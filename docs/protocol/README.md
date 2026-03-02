# Optimus Backend (Go Protocol Server)

**Source:** `protocol/`  
**Language:** Go 1.23+ (go-ethereum, chi, pgxpool)  
**Server:** `ubuntu@13.60.166.148` (EC2)

## Architecture Overview

The backend follows a **Controller → Service → Contract** layered architecture with clear dependency inversion:

```
┌─────────────────────────────────────────────────────┐
│                    HTTP Layer                        │
│   chi Router + Privy JWT Middleware                  │
├─────────┬──────────┬──────────┬──────────┬──────────┤
│ DID     │ BNPL     │ DAO      │ Loan     │ Vault    │
│ Ctrl    │ Ctrl     │ Ctrl     │ Ctrl     │ Ctrl     │
├─────────┼──────────┼──────────┼──────────┼──────────┤
│ IDid    │ IBNPL    │ IDAO     │ ILoan    │ IVault   │
│ Svc     │ Svc      │ Svc      │ Svc      │ Svc      │
├─────────┴──────────┴──────────┴──────────┴──────────┤
│     go-ethereum bindings → Sepolia (chain 11155111)  │
├─────────────────────────────────────────────────────┤
│     PostgreSQL (pgxpool) — off-chain Store           │
└─────────────────────────────────────────────────────┘
```

## Key Components

| Component         | Path                      | Purpose                              |
|-------------------|---------------------------|--------------------------------------|
| Entry point       | `main.go`                 | Wires dependencies, starts server    |
| Config            | `config.go`               | Reads env vars                       |
| Controllers (5)   | `controllers/{module}/`   | HTTP handlers, request parsing       |
| Services (5)      | `services/{module}/`      | Business logic, contract interactions|
| Models (4)        | `models/`                 | Off-chain data structures            |
| Store             | `store/store.go`          | PostgreSQL upsert helpers            |
| Middleware        | `middleware/privy_auth.go` | Privy RS256 JWT verification        |
| Transactor        | `eth/transactor.go`       | Ethereum signing from PRIVATE_KEY    |

## Environment Variables

| Variable               | Required | Description                        |
|------------------------|----------|------------------------------------|
| `PORT`                 | No (8000)| HTTP listen port                   |
| `DATABASE_URL`         | Yes      | PostgreSQL connection string       |
| `CHAIN_RPC_URL`        | Yes      | Sepolia RPC endpoint               |
| `DID_REGISTRY_ADDRESS` | Yes      | DIDRegistry contract address       |
| `BNPL_MANAGER_ADDRESS` | Yes      | BNPLManager contract address       |
| `LOAN_MANAGER_ADDRESS` | Yes      | LoanManager contract address       |
| `DAO_MANAGER_ADDRESS`  | Yes      | DAOManager contract address        |
| `TOKEN_VAULT_ADDRESS`  | Yes      | TokenVault contract address        |
| `PRIVATE_KEY`          | Yes      | Hex-encoded Ethereum private key   |
| `PRIVY_APP_ID`         | No       | Privy dashboard app identifier     |
| `PRIVY_APP_SECRET`     | No       | Privy server-side API secret       |
| `PRIVY_JWKS`           | No       | JWKS endpoint for JWT verification |

## Authentication Flow

```plantuml
@startuml Backend_Auth
actor "Flutter App" as App
participant "Privy SDK" as Privy
participant "chi Router" as Router
participant "PrivyAuth Middleware" as MW
participant "JWKS Endpoint" as JWKS
participant "Controller" as Ctrl

App -> Privy : login (email / social / wallet)
Privy --> App : accessToken (RS256 JWT)

App -> Router : POST /did + Authorization: Bearer <token>
Router -> MW : intercept
MW -> MW : decode JWT header (get kid)
MW -> JWKS : fetch RSA public keys (cached 1hr)
JWKS --> MW : JWKS key set
MW -> MW : verify RS256 signature
MW -> MW : validate exp, iss="privy.io", aud=appId
MW -> MW : extract sub (did:privy:...) + wallet_address
MW --> Ctrl : request with context (userID, wallet)
Ctrl -> Ctrl : process request
Ctrl --> App : JSON response

@enduml
```

## Request Flow (General)

```plantuml
@startuml Backend_Request_Flow
skinparam actorStyle awesome

actor "Mobile App" as App
participant "chi Router" as R
participant "Privy Middleware" as MW
participant "Controller" as C
participant "Service (Interface)" as S
participant "go-ethereum Binding" as Bind
participant "Sepolia RPC" as Chain
database "PostgreSQL" as DB

== Write Operation (e.g. createLoan) ==
App -> R : POST /loan + JWT
R -> MW : authenticate
MW --> C : authenticated context
C -> C : parse JSON body
C -> S : CreateLoan(auth, borrower, daoId, principal, rate, duration)
S -> Bind : contract.CreateLoan(opts, ...)
Bind -> Chain : eth_sendRawTransaction
Chain --> Bind : tx hash
Bind --> S : *types.Transaction
S --> C : tx
C -> S : WaitForLoanID(ctx, tx)
S -> Chain : eth_getTransactionReceipt
Chain --> S : receipt + logs
S -> S : parse LoanCreated event → loanID
S --> C : loanID
C -> DB : store.SaveLoan(ctx, Loan{...})
C --> App : {"tx": "0x...", "loan_id": 42}

== Read Operation (e.g. getLoan) ==
App -> R : GET /loan/42 + JWT
R -> MW : authenticate
MW --> C : authenticated context
C -> S : GetLoan(ctx, loanId)
S -> Bind : contract.Loans(&callOpts, id)
Bind -> Chain : eth_call
Chain --> Bind : on-chain data
Bind --> S : struct
S --> C : Loan struct
C --> App : JSON response

@enduml
```

## API Endpoints Overview

See individual controller documentation for request/response schemas:

- [DID Controller](controllers/did.md) — Identity management
- [BNPL Controller](controllers/bnpl.md) — Buy-Now-Pay-Later arrangements
- [DAO Controller](controllers/dao.md) — DAO lifecycle + governance
- [Loan Controller](controllers/loan.md) — Loan management
- [TokenVault Controller](controllers/tokenvault.md) — ERC-20 vault

## Database Schema

The backend persists minimal off-chain copies of on-chain state for CRE workflow fast lookups:

| Table          | Primary Key        | Purpose                      |
|----------------|--------------------|------------------------------|
| `arrangements` | `arrangement_id`   | BNPL arrangements            |
| `daos`         | `dao_id`           | DAO metadata                 |
| `did_profiles` | `owner`            | DID risk profiles            |
| `loans`        | `loan_id`          | Loan records                 |

All tables use **upsert** patterns (`ON CONFLICT ... DO UPDATE`).
