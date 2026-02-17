//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"cre/contracts/evm/src/generated/vault_match"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

// ── Config ──────────────────────────────────────────────────────────────

type EVMConfig struct {
	ChainName string `json:"chainName"`
	GasLimit  uint64 `json:"gasLimit"`
}

func (e *EVMConfig) GetChainSelector() (uint64, error) {
	return evm.ChainSelectorFromName(e.ChainName)
}

func (e *EVMConfig) NewEVMClient() (*evm.Client, error) {
	chainSelector, err := e.GetChainSelector()
	if err != nil {
		return nil, err
	}
	return &evm.Client{ChainSelector: chainSelector}, nil
}

type Config struct {
	Schedule       string      `json:"schedule"`
	APIBaseUrl     string      `json:"apiBaseUrl"`
	FactoryAddress string      `json:"factoryAddress"`
	EVMs           []EVMConfig `json:"evms"`
}

// ── API Response Types ──────────────────────────────────────────────────

type APIMatch struct {
	ID           uint       `json:"id"`
	Status       string     `json:"status"`
	VaultAddress string     `json:"vault_address"`
	Result       *APIResult `json:"result,omitempty"`
}

type APIResult struct {
	MatchStatus string `json:"match_status"`
	Winner      string `json:"winner"`
}

// Wrapper types with consensus tags for CRE SDK
type MatchesResponse struct {
	Matches []APIMatch `consensus_aggregation:"identical" json:"matches"`
}

type SourceResult struct {
	Source      string `json:"source"`
	MatchStatus string `json:"match_status"`
	Winner      string `json:"winner"`
	Confident   bool   `json:"confident"`
}

type SourcesResponse struct {
	Sources []SourceResult `consensus_aggregation:"identical" json:"sources"`
}

// ── Workflow Setup ──────────────────────────────────────────────────────

func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	cronTriggerCfg := &cron.Config{
		Schedule: config.Schedule,
	}

	workflow := cre.Workflow[*Config]{
		cre.Handler(
			cron.Trigger(cronTriggerCfg),
			onCronTrigger,
		),
	}

	return workflow, nil
}

// ── Cron Handler ────────────────────────────────────────────────────────

func onCronTrigger(config *Config, runtime cre.Runtime, outputs *cron.Payload) (string, error) {
	logger := runtime.Logger()
	logger.Info("[ORACLE] Tick started", "api", config.APIBaseUrl)

	// Step 1: Fetch all matches from admin API
	client := &http.Client{}
	matchesResp, err := http.SendRequest(config, runtime, client, fetchMatches, cre.ConsensusAggregationFromTags[*MatchesResponse]()).Await()
	if err != nil {
		logger.Error("[ORACLE] Failed to fetch matches from API", "err", err)
		return "", fmt.Errorf("failed to fetch matches: %w", err)
	}
	matches := matchesResp.Matches

	// Count by status for summary
	openCount, lockedCount, skippedCount := 0, 0, 0
	for _, m := range matches {
		switch m.Status {
		case "open":
			openCount++
		case "locked":
			lockedCount++
		default:
			skippedCount++
		}
	}
	logger.Info(fmt.Sprintf("[FETCH] Found %d matches (%d open, %d locked, %d skipped)", len(matches), openCount, lockedCount, skippedCount))

	evmClient, err := config.EVMs[0].NewEVMClient()
	if err != nil {
		return "", fmt.Errorf("failed to create EVM client: %w", err)
	}

	var actions []string

	// Process matches in REVERSE order (oldest first)
	for i := len(matches) - 1; i >= 0; i-- {
		m := matches[i]

		if m.Status != "open" && m.Status != "locked" {
			continue
		}
		if m.VaultAddress == "" {
			continue
		}

		vaultShort := m.VaultAddress[:6] + ".." + m.VaultAddress[len(m.VaultAddress)-4:]
		logger.Info(fmt.Sprintf("[CHECK] Match #%d | %s | vault=%s", m.ID, upper(m.Status), vaultShort))

		// Fetch aggregated results from all 3 sources
		sourcesResp, err := http.SendRequest(config, runtime, client, fetchSourceResults(m.ID), cre.ConsensusAggregationFromTags[*SourcesResponse]()).Await()
		if err != nil {
			logger.Error(fmt.Sprintf("[CHECK] Match #%d — failed to fetch sources", m.ID), "err", err)
			continue
		}
		sourceResults := sourcesResp.Sources

		// Log individual source statuses
		logSourceSummary(logger, m.ID, sourceResults)

		// Check consensus
		action := checkConsensus(m.Status, sourceResults, logger)

		if action == "lock" {
			startedCount := countByStatus(sourceResults, "started")
			logger.Info(fmt.Sprintf("[ACTION] LOCK Match #%d — consensus: %d/3 sources report started", m.ID, startedCount))
			result, err := lockMatchOnChain(runtime, evmClient, m.VaultAddress)
			if err != nil {
				logger.Error(fmt.Sprintf("[TX] Match #%d lock FAILED", m.ID), "err", err)
			} else {
				txShort := result[:10] + ".." + result[len(result)-8:]
				logger.Info(fmt.Sprintf("[TX] Match #%d locked on-chain | tx=%s", m.ID, txShort))
				actions = append(actions, fmt.Sprintf("locked:#%d", m.ID))
			}
		} else if action == "resolve" {
			winner := getConsensusWinner(sourceResults)
			winnerLabel := "TeamA"
			if winner == 2 {
				winnerLabel = "TeamB"
			}
			endedCount := countByStatus(sourceResults, "ended")
			consensusWinnerName := getConsensusWinnerName(sourceResults)
			logger.Info(fmt.Sprintf("[ACTION] RESOLVE Match #%d — consensus: %d/3 sources agree winner=%s (%s)", m.ID, endedCount, consensusWinnerName, winnerLabel))
			result, err := resolveMatchOnChain(runtime, evmClient, m.VaultAddress, winner)
			if err != nil {
				logger.Error(fmt.Sprintf("[TX] Match #%d resolve FAILED", m.ID), "err", err)
			} else {
				txShort := result[:10] + ".." + result[len(result)-8:]
				logger.Info(fmt.Sprintf("[TX] Match #%d resolved on-chain | winner=%s | tx=%s", m.ID, consensusWinnerName, txShort))
				actions = append(actions, fmt.Sprintf("resolved:#%d→%s", m.ID, consensusWinnerName))
			}
		} else {
			logger.Info(fmt.Sprintf("[CHECK] Match #%d — no consensus yet, waiting", m.ID))
		}

		// LIMIT: Stop after processing 1 match to stay under 5 HTTP calls
		break
	}

	summary := fmt.Sprintf("[DONE] Processed %d matches → %v", len(matches), actions)
	if len(actions) == 0 {
		summary = fmt.Sprintf("[DONE] Processed %d matches → no actions taken", len(matches))
	}
	logger.Info(summary)
	return summary, nil
}

// ── HTTP Fetchers ───────────────────────────────────────────────────────

func fetchMatches(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*MatchesResponse, error) {
	url := config.APIBaseUrl + "/admin/matches"
	resp, err := sendRequester.SendRequest(&http.Request{
		Method: "GET",
		Url:    url,
	}).Await()
	if err != nil {
		return nil, fmt.Errorf("GET %s failed: %w", url, err)
	}

	var response MatchesResponse
	if err := json.Unmarshal(resp.Body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal matches: %w", err)
	}
	return &response, nil
}

func fetchSourceResults(matchID uint) func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*SourcesResponse, error) {
	return func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*SourcesResponse, error) {
		sources := []string{"pandascore", "vlr", "liquipedia"}
		results := make([]SourceResult, 0, len(sources))

		for _, source := range sources {
			url := fmt.Sprintf("%s/%s/matches/%d", config.APIBaseUrl, source, matchID)

			resp, err := sendRequester.SendRequest(&http.Request{
				Method: "GET",
				Url:    url,
			}).Await()
			if err != nil {
				logger.Error("failed to fetch from source", "source", source, "err", err)
				continue
			}

			var match APIMatch
			if err := json.Unmarshal(resp.Body, &match); err != nil {
				logger.Error("failed to unmarshal source response", "source", source, "err", err)
				continue
			}

			sr := SourceResult{
				Source:    source,
				Confident: true,
			}
			if match.Result != nil {
				sr.MatchStatus = match.Result.MatchStatus
				sr.Winner = match.Result.Winner
			}
			results = append(results, sr)
		}

		return &SourcesResponse{Sources: results}, nil
	}
}

// ── Consensus Logic ─────────────────────────────────────────────────────

// ── Log Helpers ──────────────────────────────────────────────────────────

func upper(s string) string {
	if len(s) == 0 {
		return s
	}
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(c - 32)
		} else {
			result += string(c)
		}
	}
	return result
}

func logSourceSummary(logger *slog.Logger, matchID uint, sources []SourceResult) {
	parts := make([]string, 0, len(sources))
	for _, s := range sources {
		status := s.MatchStatus
		if status == "" {
			status = "pending"
		}
		entry := s.Source + "=" + status
		if s.Winner != "" {
			entry += "(" + s.Winner + ")"
		}
		parts = append(parts, entry)
	}
	missing := 3 - len(sources)
	for i := 0; i < missing; i++ {
		parts = append(parts, "?=no-response")
	}
	summary := ""
	for i, p := range parts {
		if i > 0 {
			summary += ", "
		}
		summary += p
	}
	logger.Info(fmt.Sprintf("[CHECK] Match #%d sources: %s", matchID, summary))
}

func countByStatus(sources []SourceResult, status string) int {
	count := 0
	for _, s := range sources {
		if s.MatchStatus == status {
			count++
		}
	}
	return count
}

func getConsensusWinnerName(sources []SourceResult) string {
	votes := map[string]int{}
	for _, s := range sources {
		if s.MatchStatus == "ended" && s.Winner != "" {
			votes[s.Winner]++
		}
	}
	bestWinner := ""
	bestVotes := 0
	for winner, count := range votes {
		if count > bestVotes {
			bestWinner = winner
			bestVotes = count
		}
	}
	return bestWinner
}

func checkConsensus(currentStatus string, sources []SourceResult, logger *slog.Logger) string {
	if len(sources) < 2 {
		return "wait" // need at least 2 sources
	}

	if currentStatus == "open" {
		// Check if 2/3 sources report "started"
		startedCount := 0
		for _, s := range sources {
			if s.MatchStatus == "started" || s.MatchStatus == "ended" {
				startedCount++
			}
		}
		if startedCount >= 2 {
			return "lock"
		}
	}

	if currentStatus == "locked" {
		// Check if 2/3 sources report "ended" and agree on winner
		endedCount := 0
		winnerVotes := map[string]int{}
		for _, s := range sources {
			if s.MatchStatus == "ended" && s.Winner != "" {
				endedCount++
				winnerVotes[s.Winner]++
			}
		}
		if endedCount >= 2 {
			// Check winner consensus
			for winner, votes := range winnerVotes {
				if votes >= 2 {
					return "resolve"
				}
			}
		}
	}

	return "wait"
}

func getConsensusWinner(sources []SourceResult) uint8 {
	// Count votes for each winner
	votes := map[string]int{}
	for _, s := range sources {
		if s.MatchStatus == "ended" && s.Winner != "" {
			votes[s.Winner]++
		}
	}

	// Find winner with most votes
	bestWinner := ""
	bestVotes := 0
	for winner, count := range votes {
		if count > bestVotes {
			bestWinner = winner
			bestVotes = count
		}
	}

	// Map team name to uint8 (1 = TeamA, 2 = TeamB)
	// The winner string from the API is the team tag/name
	// We return 1 for first team, 2 for second team
	// The actual mapping will depend on the match data
	if bestWinner == "teamA" || bestWinner == "1" {
		return 1
	}
	return 2 // default to team B if not explicitly team A
}

// ── On-Chain Actions ────────────────────────────────────────────────────

func lockMatchOnChain(runtime cre.Runtime, evmClient *evm.Client, vaultAddr string) (string, error) {
	codec, err := vault_match.NewCodec()
	if err != nil {
		return "", fmt.Errorf("failed to create codec: %w", err)
	}

	calldata, err := codec.EncodeLockMatchMethodCall()
	if err != nil {
		return "", fmt.Errorf("failed to encode lockMatch: %w", err)
	}

	addr := common.HexToAddress(vaultAddr)
	report, err := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: calldata,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return "", fmt.Errorf("failed to generate report: %w", err)
	}

	resp, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  addr.Bytes(),
		Report:    report,
		GasConfig: nil,
	}).Await()
	if err != nil {
		return "", fmt.Errorf("lockMatch tx failed: %w", err)
	}

	return common.BytesToHash(resp.TxHash).Hex(), nil
}

func resolveMatchOnChain(runtime cre.Runtime, evmClient *evm.Client, vaultAddr string, winner uint8) (string, error) {
	codec, err := vault_match.NewCodec()
	if err != nil {
		return "", fmt.Errorf("failed to create codec: %w", err)
	}

	calldata, err := codec.EncodeResolveMatchMethodCall(vault_match.ResolveMatchInput{
		Winner: winner,
	})
	if err != nil {
		return "", fmt.Errorf("failed to encode resolveMatch: %w", err)
	}

	addr := common.HexToAddress(vaultAddr)
	report, err := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: calldata,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return "", fmt.Errorf("failed to generate report: %w", err)
	}

	resp, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  addr.Bytes(),
		Report:    report,
		GasConfig: nil,
	}).Await()
	if err != nil {
		return "", fmt.Errorf("resolveMatch tx failed: %w", err)
	}

	return common.BytesToHash(resp.TxHash).Hex(), nil
}

// unused but kept for potential future use
var _ = big.NewInt
