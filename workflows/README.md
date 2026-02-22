# Optimus CRE Workflows

Chainlink Runtime Environment (CRE) workflow implementations for the Optimus protocol.  
Every complex multi-step process is expressed as a CRE workflow compiled to WASM and executed by the DON (Decentralized Oracle Network).

## Directory Layout

```
workflows/
├── project.yaml                   # Global CRE project config (RPCs, secrets path)
├── secrets.yaml                   # Vault DON secrets schema
├── .env                           # Local secret values (not committed)
├── go.mod                         # Module for all workflows
├── shared/                        # Shared helpers imported by every workflow
│   ├── abi.go                     # ABI Pack/Unpack for all contract functions
│   ├── events.go                  # Event data decoders (log.Data → structs)
│   ├── risk.go                    # Risk score arithmetic, credit tiers
│   └── config.go                  # Config struct loaded from config.*.json
│
├── bnpl_late_fee/                 # Cron: every 30 min
│   ├── main.go                    #   Scan active BNPLs → applyLateFee on-chain
│   ├── workflow.yaml
│   └── config.staging.json
│
├── loan_default_monitor/          # Cron: every 1 hour
│   ├── main.go                    #   Scan active loans → markDefaulted on-chain
│   ├── workflow.yaml
│   └── config.staging.json
│
├── dao_proposal_finalizer/        # Cron: every 1 hour
│   ├── main.go                    #   Finalize expired proposals on-chain
│   ├── workflow.yaml
│   └── config.staging.json
│
├── treasury_monitor/              # Cron: every 6 hours
│   ├── main.go                    #   Monitor TokenVault + DAO treasury balances
│   ├── workflow.yaml
│   └── config.staging.json
│
├── state_reconciler/              # Cron: every 2 hours
│   ├── main.go                    #   Reconcile offchain state vs on-chain
│   ├── workflow.yaml
│   └── config.staging.json
│
├── bnpl_created/                  # EVM Log: BNPLCreated
│   ├── main.go                    #   Validate terms, schedule reminders
│   ├── workflow.yaml
│   └── config.staging.json
│
├── bnpl_payment/                  # EVM Log: BNPLPaymentMade
│   ├── main.go                    #   Track installments, detect completion
│   ├── workflow.yaml
│   └── config.staging.json
│
├── bnpl_completed/                # EVM Log: BNPLCompleted
│   ├── main.go                    #   Final accounting, treasury rebalance
│   ├── workflow.yaml
│   └── config.staging.json
│
├── loan_created/                  # EVM Log: LoanCreated
│   ├── main.go                    #   Register loan, risk profile check
│   ├── workflow.yaml
│   └── config.staging.json
│
├── loan_approved/                 # EVM Log: LoanApproved
│   ├── main.go                    #   Start interest monitoring
│   ├── workflow.yaml
│   └── config.staging.json
│
├── loan_payment/                  # EVM Log: PaymentMade
│   ├── main.go                    #   Track repayment, detect full repayment
│   ├── workflow.yaml
│   └── config.staging.json
│
├── dao_proposal_opened/           # EVM Log: ProposalOpened
│   ├── main.go                    #   Parse proposal, register in backend
│   ├── workflow.yaml
│   └── config.staging.json
│
├── dao_vote_cast/                 # EVM Log: VoteCast
│   ├── main.go                    #   Update tally, check quorum
│   ├── workflow.yaml
│   └── config.staging.json
│
└── risk_profile_updater/          # EVM Log: RiskProfileUpdated
    ├── main.go                    #   Credit tier re-evaluation
    ├── workflow.yaml
    └── config.staging.json
```

## Workflows Summary

### Cron-Triggered (Scheduled)

| Workflow ID | Schedule | Description |
|---|---|---|
| `optimus_bnpl_late_fee` | `*/30 * * * *` | Scans active BNPL arrangements for overdue installments, applies late fees on-chain, and lowers payer risk scores |
| `optimus_loan_default_monitor` | `0 * * * *` | Identifies loans past their EndTime with outstanding balance, marks them defaulted, and lowers borrower risk scores |
| `optimus_dao_proposal_finalizer` | `0 * * * *` | Finalizes DAO proposals whose voting period has expired |
| `optimus_treasury_monitor` | `0 */6 * * *` | Monitors TokenVault and DAO treasury balances, sends alerts when below threshold |
| `optimus_state_reconciler` | `0 */2 * * *` | Reconciles backend offchain state with on-chain contract state for all entities |

### EVM Log-Triggered (Event-Driven)

| Workflow ID | Event | Contract | Description |
|---|---|---|---|
| `optimus_bnpl_created` | `BNPLCreated` | BNPLManager | Validates arrangement against DAO terms, schedules payments |
| `optimus_bnpl_payment` | `BNPLPaymentMade` | BNPLManager | Tracks installments, detects completion, adjusts risk score |
| `optimus_bnpl_completed` | `BNPLCompleted` | BNPLManager | Final accounting, treasury rebalance proposal |
| `optimus_loan_created` | `LoanCreated` | LoanManager | Registers loan, checks borrower credit tier |
| `optimus_loan_approved` | `LoanApproved` | LoanManager | Starts interest monitoring, notifies backend |
| `optimus_loan_payment` | `PaymentMade` | LoanManager | Tracks repayment, detects full repayment, boosts risk |
| `optimus_dao_proposal_opened` | `ProposalOpened` | DAOManager | Parses treasury proposals, starts vote tracking |
| `optimus_dao_vote_cast` | `VoteCast` | DAOManager | Updates tally cache, checks quorum |
| `optimus_risk_profile_updater` | `RiskProfileUpdated` | DIDRegistry | Re-evaluates credit tier limits, detects tier changes |

## Risk Score Model

All workflows share a unified credit scoring system via `shared/risk.go`:

| Event | Delta (bps) | Direction |
|---|---|---|
| On-time BNPL installment | +50 | Positive |
| BNPL completed on-time | +500 | Positive |
| Late fee applied | -300 | Negative |
| On-time loan payment | +100 | Positive |
| Loan fully repaid | +700 | Positive |
| Loan defaulted | -2000 | Negative |

**Credit Tiers:**

| Tier | Score Range | Max BNPL | Max Loan |
|---|---|---|---|
| EXCELLENT | 7000–10000 | 10 ETH | 50 ETH |
| GOOD | 5000–6999 | 5 ETH | 20 ETH |
| FAIR | 3000–4999 | 1 ETH | 5 ETH |
| POOR | 0–2999 | 0.5 ETH | 1 ETH |

## Development

### Simulate a workflow

```bash
cd workflows/bnpl_late_fee
cre workflow simulate --config config.staging.json
```

### Deploy a workflow

```bash
cre login
cd workflows/bnpl_late_fee
cre workflow deploy --config config.staging.json
cre workflow activate --workflow-id optimus_bnpl_late_fee
```

### Deploy all workflows

```bash
for dir in workflows/*/; do
  if [ -f "$dir/workflow.yaml" ]; then
    echo "Deploying $(basename $dir)..."
    cd "$dir"
    cre workflow deploy --config config.staging.json
    wid=$(grep workflow_id workflow.yaml | awk '{print $2}')
    cre workflow activate --workflow-id "$wid"
    cd ../..
  fi
done
```

## Contract Addresses (Sepolia)

| Contract | Address |
|---|---|
| BNPLManager | `0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba` |
| LoanManager | `0xbB0D4067488edf4a007822407e2486412dC8D39D` |
| DAOManager | `0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b` |
| TokenVault | `0x4C704D51fc47cfe582F8c5477de3AE398B344907` |
| DIDRegistry | `0x0E9D8959bCD99e7AFD7C693e51781058A998b756` |

## Architecture

```
On-Chain Events ──┐
                  ├──▶ CRE DON (WASM Workflows) ──▶ On-Chain Writes
Cron Triggers ────┘              │                   (applyLateFee,
                                 │                    markDefaulted,
                                 │                    finalizeProposal,
                                 ▼                    updateRiskProfile)
                          Backend API
                    (offchain state, notifications,
                     tally cache, reconciliation)
```

Each workflow is compiled to WASM and executed on the Chainlink DON.  
The workflows use:
- **EVMClient** for reading contract state (`callContract`) and writing signed reports (`writeReport`)
- **HTTPClient** for backend API communication
- **Runtime** for logging and report signing
