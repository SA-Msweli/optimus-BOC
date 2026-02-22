# Optimus Protocol Backend

Go backend for the Optimus Buy-Own-Credit (BOC) protocol. Provides an HTTP
API that wraps on-chain smart-contracts via `go-ethereum` bindings and
persists minimal CRE-relevant data in PostgreSQL.

## Architecture

The project follows SOLID principles with a structure similar to .NET MVC:

```
protocol/
├── bindings/          # Auto-generated Go bindings from Solidity ABIs
├── config.go          # Environment configuration (LoadConfig)
├── controllers/       # HTTP controllers (one package per domain)
│   ├── bnpl/          # BNPL arrangement endpoints
│   ├── dao/           # DAO management endpoints
│   ├── did/           # DID identity endpoints
│   ├── loan/          # Loan management endpoints
│   └── tokenvault/    # Token vault endpoints
├── db/migrations/     # SQL migration files
├── eth/               # Ethereum transactor helper
├── main.go            # Composition root — wires services, controllers, router
├── models/            # Data models (Arrangement, DAO, DIDProfile, Loan)
├── services/          # Business logic (one sub-package per domain)
│   ├── bnpl/          # IBNPLService interface + Service implementation
│   ├── dao/           # IDAOService interface + Service implementation
│   ├── did/           # IDid interface + Did implementation
│   ├── loan/          # ILoanService interface + Service implementation
│   └── tokenvault/    # ITokenVaultService interface + Service implementation
└── store/             # PostgreSQL persistence layer
```

Each domain follows the pattern:
- **Interface** (`interface.go`) — defines the contract (Dependency Inversion)
- **Service** (`*_service.go`) — implements the interface against real chain bindings
- **Controller** (`*.go`) — HTTP handlers that depend only on interfaces

## Prerequisites

- Go 1.24+
- PostgreSQL 15+
- `abigen` (from go-ethereum)
  ```bash
  go install github.com/ethereum/go-ethereum/cmd/abigen@latest
  ```

## Configuration

All configuration is provided via environment variables. Use `set -a && . .env && set +a` before running.

| Variable                 | Required | Description                              |
|--------------------------|----------|------------------------------------------|
| `PORT`                   | No       | HTTP listen port (default `8000`)        |
| `DATABASE_URL`           | Yes      | PostgreSQL connection string             |
| `CHAIN_RPC_URL`          | Yes      | Ethereum JSON-RPC endpoint               |
| `PRIVATE_KEY`            | Yes      | Hex-encoded private key for signing txs  |
| `BNPL_MANAGER_ADDRESS`   | Yes      | Deployed BNPLManager contract address    |
| `DAO_MANAGER_ADDRESS`    | Yes      | Deployed DAOManager contract address     |
| `DID_REGISTRY_ADDRESS`   | Yes      | Deployed DIDRegistry contract address    |
| `LOAN_MANAGER_ADDRESS`   | Yes      | Deployed LoanManager contract address    |
| `TOKEN_VAULT_ADDRESS`    | Yes      | Deployed TokenVault contract address     |
| `PRIVY_APP_ID`           | No       | Privy application ID                     |
| `PRIVY_APP_SECRET`       | No       | Privy application secret for auth        |
| `PRIVY_JWKS`             | No       | Privy JWKS endpoint URL                  |

## Generating ABI Bindings

```bash
cd protocol
./generate_bindings.sh
```

## Building and Running

```bash
cd protocol
set -a && . .env && set +a
go run ./...
```

Or build a binary:

```bash
go build -o optimus .
./optimus
```

## HTTP API Endpoints

All endpoints return JSON. Write endpoints return `{"tx": "0x..."}`.

### General

| Method | Path      | Description          |
|--------|-----------|----------------------|
| GET    | `/health` | Health check → `OK`  |
| GET    | `/`       | Greeting message     |

### DID (`/did`)

| Method | Path                | Body / Params                                   | Description                          |
|--------|---------------------|-------------------------------------------------|--------------------------------------|
| POST   | `/did`              | `{"owner":"0x..."}`                             | Create DID on-chain                  |
| GET    | `/did/{owner}`      | —                                               | Check if DID exists                  |
| POST   | `/did/{owner}/link` | `{"hash":"0x..."}`                              | Link Privy credential hash           |
| GET    | `/did/{owner}/privy`| —                                               | Get stored Privy hash                |
| POST   | `/did/{owner}/risk` | `{"new_score":"100","risk_profile_hash":"0x..."}`| Update risk profile on-chain         |
| GET    | `/did/{owner}/risk` | —                                               | Get risk profile score               |

### BNPL (`/bnpl`)

| Method | Path                                  | Body / Params                                                                 | Description                 |
|--------|---------------------------------------|-------------------------------------------------------------------------------|-----------------------------|
| POST   | `/bnpl/arrangements`                  | `{"dao_id":"1","recipient":"0x...","total_amount":"1000","start_timestamp":0,"interval_seconds":86400}` | Create arrangement |
| GET    | `/bnpl/arrangements/{id}`             | —                                                                             | Get arrangement from chain  |
| POST   | `/bnpl/arrangements/{id}/payment`     | `{"installment":1}`                                                           | Make installment payment    |
| POST   | `/bnpl/arrangements/{id}/activate`    | —                                                                             | Activate arrangement        |
| POST   | `/bnpl/arrangements/{id}/latefee`     | `{"installment":1}`                                                           | Apply late fee              |
| POST   | `/bnpl/arrangements/{id}/reschedule`  | `{"new_start_timestamp":0,"new_interval_seconds":86400}`                      | Reschedule arrangement      |

### DAO (`/dao`)

| Method | Path                          | Body / Params                                                                                                      | Description                |
|--------|-------------------------------|--------------------------------------------------------------------------------------------------------------------|----------------------------|
| POST   | `/dao`                        | `{"goal":1,"voting_period_days":7}`                                                                                | Create a DAO               |
| POST   | `/dao/{id}/join`              | `{"investment":"1000000000000000000"}`                                                                             | Join a DAO                 |
| POST   | `/dao/{daoId}/propose`        | `{"data":"0x..."}`                                                                                                 | Submit proposal (members)  |
| POST   | `/dao/proposal/{id}/vote`     | `{"dao_id":"1","support":true}`                                                                                    | Vote on proposal (members) |
| POST   | `/dao/proposal/{id}/finalize` | —                                                                                                                  | Finalize proposal          |
| POST   | `/dao/proposal/{id}/execute`  | —                                                                                                                  | Execute proposal           |
| POST   | `/dao/{id}/bnpl-terms`        | `{"num_installments":"12","min_days":"7","max_days":"30","late_fee_bps":"500","grace_days":"3","reschedule_allowed":true,"min_down_bps":"1000"}` | Set BNPL terms |
| GET    | `/dao/{id}/bnpl-terms`        | —                                                                                                                  | Get BNPL terms             |
| GET    | `/dao/{id}/treasury`          | —                                                                                                                  | Get treasury balance       |
| POST   | `/dao/{id}/treasury/credit`   | `{"amount":"1000"}`                                                                                                | Credit treasury            |

### Loan (`/loan`)

| Method | Path                     | Body / Params                                                                                         | Description              |
|--------|--------------------------|-------------------------------------------------------------------------------------------------------|--------------------------|
| POST   | `/loan`                  | `{"borrower":"0x...","dao_id":"1","principal":"1000","interest_rate_bps":"500","duration_seconds":"86400"}` | Create loan         |
| POST   | `/loan/{id}/approve`     | —                                                                                                     | Approve loan             |
| POST   | `/loan/{id}/payment`     | —                                                                                                     | Make loan payment        |
| POST   | `/loan/{id}/default`     | —                                                                                                     | Mark loan as defaulted   |
| GET    | `/loan/{id}`             | —                                                                                                     | Get loan details         |
| GET    | `/loan/{id}/interest`    | —                                                                                                     | Get accrued interest     |
| GET    | `/loan/{id}/owed`        | —                                                                                                     | Get total amount owed    |

### Token Vault (`/vault`)

| Method | Path                    | Body / Params                      | Description                |
|--------|-------------------------|------------------------------------|----------------------------|
| POST   | `/vault/deposit`        | `{"token":"0x...","amount":"1000"}`| Deposit tokens             |
| POST   | `/vault/withdraw`       | `{"token":"0x...","amount":"1000"}`| Withdraw tokens            |
| GET    | `/vault/balance/{token}`| —                                  | Get balance for token addr |

## Database

The schema is defined in `db/migrations/001_init.sql`. Apply it manually:

```bash
psql $DATABASE_URL -f db/migrations/001_init.sql
```

Tables: `arrangements`, `loans`, `daos`, `did_profiles`.

Only minimal data needed for off-chain CRE workflows is persisted; canonical
state lives on-chain.
