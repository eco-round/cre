# EcoRound — Chainlink CRE Oracle Workflow

Go WASM workflow running on the **Chainlink Compute Runtime Environment (CRE)**. Acts as the decentralized oracle for EcoRound — polls 3 data sources, reaches 2/3 consensus, and calls `lockMatch()` / `resolveMatch()` on-chain.

## How It Works

```
cron trigger (every 30s)
  └── fetch all matches from Admin API (localhost:8080)
       └── for each open/locked match with a vault:
            ├── Confidential HTTP → PandaScore  ─┐
            ├── Confidential HTTP → VLR          ├── 2/3 consensus check
            └── Confidential HTTP → Liquipedia  ─┘
                 ├── 2+ report "started" + match is Open
                 │    └── EVM write: lockMatch()  → USDC into Morpho
                 └── 2+ report "ended" + same winner + match is Locked
                      └── EVM write: resolveMatch(winner) → yield distributed
```

## CRE SDK Features Used

| Feature | Usage |
|---|---|
| `cron.Trigger` | Periodic execution every 30s |
| `http.Fetch` | Confidential HTTP to data sources (API key injected inside enclave) |
| `evm.Client` | On-chain writes to `lockMatch` and `resolveMatch` |
| `cre.SecretsProvider` | API keys stored in DON vault, never exposed |

## Key Files

| File | Description |
|---|---|
| `EcoRound-workflow/main.go` | Workflow entry point (WASM runner) |
| `EcoRound-workflow/workflow.yaml` | CRE targets (staging / production) |
| `EcoRound-workflow/config.production.json` | Production config — schedule, factory address, chain |
| `EcoRound-workflow/config.staging.json` | Staging config |
| `secrets.yaml` | Local secrets for simulation (API keys) |
| `project.yaml` | CRE project config — chain selectors, RPC endpoints |

## Configuration

### `config.production.json`
```json
{
  "schedule": "*/30 * * * * *",
  "apiBaseUrl": "http://localhost:8080/api/v1",
  "factoryAddress": "0x602473fc59ff5eefbe5d6c86d3af5c64ac7987bc",
  "evms": [{ "chainName": "ethereum-mainnet-base-1", "gasLimit": 500000 }]
}
```

### `secrets.yaml` (local simulation only)
```yaml
secretsNames:
    PANDASCORE_API_KEY:
        - pandascore123ECO
    VLR_API_KEY:
        - vlr123ECO
    LIQUIPEDIA_API_KEY:
        - liquipedia123ECO
```

> In production, secrets are stored in the DON vault via `cre secrets set`.

## Running the Workflow

### Simulation (no on-chain writes)
```bash
cd cre
cre workflow simulate EcoRound-workflow --target production-settings
```

### Simulation with on-chain broadcast (writes to Tenderly fork)
```bash
export CRE_ETH_PRIVATE_KEY=0x...
cre workflow simulate EcoRound-workflow --target production-settings --broadcast
```

### Deploy to CRE DON
```bash
cre workflow deploy EcoRound-workflow --target production-settings
```

## Log Prefixes

| Prefix | Meaning |
|---|---|
| `[ORACLE]` | Workflow tick lifecycle |
| `[FETCH]` | Match list from Admin API |
| `[CHECK]` | Per-match consensus evaluation |
| `[ACTION]` | On-chain write triggered |
| `[TX]` | Transaction hash |
| `[DONE]` | Summary of processed matches |

## Prerequisites

- [CRE CLI](https://github.com/smartcontractkit/cre-cli) installed
- `api-simulator` running on `localhost:8080`
- For `--broadcast`: `CRE_ETH_PRIVATE_KEY` env var set, Tenderly fork RPC in `project.yaml`
