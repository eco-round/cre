//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

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
	workflow := cre.Workflow[*Config]{
		cre.Handler(
			cron.Trigger(&cron.Config{Schedule: config.Schedule}),
			onCronTrigger,
		),
	}
	return workflow, nil
}

// ── Confidential HTTP — Secret Helper ───────────────────────────────────

// getSecretValue fetches a secret from the Chainlink DON Vault via enclave-protected
// runtime.GetSecret(). In simulation, secrets are resolved from the .env file.
// Secrets are always fetched SEQUENTIALLY — never in parallel.
func getSecretValue(runtime cre.Runtime, logger *slog.Logger, id string) string {
	result, err := runtime.GetSecret(&pb2.SecretRequest{Id: id}).Await()
	if err != nil || result == nil {
		logger.Warn("[ORACLE] Could not fetch secret — proceeding without auth", "id", id, "err", err)
		return ""
	}
	return result.Value
}

// ── Cron Handler ────────────────────────────────────────────────────────

func onCronTrigger(config *Config, runtime cre.Runtime, outputs *cron.Payload) (string, error) {
	logger := runtime.Logger()
	logger.Info("[ORACLE] Tick started", "api", config.APIBaseUrl)

	// ── Confidential HTTP: fetch API keys sequentially from DON Vault ──────
	// Keys are injected inside the CRE enclave — never visible to node operators.
	// In local simulation, values are resolved from PANDASCORE_API_KEY etc. in .env
	sourceKeys := map[string]string{
		"pandascore": getSecretValue(runtime, logger, "PANDASCORE_API_KEY"),
		"vlr":        getSecretValue(runtime, logger, "VLR_API_KEY"),
		"liquipedia": getSecretValue(runtime, logger, "LIQUIPEDIA_API_KEY"),
	}
	logger.Info("[ORACLE] Confidential keys loaded",
		"pandascore", sourceKeys["pandascore"] != "",
		"vlr", sourceKeys["vlr"] != "",
		"liquipedia", sourceKeys["liquipedia"] != "",
	)

	// Step 1: Fetch all matches from admin API
	client := &http.Client{}
	matchesResp, err := http.SendRequest(config, runtime, client, fetchMatches, cre.ConsensusAggregationFromTags[*MatchesResponse]()).Await()
	if err != nil {
		logger.Error("[ORACLE] Failed to fetch matches from API", "err", err)
		return "", fmt.Errorf("failed to fetch matches: %w", err)
	}
	matches := matchesResp.Matches

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

	// Process matches in reverse order (oldest first).
	// Limited to 1 match per tick to stay within CRE HTTP call budget.
	for i := len(matches) - 1; i >= 0; i-- {
		m := matches[i]

		if m.Status != "open" && m.Status != "locked" {
			continue
		}
		if m.VaultAddress == "" {
			continue
		}

		vaultShort := m.VaultAddress[:6] + ".." + m.VaultAddress[len(m.VaultAddress)-4:]
		logger.Info(fmt.Sprintf("[CHECK] Match #%d | %s | vault=%s", m.ID, strings.ToUpper(m.Status), vaultShort))

		sourcesResp, err := http.SendRequest(config, runtime, client, fetchSourceResults(m.ID, sourceKeys), cre.ConsensusAggregationFromTags[*SourcesResponse]()).Await()
		if err != nil {
			logger.Error(fmt.Sprintf("[CHECK] Match #%d — failed to fetch sources", m.ID), "err", err)
			continue
		}

		logSourceSummary(logger, m.ID, sourcesResp.Sources)

		switch checkConsensus(m.Status, sourcesResp.Sources) {
		case "lock":
			startedCount := countByStatus(sourcesResp.Sources, "started")
			logger.Info(fmt.Sprintf("[ACTION] LOCK Match #%d — consensus: %d/3 sources report started", m.ID, startedCount))
			txHash, err := lockMatchOnChain(runtime, evmClient, m.VaultAddress)
			if err != nil {
				logger.Error(fmt.Sprintf("[TX] Match #%d lock FAILED", m.ID), "err", err)
			} else {
				logger.Info(fmt.Sprintf("[TX] Match #%d locked | tx=%s", m.ID, shortHash(txHash)))
				actions = append(actions, fmt.Sprintf("locked:#%d", m.ID))
			}

		case "resolve":
			winnerUint, winnerName := consensusWinner(sourcesResp.Sources)
			winnerLabel := "TeamA"
			if winnerUint == 2 {
				winnerLabel = "TeamB"
			}
			endedCount := countByStatus(sourcesResp.Sources, "ended")
			logger.Info(fmt.Sprintf("[ACTION] RESOLVE Match #%d — consensus: %d/3 agree winner=%s (%s)", m.ID, endedCount, winnerName, winnerLabel))
			txHash, err := resolveMatchOnChain(runtime, evmClient, m.VaultAddress, winnerUint)
			if err != nil {
				logger.Error(fmt.Sprintf("[TX] Match #%d resolve FAILED", m.ID), "err", err)
			} else {
				logger.Info(fmt.Sprintf("[TX] Match #%d resolved | winner=%s | tx=%s", m.ID, winnerName, shortHash(txHash)))
				actions = append(actions, fmt.Sprintf("resolved:#%d→%s", m.ID, winnerName))
			}

		default:
			logger.Info(fmt.Sprintf("[CHECK] Match #%d — no consensus yet, waiting", m.ID))
		}

		break // one match per tick (CRE HTTP budget)
	}

	summary := fmt.Sprintf("[DONE] Processed %d matches → no actions taken", len(matches))
	if len(actions) > 0 {
		summary = fmt.Sprintf("[DONE] Processed %d matches → %v", len(matches), actions)
	}
	logger.Info(summary)
	return summary, nil
}

// ── HTTP Fetchers ───────────────────────────────────────────────────────

func fetchMatches(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*MatchesResponse, error) {
	url := config.APIBaseUrl + "/admin/matches"
	resp, err := sendRequester.SendRequest(&http.Request{Method: "GET", Url: url}).Await()
	if err != nil {
		return nil, fmt.Errorf("GET %s failed: %w", url, err)
	}
	var response MatchesResponse
	if err := json.Unmarshal(resp.Body, &response); err != nil {
		return nil, fmt.Errorf("unmarshal matches: %w", err)
	}
	return &response, nil
}

// fetchSourceResults returns a fetcher that queries all 3 data sources for a match.
// API keys are injected via Chainlink Confidential HTTP — resolved inside the CRE
// enclave from the DON Vault, never visible to node operators or logs.
func fetchSourceResults(matchID uint, sourceKeys map[string]string) func(*Config, *slog.Logger, *http.SendRequester) (*SourcesResponse, error) {
	return func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*SourcesResponse, error) {
		sources := []string{"pandascore", "vlr", "liquipedia"}
		results := make([]SourceResult, 0, len(sources))

		for _, source := range sources {
			url := fmt.Sprintf("%s/%s/matches/%d", config.APIBaseUrl, source, matchID)

			headers := map[string]string{}
			if key := sourceKeys[source]; key != "" {
				headers["X-Api-Key"] = key
			}

			resp, err := sendRequester.SendRequest(&http.Request{
				Method:  "GET",
				Url:     url,
				Headers: headers,
			}).Await()
			if err != nil {
				logger.Error("failed to fetch from source", "source", source, "err", err)
				continue
			}
			if resp.StatusCode == 401 {
				logger.Warn(fmt.Sprintf("[CHECK] Source %s rejected API key (401) — skipping", source))
				continue
			}

			var match APIMatch
			if err := json.Unmarshal(resp.Body, &match); err != nil {
				logger.Error("failed to unmarshal source response", "source", source, "err", err)
				continue
			}

			sr := SourceResult{Source: source, Confident: true}
			if match.Result != nil {
				sr.MatchStatus = match.Result.MatchStatus
				sr.Winner = match.Result.Winner
			}
			results = append(results, sr)
		}

		return &SourcesResponse{Sources: results}, nil
	}
}

// ── Consensus Logic ──────────────────────────────────────────────────────

func checkConsensus(currentStatus string, sources []SourceResult) string {
	if len(sources) < 2 {
		return "wait"
	}
	switch currentStatus {
	case "open":
		count := 0
		for _, s := range sources {
			if s.MatchStatus == "started" || s.MatchStatus == "ended" {
				count++
			}
		}
		if count >= 2 {
			return "lock"
		}
	case "locked":
		endedVotes := map[string]int{}
		for _, s := range sources {
			if s.MatchStatus == "ended" && s.Winner != "" {
				endedVotes[s.Winner]++
			}
		}
		for _, votes := range endedVotes {
			if votes >= 2 {
				return "resolve"
			}
		}
	}
	return "wait"
}

// consensusWinner returns the winner's uint8 enum value and display name.
func consensusWinner(sources []SourceResult) (uint8, string) {
	votes := map[string]int{}
	for _, s := range sources {
		if s.MatchStatus == "ended" && s.Winner != "" {
			votes[s.Winner]++
		}
	}
	best, bestVotes := "", 0
	for winner, count := range votes {
		if count > bestVotes {
			best, bestVotes = winner, count
		}
	}
	if best == "teamA" || best == "1" {
		return 1, best
	}
	return 2, best
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

// ── Log Helpers ──────────────────────────────────────────────────────────

func logSourceSummary(logger *slog.Logger, matchID uint, sources []SourceResult) {
	parts := make([]string, 0, 3)
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
	for len(parts) < 3 {
		parts = append(parts, "?=no-response")
	}
	logger.Info(fmt.Sprintf("[CHECK] Match #%d sources: %s", matchID, strings.Join(parts, ", ")))
}

func shortHash(hash string) string {
	if len(hash) < 14 {
		return hash
	}
	return hash[:10] + ".." + hash[len(hash)-8:]
}

// ── On-Chain Writes ──────────────────────────────────────────────────────

func lockMatchOnChain(runtime cre.Runtime, evmClient *evm.Client, vaultAddr string) (string, error) {
	codec, err := vault_match.NewCodec()
	if err != nil {
		return "", fmt.Errorf("new codec: %w", err)
	}
	calldata, err := codec.EncodeLockMatchMethodCall()
	if err != nil {
		return "", fmt.Errorf("encode lockMatch: %w", err)
	}
	return writeOnChain(runtime, evmClient, vaultAddr, calldata)
}

func resolveMatchOnChain(runtime cre.Runtime, evmClient *evm.Client, vaultAddr string, winner uint8) (string, error) {
	codec, err := vault_match.NewCodec()
	if err != nil {
		return "", fmt.Errorf("new codec: %w", err)
	}
	calldata, err := codec.EncodeResolveMatchMethodCall(vault_match.ResolveMatchInput{Winner: winner})
	if err != nil {
		return "", fmt.Errorf("encode resolveMatch: %w", err)
	}
	return writeOnChain(runtime, evmClient, vaultAddr, calldata)
}

// writeOnChain generates a CRE report and submits it as an EVM transaction.
func writeOnChain(runtime cre.Runtime, evmClient *evm.Client, vaultAddr string, calldata []byte) (string, error) {
	report, err := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: calldata,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await() 
	if err != nil {
		return "", fmt.Errorf("generate report: %w", err)
	}

	resp, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  common.HexToAddress(vaultAddr).Bytes(),
		Report:    report,
		GasConfig: nil,
	}).Await()
	if err != nil {
		return "", fmt.Errorf("write report: %w", err)
	}

	return common.BytesToHash(resp.TxHash).Hex(), nil
}
