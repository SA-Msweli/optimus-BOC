# TokenVault Controller

**Source:** `protocol/controllers/tokenvault/tokenvault.go`  
**Mount:** `/vault` (protected — Privy JWT required)  
**Service:** `services/tokenvault.ITokenVaultService`  
**Contract:** `TokenVault` (`0x4C704D51fc47cfe582F8c5477de3AE398B344907`)

## Routes

| Method | Path                  | Handler      | Description              |
|--------|-----------------------|-------------|--------------------------|
| POST   | `/vault/deposit`      | `deposit`    | Deposit ERC-20 tokens    |
| POST   | `/vault/withdraw`     | `withdraw`   | Withdraw ERC-20 tokens   |
| GET    | `/vault/balance/{token}` | `getBalance` | Get token balance      |

## Request / Response Schemas

### POST `/vault/deposit` — Deposit Tokens

**Request:**
```json
{
  "token": "0x...",
  "amount": "1000000000000000000"
}
```
**Response:**
```json
{ "tx": "0x..." }
```
**Note:** Requires prior ERC-20 approval from caller to vault contract.

---

### POST `/vault/withdraw` — Withdraw Tokens

**Request:**
```json
{
  "token": "0x...",
  "amount": "1000000000000000000"
}
```
**Response:**
```json
{ "tx": "0x..." }
```
**Note:** Caller must have `ADMIN` or `VAULT_MANAGER_ROLE` on the contract.

---

### GET `/vault/balance/{token}` — Get Balance

**Response:**
```json
{ "balance": "5000000000000000000" }
```

## Data Flow Diagram

```plantuml
@startuml Vault_Controller_Flow
skinparam actorStyle awesome

actor "Flutter App\n(Vault Screen)" as App
participant "Vault Controller" as Ctrl
participant "ITokenVaultService" as Svc
participant "TokenVault\nContract" as TV
participant "ERC-20 Token" as Token

== Deposit ==
App -> Ctrl : POST /vault/deposit\n{"token": "0x...", "amount": "1000..."}
Ctrl -> Svc : Deposit(auth, tokenAddr, amount)
Svc -> TV : deposit(token, amount)
TV -> Token : transferFrom(caller, vault, amount)
TV --> Svc : tx
Ctrl --> App : {"tx": "0x..."}

== Withdraw ==
App -> Ctrl : POST /vault/withdraw\n{"token": "0x...", "amount": "1000..."}
Ctrl -> Svc : Withdraw(auth, tokenAddr, amount)
Svc -> TV : withdraw(token, amount)
TV -> Token : transfer(recipient, amount)
TV --> Svc : tx
Ctrl --> App : {"tx": "0x..."}

== Balance Check ==
App -> Ctrl : GET /vault/balance/{token}
Ctrl -> Svc : GetBalance(ctx, tokenAddr)
Svc -> TV : getBalance(token)
TV -> Token : balanceOf(vault)
TV --> Svc : balance
Ctrl --> App : {"balance": "5000..."}

@enduml
```

## No Off-Chain Storage

Unlike other controllers, TokenVault does **not** use the PostgreSQL store — all state is on-chain.
