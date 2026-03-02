# Optimus Protocol

Optimus is a decentralised financial protocol built on Ethereum (Sepolia testnet) that combines **on-chain smart contracts**, a **Go backend API**, a **Flutter mobile app**, and **Chainlink CRE (Compute Runtime Environment) workflows** to deliver identity-aware DeFi services — including BNPL, lending, DAO governance, and token vaults.

The protocol's defining feature is its deep integration with Chainlink CRE: 14 autonomous WASM workflows running on the CRE DON (Decentralised Oracle Network) handle event reactions, risk scoring, default detection, late fees, proposal finalization, and state reconciliation — keeping all off-chain automation trustless and verifiable.

---

## Table of Contents

- [Project Structure](#project-structure)
- [Smart Contracts](#smart-contracts)
- [Go Backend (Protocol)](#go-backend-protocol)
- [Flutter Mobile App (Client)](#flutter-mobile-app-client)
- [Chainlink CRE Integration](#chainlink-cre-integration)
- [CRE Workflow Catalogue](#cre-workflow-catalogue)
- [How CRE Is Used](#how-cre-is-used)
- [Deployment](#deployment)
- [Documentation](#documentation)

---

## Project Structure

```
optimus-BOC/
├── contracts/              # Foundry project — Solidity smart contracts
│   ├── src/                # 5 core contracts + interfaces
│   ├── test/               # Forge tests
│   └── script/             # Deployment scripts
├── protocol/               # Go backend API server
│   ├── controllers/        # HTTP handlers per domain (did, bnpl, loan, dao, vault)
│   ├── services/           # Business logic interfaces + implementations
│   ├── bindings/           # Auto-generated Go bindings (abigen)
│   ├── middleware/          # Privy JWT auth middleware
│   ├── store/              # PostgreSQL data layer
│   └── main.go             # Entry point (chi router, :8000)
├── client/prime/           # Flutter mobile app
│   └── lib/
│       ├── screens/        # 5 domain screens + login
│       ├── services/       # Provider-based services + API client
│       └── widgets/        # Shared UI components
├── workflows/              # 14 Chainlink CRE workflows (Go → WASM)
│   ├── shared/             # Common config, helpers, contract addresses
│   ├── bnpl_*/             # 4 BNPL workflows
│   ├── loan_*/             # 4 loan workflows
│   ├── dao_*/              # 3 DAO workflows
│   ├── risk_profile_updater/
│   ├── state_reconciler/
│   └── treasury_monitor/
├── docs/                   # PlantUML diagrams & markdown documentation
└── scripts/                # Utility scripts
```

---

## Smart Contracts

Five Solidity contracts deployed on **Sepolia (chain 11155111)**:

| Contract | Address | Source |
|----------|---------|--------|
| **DIDRegistry** | `0x0E9D33b0cA7c7e93F4F5e413C35cE0E518040296` | [`contracts/src/DIDRegistry.sol`](contracts/src/DIDRegistry.sol) |
| **BNPLManager** | `0x4d99F01530F9b5Ee58C13E52a114d79cF397c1f4` | [`contracts/src/BNPLManager.sol`](contracts/src/BNPLManager.sol) |
| **LoanManager** | `0xbB0D4067488edf4a007822407e2486412dC8D39D` | [`contracts/src/LoanManager.sol`](contracts/src/LoanManager.sol) |
| **DAOManager** | `0x5612dB018f17E82D76E79a77aA75Bbf42E40C9B1` | [`contracts/src/DAOManager.sol`](contracts/src/DAOManager.sol) |
| **TokenVault** | `0x4C70E0Ae715d04e3a6e65a70191276a5D0E5B66a` | [`contracts/src/TokenVault.sol`](contracts/src/TokenVault.sol) |

**Deployer:** `0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780`

Contract documentation: [`docs/contracts/`](docs/contracts/)

---

## Go Backend (Protocol)

A chi-based HTTP API server that mediates between the Flutter client and on-chain contracts.

- **Auth:** Privy RS256 JWT verification middleware
- **Transactor:** Ethereum transaction signing via `PRIVATE_KEY` env var
- **Database:** PostgreSQL (pgxpool) for off-chain state (loans, arrangements, DAOs)
- **Port:** `:8000`

32 API routes across 5 domain controllers:

| Domain | Routes | Controller |
|--------|--------|------------|
| DID | 4 routes | [`protocol/controllers/did/`](protocol/controllers/did/) |
| BNPL | 8 routes | [`protocol/controllers/bnpl/`](protocol/controllers/bnpl/) |
| Loan | 7 routes | [`protocol/controllers/loan/`](protocol/controllers/loan/) |
| DAO | 12 routes | [`protocol/controllers/dao/`](protocol/controllers/dao/) |
| Vault | 3 routes | [`protocol/controllers/vault/`](protocol/controllers/vault/) |

Controller documentation: [`docs/protocol/`](docs/protocol/)

---

## Flutter Mobile App (Client)

A cross-platform mobile application built with Flutter 3.11+ and Provider state management.

- **Auth:** Privy Flutter SDK (`privy_flutter: ^0.6.0`) — email/OTP login
- **5 screens:** Identity (DID), BNPL, Loans, DAO Governance, Token Vault
- **API:** 32-endpoint HTTP client with Bearer JWT auth

Client documentation: [`docs/client/`](docs/client/)

---

## Chainlink CRE Integration

### What Is CRE?

Chainlink CRE (Compute Runtime Environment) is a decentralised compute platform that runs custom WASM binaries on the Chainlink DON (Decentralised Oracle Network). Workflows are triggered by on-chain events (EVM log filters) or cron schedules and can read/write to smart contracts and call external HTTP APIs.

### Why CRE?

Optimus uses CRE to move **all autonomous protocol logic off the backend** and into a trustless, decentralised execution environment. Instead of running cron jobs on a centralised server or relying on keeper bots, Optimus delegates critical financial operations to CRE workflows:

1. **Risk Scoring** — Every BNPL payment, loan payment, late fee, and default adjusts the user's on-chain risk profile via CRE. The Flutter app and backend never calculate or set risk scores directly.

2. **Late Fee Enforcement** — The `bnpl_late_fee` workflow runs every 30 minutes, scans active BNPL arrangements, and applies late fees + risk penalties autonomously.

3. **Default Detection** — The `loan_default_monitor` workflow runs hourly, checks whether loans have exceeded their duration, and calls `markDefaulted` on-chain — triggering a -2000 risk score penalty.

4. **Proposal Finalization** — The `dao_proposal_finalizer` workflow runs hourly and auto-finalizes DAO proposals whose voting period has ended.

5. **State Reconciliation** — The `state_reconciler` workflow runs every 2 hours and cross-checks on-chain state against the backend database, notifying if discrepancies exist.

6. **Treasury Monitoring** — The `treasury_monitor` workflow runs every 6 hours and alerts if a DAO treasury balance drops below threshold.

7. **Event Enrichment** — Event-triggered workflows (10 of 14) fire on specific contract events (e.g., `BNPLCreated`, `LoanApproved`, `VoteCast`) and enrich the backend with contextual data — risk scores, accrued interest, arrangement details — that the backend cannot compute on its own.

### CRE Architecture

```
  On-Chain Events              CRE DON                    Backend API
  ──────────────          ─────────────────          ──────────────────
  BNPLCreated ───────────→ bnpl_created ──────────→ POST /api/bnpl/created
  BNPLPaymentMade ───────→ bnpl_payment ──────────→ POST /api/bnpl/payment
  BNPLCompleted ─────────→ bnpl_completed ────────→ POST /api/bnpl/completed
  ProposalOpened ────────→ dao_proposal_opened ───→ POST /api/dao/proposals/opened
  VoteCast ──────────────→ dao_vote_cast ─────────→ POST /api/dao/votes
  LoanCreated ───────────→ loan_created ──────────→ POST /api/loans/created
  LoanApproved ──────────→ loan_approved ─────────→ POST /api/loans/approved
  PaymentMade ───────────→ loan_payment ──────────→ POST /api/loans/payment
  RiskProfileUpdated ────→ risk_profile_updater ──→ POST /api/did/risk-updated

  Cron (30min) ──────────→ bnpl_late_fee ─────────→ applyLateFee() on-chain
  Cron (1hr) ────────────→ dao_proposal_finalizer → finalizeProposal() on-chain
  Cron (1hr) ────────────→ loan_default_monitor ──→ markDefaulted() on-chain
  Cron (2hr) ────────────→ state_reconciler ──────→ POST /api/reconcile/*
  Cron (6hr) ────────────→ treasury_monitor ──────→ POST /api/treasury/alert
```

### SDK & Build

- **SDK:** `github.com/smartcontractkit/cre-sdk-go v1.1.0`
- **Build target:** `GOOS=wasip1 GOARCH=wasm go build ./...`
- **Shared config:** [`workflows/shared/`](workflows/shared/) — contract addresses, RPC URL, API base URL
- **Deployment spec:** [`workflows/project.yaml`](workflows/project.yaml)

---

## CRE Workflow Catalogue

### Event-Triggered Workflows (10)

| Workflow | Trigger Event | Contract Reads | Contract Writes | Risk Δ | Source |
|----------|---------------|----------------|-----------------|--------|--------|
| `bnpl_created` | `BNPLCreated` | getArrangement, getBnplTerms | — | — | [`workflows/bnpl_created/`](workflows/bnpl_created/) |
| `bnpl_payment` | `BNPLPaymentMade` | getArrangement, getRiskProfileScore | updateRiskProfile | +50 / +500 | [`workflows/bnpl_payment/`](workflows/bnpl_payment/) |
| `bnpl_completed` | `BNPLCompleted` | getArrangement, getTreasuryBalance | — | — | [`workflows/bnpl_completed/`](workflows/bnpl_completed/) |
| `dao_proposal_opened` | `ProposalOpened` | — | — | — | [`workflows/dao_proposal_opened/`](workflows/dao_proposal_opened/) |
| `dao_vote_cast` | `VoteCast` | — | — | — | [`workflows/dao_vote_cast/`](workflows/dao_vote_cast/) |
| `loan_created` | `LoanCreated` | getRiskProfileScore | — | — | [`workflows/loan_created/`](workflows/loan_created/) |
| `loan_approved` | `LoanApproved` | getLoan, getAmountOwed, getAccruedInterest | — | — | [`workflows/loan_approved/`](workflows/loan_approved/) |
| `loan_payment` | `PaymentMade` | getRiskProfileScore, getAccruedInterest | updateRiskProfile | +100 / +700 | [`workflows/loan_payment/`](workflows/loan_payment/) |
| `risk_profile_updater` | `RiskProfileUpdated` | — | — | — | [`workflows/risk_profile_updater/`](workflows/risk_profile_updater/) |

### Cron-Triggered Workflows (4)

| Workflow | Schedule | Actions | Risk Δ | Source |
|----------|----------|---------|--------|--------|
| `bnpl_late_fee` | Every 30 min | Reads arrangement + terms → calls `applyLateFee()` + `updateRiskProfile()` | -300 | [`workflows/bnpl_late_fee/`](workflows/bnpl_late_fee/) |
| `dao_proposal_finalizer` | Every 1 hr | Scans open proposals → calls `finalizeProposal()` | — | [`workflows/dao_proposal_finalizer/`](workflows/dao_proposal_finalizer/) |
| `loan_default_monitor` | Every 1 hr | Checks loan duration → calls `markDefaulted()` + `updateRiskProfile()` | -2000 | [`workflows/loan_default_monitor/`](workflows/loan_default_monitor/) |
| `state_reconciler` | Every 2 hr | Cross-checks on-chain vs DB state → notifies backend of discrepancies | — | [`workflows/state_reconciler/`](workflows/state_reconciler/) |
| `treasury_monitor` | Every 6 hr | Reads DAO treasury balance → alerts if below threshold | — | [`workflows/treasury_monitor/`](workflows/treasury_monitor/) |

### Risk Score Summary

| Event | Points | Direction | Workflow |
|-------|--------|-----------|----------|
| BNPL payment (on-time) | +50 | ↑ | `bnpl_payment` |
| BNPL final payment | +500 | ↑ | `bnpl_payment` |
| Loan payment (on-time) | +100 | ↑ | `loan_payment` |
| Loan fully repaid | +700 | ↑ | `loan_payment` |
| BNPL late fee applied | -300 | ↓ | `bnpl_late_fee` |
| Loan defaulted | -2000 | ↓ | `loan_default_monitor` |

Workflow documentation: [`docs/workflows/`](docs/workflows/)

---

## How CRE Is Used

The relationship between the app, backend, contracts, and CRE workflows follows a clear pattern:

### 1. User Action → Contract Event → CRE Reaction

When a user performs an action in the Flutter app (e.g., makes a BNPL payment), the request flows through the Go backend to the smart contract. The contract emits an event (e.g., `BNPLPaymentMade`). A CRE event-triggered workflow picks up this event, reads additional on-chain data (risk scores, arrangement details), performs computations (risk score adjustments), optionally writes back to the contract (`updateRiskProfile`), and notifies the backend API with enriched data.

```
User (Flutter) → Backend API → Smart Contract → Event Log
                                                     ↓
                                              CRE DON picks up event
                                                     ↓
                                              Workflow reads chain data
                                                     ↓
                                              Workflow writes to chain (risk update)
                                                     ↓
                                              Workflow notifies backend
```

### 2. Autonomous Cron Operations

Four CRE cron workflows run independently of user actions to enforce protocol rules:

- **Late fees** (`bnpl_late_fee`): Scans every 30 minutes for overdue BNPL installments, applies fees, and penalises risk scores. No user or backend involvement.
- **Defaults** (`loan_default_monitor`): Checks hourly whether loans have exceeded their duration. Auto-marks them as defaulted on-chain with a severe risk penalty (-2000).
- **Governance** (`dao_proposal_finalizer`): Finalises DAO proposals whose voting window has closed. Ensures governance moves forward without manual intervention.
- **Integrity** (`state_reconciler`, `treasury_monitor`): Periodically verify that the off-chain database matches on-chain reality and that treasury balances remain healthy.

### 3. What the Flutter App Does NOT Do

Because CRE handles autonomous operations, the Flutter app intentionally omits certain actions:

| Action | Why it's not in the app | CRE Workflow |
|--------|------------------------|--------------|
| Apply late fees | Automated on schedule | [`bnpl_late_fee`](workflows/bnpl_late_fee/) |
| Mark loans as defaulted | Automated on schedule | [`loan_default_monitor`](workflows/loan_default_monitor/) |
| Update risk scores | Triggered by contract events | [`bnpl_payment`](workflows/bnpl_payment/), [`loan_payment`](workflows/loan_payment/) |
| Activate BNPL arrangements | Auto-activates on first payment | (contract logic) |
| Finalize DAO proposals | Automated on schedule | [`dao_proposal_finalizer`](workflows/dao_proposal_finalizer/) |

This separation ensures that critical financial logic runs in a decentralised, tamper-proof environment rather than depending on a centralised backend or user action.

---

## Deployment

See [`DEPLOY.md`](DEPLOY.md) for the full deployment guide.

### Contracts

Deployed via Foundry (`forge script`) from the `contracts/` directory.

### Backend

Go binary deployed to EC2 as a systemd service:

```bash
# One-time EC2 setup (Go, PostgreSQL, systemd unit, env template)
./scripts/setup-ec2.sh ubuntu@13.60.166.148 optimus.pem

# Deploy / redeploy
./scripts/deploy-ec2.sh ubuntu@13.60.166.148 optimus.pem
```

### CRE Workflows

Built as WASM and deployed to the Chainlink DON (not EC2):

```bash
cd workflows
GOOS=wasip1 GOARCH=wasm go build ./...
# Deploy via CRE CLI per workflows/project.yaml
```

---

## Documentation

Full documentation with PlantUML diagrams:

| Area | Path | Contents |
|------|------|----------|
| Contracts | [`docs/contracts/`](docs/contracts/) | Per-contract specifications, event schemas, storage layout |
| Backend Controllers | [`docs/protocol/`](docs/protocol/) | Route tables, request/response schemas, data flow diagrams |
| Flutter Client | [`docs/client/`](docs/client/) | Screen docs, auth flow, provider architecture |
| CRE Workflows | [`docs/workflows/`](docs/workflows/) | Trigger types, risk adjustments, shared config, per-workflow specs |

---

## License

See [LICENSE](LICENSE).
