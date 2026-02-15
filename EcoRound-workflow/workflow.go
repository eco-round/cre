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
	logger.Info("EcoRound oracle tick", "apiBaseUrl", config.APIBaseUrl)

	// Step 1: Fetch all matches from admin API
	client := &http.Client{}
	matchesResp, err := http.SendRequest(config, runtime, client, fetchMatches, cre.ConsensusAggregationFromTags[*MatchesResponse]()).Await()
	if err != nil {
		logger.Error("failed to fetch matches", "err", err)
		return "", fmt.Errorf("failed to fetch matches: %w", err)
	}
	matches := matchesResp.Matches

	logger.Info("fetched matches", "count", len(matches))

	evmClient, err := config.EVMs[0].NewEVMClient()
	if err != nil {
		return "", fmt.Errorf("failed to create EVM client: %w", err)
	}

	results := make([]string, 0)

	// Process matches in REVERSE order (Oldest first)
	// This ensures we finish earlier games before newer ones, and avoids getting stuck
	// on a new "waiting" match while an older "ready" match sits in queue.
	for i := len(matches) - 1; i >= 0; i-- {
		m := matches[i]

		if m.Status != "open" && m.Status != "locked" {
			continue // skip finished/cancelled
		}
		if m.VaultAddress == "" {
			continue // no vault deployed yet
		}

		logger.Info("checking match", "id", m.ID, "status", m.Status, "vault", m.VaultAddress)

		// Fetch aggregated results from all 3 sources
		sourcesResp, err := http.SendRequest(config, runtime, client, fetchSourceResults(m.ID), cre.ConsensusAggregationFromTags[*SourcesResponse]()).Await()
		if err != nil {
			logger.Error("failed to fetch source results", "matchId", m.ID, "err", err)
			continue
		}
		sourceResults := sourcesResp.Sources

		// Check consensus
		action := checkConsensus(m.Status, sourceResults, logger)

		if action == "lock" {
			logger.Info("consensus reached: LOCK match", "matchId", m.ID)
			result, err := lockMatchOnChain(runtime, evmClient, m.VaultAddress)
			if err != nil {
				logger.Error("failed to lock match on-chain", "matchId", m.ID, "err", err)
			} else {
				logger.Info("match locked on-chain", "matchId", m.ID, "txHash", result)
				results = append(results, fmt.Sprintf("locked:%d", m.ID))
			}
		} else if action == "resolve" {
			winner := getConsensusWinner(sourceResults)
			logger.Info("consensus reached: RESOLVE match", "matchId", m.ID, "winner", winner)
			result, err := resolveMatchOnChain(runtime, evmClient, m.VaultAddress, winner)
			if err != nil {
				logger.Error("failed to resolve match on-chain", "matchId", m.ID, "err", err)
			} else {
				logger.Info("match resolved on-chain", "matchId", m.ID, "winner", winner, "txHash", result)
				results = append(results, fmt.Sprintf("resolved:%d:winner=%d", m.ID, winner))
			}
		} else {
			logger.Info("no consensus yet", "matchId", m.ID, "action", action)
		}

		// LIMIT: Stop after processing 1 match to stay under 5 HTTP calls
		break 
	}

	summary := fmt.Sprintf("processed %d matches, actions: %v", len(matches), results)
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
			logger.Info("consensus: 2+ sources report started", "count", startedCount)
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
					logger.Info("consensus: 2+ sources agree on winner", "winner", winner, "votes", votes)
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
