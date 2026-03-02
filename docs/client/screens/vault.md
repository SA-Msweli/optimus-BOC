# Token Vault Screen

**Source:** `client/prime/lib/screens/vault_screen.dart`  
**Service:** `VaultService`  
**Tab:** Vault (index 4)

## UI Elements

| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| ERC-20 Token Address | `_tokenCtrl` | All 3 endpoints |
| Amount (wei) | `_amountCtrl` | `POST /vault/deposit`, `POST /vault/withdraw` |
| Balance Display | read-only | `GET /vault/balance?token=` |

> **Changed:** Token field no longer pre-fills with `AppConfig.tokenVaultAddress` (which was incorrectly filling the vault contract address instead of an ERC-20 token address). Users must enter the specific ERC-20 token contract address they wish to deposit or withdraw.

## Actions

| Button | Method | API Call |
|--------|--------|----------|
| Deposit | `_deposit()` | `POST /vault/deposit {token, amount}` |
| Withdraw | `_withdraw()` | `POST /vault/withdraw {token, amount}` |
| Refresh Balance | `_checkBalance()` | `GET /vault/balance?token=` |

## Data Flow Diagram

```plantuml
@startuml Vault_Screen_Flow
skinparam actorStyle awesome

actor User
participant "Vault Screen" as Scr
participant "VaultService" as Svc
participant "ApiClient" as API
participant "Backend\n/vault" as BE
participant "TokenVault" as TV

== Deposit ==
User -> Scr : enter token + amount, tap "Deposit"
Scr -> Svc : deposit(token, amount)
Svc -> API : POST /vault/deposit {token, amount}
API -> BE : HTTP + JWT
BE -> TV : deposit(token, amount)
TV -> TV : IERC20(token).transferFrom(msg.sender, this, amount)
TV --> BE : tx receipt
BE --> API : {tx_hash}
API --> Scr : success

== Withdraw ==
User -> Scr : enter token + amount, tap "Withdraw"
Scr -> Svc : withdraw(token, amount)
Svc -> API : POST /vault/withdraw {token, amount}
API -> BE : HTTP + JWT
BE -> TV : withdraw(token, amount)
note right: Requires VAULT_MANAGER_ROLE\nor ADMIN
TV -> TV : IERC20(token).transfer(msg.sender, amount)
TV --> BE : tx receipt
BE --> API : {tx_hash}
API --> Scr : success

== Check Balance ==
User -> Scr : tap "Refresh Balance"
Scr -> Svc : getBalance(token)
Svc -> API : GET /vault/balance?token=...
API -> BE : HTTP + JWT
BE -> TV : getBalance(token) [view]
TV --> BE : uint256
BE --> API : {balance}
API --> Scr : display balance

note over TV: DAOManager.executeProposal()\ncan also call withdraw()\nto fund approved proposals

@enduml
```
