# Optimus Protocol — System Documentation

> Generated from the implemented codebase targeting **EVM / Sepolia** (chain 11155111).

## Documentation Index

| Layer | Path | Contents |
|-------|------|----------|
| Smart Contracts | [docs/contracts/](contracts/README.md) | 5 Solidity contracts with roles, storage, functions, events |
| Backend (Go) | [docs/protocol/](protocol/README.md) | Chi router, Privy JWT auth, 5 controller groups, PostgreSQL store |
| CRE Workflows | [docs/workflows/](workflows/README.md) | 14 Chainlink workflows (10 event-triggered, 4 cron) |
| Mobile App (Flutter) | [docs/client/](client/README.md) | 5 screens, API client (31 endpoints), Privy auth |

## Deployed Addresses (Sepolia)

| Contract | Address |
|----------|---------|
| DIDRegistry | `0x0E9D8959bCD99e7AFD7C693e51781058A998b756` |
| BNPLManager | `0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba` |
| LoanManager | `0xbB0D4067488edf4a007822407e2486412dC8D39D` |
| DAOManager | `0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b` |
| TokenVault | `0x4C704D51fc47cfe582F8c5477de3AE398B344907` |
| Deployer | `0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780` |

## Architecture Overview

```plantuml
@startuml Architecture_Overview
skinparam componentStyle rectangle

package "Mobile App (Flutter)" as FE {
  [AuthService\n(Privy OTP)] as Auth
  [ApiClient\n(31 endpoints)] as AC
  [5 Domain Services] as DS
  [5 Screens] as Screens
}

package "Backend (Go + Chi)" as BE {
  [Privy JWT Middleware] as MW
  [5 Controllers] as Ctrl
  [5 Services] as Svc
  [Transactor\n(PRIVATE_KEY)] as TX
  [PostgreSQL Store] as DB
}

package "Sepolia (EVM)" as Chain {
  [DIDRegistry] as DID
  [BNPLManager] as BNPL
  [LoanManager] as Loan
  [DAOManager] as DAO
  [TokenVault] as Vault
}

package "Chainlink CRE" as CRE {
  [10 Event Workflows] as EW
  [4 Cron Workflows] as CW
}

Screens --> DS
DS --> AC
AC --> MW : HTTPS + JWT
MW --> Ctrl
Ctrl --> Svc
Svc --> TX
TX --> Chain
Svc --> DB : upsert models

Chain ..> CRE : contract events
CRE --> DID : updateRiskProfile
CRE --> BNPL : applyLateFee
CRE --> DAO : finalizeProposal
CRE ..> BE : HTTP notifications (future)

@enduml
```

## End-to-End Data Flows

### 1. Identity (DID) Flow

```plantuml
@startuml E2E_DID
actor User
participant "DID Screen" as FE
participant "Backend" as BE
participant "DIDRegistry" as DID
participant "CRE: risk_profile\n_updater" as CRE

User -> FE : enter wallet + metadata
FE -> BE : POST /did/ {wallet, metadata_hash}
BE -> DID : createDID(wallet, metadataHash)
DID --> BE : tx (DIDCreated event)
BE -> BE : store DIDProfile (risk_score=500)
BE --> FE : {tx_hash}

== Privy Link ==
User -> FE : tap "Link Privy"
FE -> BE : POST /did/link-privy {wallet, credential_hash}
BE -> DID : linkPrivyCredential(wallet, hash)
DID --> BE : tx (PrivyCredentialLinked)

== Risk Updates (via other domain events) ==
CRE -> DID : updateRiskProfile(wallet, newScore, profileHash)
note right: Triggered by loan/BNPL events\nadjusting risk score

@enduml
```

### 2. BNPL Flow (Full Lifecycle)

```plantuml
@startuml E2E_BNPL
actor User
participant "BNPL Screen" as FE
participant "Backend" as BE
participant "BNPLManager" as BNPL
participant "DAOManager" as DAO
participant "DIDRegistry" as DID

== Creation ==
User -> FE : fill borrower, daoId, installments, tap "Create"
FE -> BE : POST /bnpl/ {borrower, dao_id, installment_amounts[]}
BE -> BNPL : createBNPL(borrower, daoId, amounts[])
BNPL -> DAO : getBnplTerms(daoId)
note right: Validates numInstallments,\nmin/max interval, downPayment
BNPL --> BE : tx (BNPLCreated event)
BE -> BE : WaitForBNPLID(receipt) → bnplId
BE -> BE : SaveArrangement(bnplId, ...)
BE --> FE : {tx_hash, bnpl_id}

note over BNPL : CRE bnpl_created workflow\nadjusts risk -0 (log only)

== Payments ==
User -> FE : enter BNPL ID + amount, tap "Pay"
FE -> BE : POST /bnpl/{id}/pay {amount}
BE -> BNPL : makePayment(bnplId, amount)
BNPL -> BNPL : auto-ACTIVE on 1st payment
BNPL -> BNPL : auto-COMPLETED if all paid
BNPL --> BE : tx

note over BNPL : CRE bnpl_payment\nrisk_score += 50

alt All installments paid
  note over BNPL : CRE bnpl_completed\nrisk_score += 500
end

== Late Fees (Cron 30min) ==
note over BNPL : CRE bnpl_late_fee\nchecks overdue, calls applyLateFee()
BNPL -> DAO : creditTreasury(daoId, fee)
note over BNPL : risk_score -= 300

@enduml
```

### 3. Loan Flow (Full Lifecycle)

```plantuml
@startuml E2E_Loan
actor User
participant "Loan Screen" as FE
participant "Backend" as BE
participant "LoanManager" as LM
participant "DIDRegistry" as DID

== Creation ==
User -> FE : fill borrower, principal, rate, tap "Create"
FE -> BE : POST /loan/ {borrower, principal, interest_rate_bps}
BE -> LM : createLoan(borrower, principal, rateBps)
LM --> BE : tx (LoanCreated event)
BE -> BE : WaitForLoanID(receipt) → loanId
BE -> BE : SaveLoan(loanId, ...)
BE --> FE : {tx_hash, loan_id}

note over LM : CRE loan_created (log only)

== Approval ==
User -> FE : enter loan ID, tap "Approve"
FE -> BE : POST /loan/{id}/approve
BE -> LM : approveLoan(loanId)
LM --> BE : tx (LoanApproved)
BE -> BE : UpdateLoanStatus(loanId, APPROVED)
BE --> FE : {tx_hash}

note over LM : CRE loan_approved\nrisk_score unchanged

== Payments ==
User -> FE : enter loan ID + amount, tap "Pay"
FE -> BE : POST /loan/{id}/pay {amount}
BE -> LM : makePayment(loanId, amount)
LM -> LM : auto-REPAID if amount >= amountOwed
LM --> BE : tx (PaymentMade)

note over LM : CRE loan_payment\nrisk_score += 100
alt Full repayment
  note over LM : risk_score += 700 (payoff bonus)
end

== Default Monitor (Cron 1hr) ==
note over LM : CRE loan_default_monitor\nchecks amountOwed > threshold\ncalls markDefaulted(loanId)
LM --> BE : LoanDefaulted event

note over LM : risk_score -= 2000

@enduml
```

### 4. DAO Governance Flow

```plantuml
@startuml E2E_DAO
actor User
participant "DAO Screen" as FE
participant "Backend" as BE
participant "DAOManager" as DM
participant "TokenVault" as TV
participant "DIDRegistry" as DID

== Create DAO ==
User -> FE : name, goal(0-2), period
FE -> BE : POST /dao/ {goal, voting_period_days}
BE -> DM : createDAO(from, goal, days)
DM --> BE : tx
BE -> BE : WaitForDaoID → daoId
BE -> BE : SaveDAO(daoId, ...)
BE --> FE : {tx, dao_id}

== Register Member ==
User -> FE : DAO ID, stake (wei)
FE -> BE : POST /dao/{id}/join {investment}
BE -> DM : joinDAO(daoId, from, investment)
DM -> DM : members[daoId][addr] = investment
DM -> DM : memberCount++, treasury += investment
DM --> BE : tx (MemberJoined)

== Proposal Lifecycle ==
User -> FE : DAO ID, proposal data (hex)
FE -> BE : POST /dao/{daoId}/propose {data}
BE -> BE : IsMember check (scan MemberJoined events)
BE -> DM : propose(daoId, data)
DM -> DM : expiry = now + votingPeriod
DM --> BE : tx (ProposalOpened)

note over DM : CRE dao_proposal_opened (log)

User -> FE : proposal ID, FOR/AGAINST
FE -> BE : POST /dao/proposal/{id}/vote {support}
BE -> DM : vote(proposalId, support)
DM -> DM : weight = memberInvestment (min 1)
DM --> BE : tx (VoteCast)

note over DM : CRE dao_vote_cast (log)

== Finalization ==
alt Manual
  User -> FE : tap "Finalize"
  FE -> BE : POST /dao/proposal/{id}/finalize
  BE -> DM : finalizeProposal(id)
else Cron (1hr)
  note over DM : CRE dao_proposal_finalizer\nauto-finalizes expired proposals
end
DM -> DM : approved = yesWeight > noWeight

== Execution ==
User -> FE : tap "Execute"
FE -> BE : POST /dao/proposal/{id}/execute
BE -> DM : executeProposal(id)
alt proposal data decodes to (token, amount, recipient)
  DM -> TV : withdraw(token, amount)
  DM -> DM : transfer to recipient
end

== BNPL Terms ==
User -> FE : 7 fields via "Terms" tab
FE -> BE : POST /dao/{id}/bnpl-terms {7 params}
BE -> DM : setBnplTerms(daoId, ...)
DM --> BE : tx (BnplTermsUpdated)

== Treasury ==
User -> FE : amount
FE -> BE : POST /dao/{id}/treasury/credit {amount}
BE -> DM : creditTreasury(daoId, amount)
note right: Requires TREASURY_FUNDER or ADMIN

@enduml
```

### 5. Token Vault Flow

```plantuml
@startuml E2E_Vault
actor User
participant "Vault Screen" as FE
participant "Backend" as BE
participant "TokenVault" as TV

User -> FE : token (pre-filled), amount
FE -> BE : POST /vault/deposit {token, amount}
BE -> TV : deposit(token, amount)
TV -> TV : transferFrom(sender, vault, amount)
TV --> BE : tx
BE --> FE : {tx_hash}

User -> FE : tap "Check Balance"
FE -> BE : GET /vault/balance?token=...
BE -> TV : getBalance(token) [view]
TV --> BE : uint256
BE --> FE : {balance}

User -> FE : amount, tap "Withdraw"
FE -> BE : POST /vault/withdraw {token, amount}
BE -> TV : withdraw(token, amount)
note right: VAULT_MANAGER or ADMIN
TV --> BE : tx
BE --> FE : {tx_hash}

@enduml
```

## Risk Score System

All risk scores flow through `DIDRegistry.updateRiskProfile()` via CRE workflows.

| Event | Workflow | Score Δ |
|-------|----------|---------|
| BNPL payment made | `bnpl_payment` | **+50** |
| BNPL completed | `bnpl_completed` | **+500** |
| BNPL late fee applied | `bnpl_late_fee` | **−300** |
| Loan payment made | `loan_payment` | **+100** |
| Loan fully repaid | `loan_payment` | **+700** |
| Loan defaulted | `loan_default_monitor` | **−2000** |

Initial score: **500** (set on DID creation in backend store).

## Event → Workflow Mapping

```plantuml
@startuml Event_Workflow_Map
skinparam componentStyle rectangle

package "Contract Events" as Events {
  [BNPLCreated] as E1
  [BNPLPaymentMade] as E2
  [BNPLCompleted] as E3
  [LoanCreated] as E4
  [LoanApproved] as E5
  [PaymentMade] as E6
  [ProposalOpened] as E7
  [VoteCast] as E8
}

package "Event-Triggered Workflows" as ETW {
  [bnpl_created] as W1
  [bnpl_payment] as W2
  [bnpl_completed] as W3
  [loan_created] as W4
  [loan_approved] as W5
  [loan_payment] as W6
  [dao_proposal_opened] as W7
  [dao_vote_cast] as W8
  [risk_profile_updater] as W9
}

package "Cron Workflows" as CW {
  [bnpl_late_fee\n(30 min)] as C1
  [dao_proposal_finalizer\n(1 hr)] as C2
  [loan_default_monitor\n(1 hr)] as C3
  [state_reconciler\n(2 hr)] as C4
  [treasury_monitor\n(6 hr)] as C5
}

E1 --> W1
E2 --> W2
E3 --> W3
E4 --> W4
E5 --> W5
E6 --> W6
E7 --> W7
E8 --> W8

W2 --> W9 : triggers
W3 --> W9 : triggers
W6 --> W9 : triggers
C1 --> W9 : triggers
C3 --> W9 : triggers

@enduml
```

## Contract Dependency Graph

```plantuml
@startuml Contract_Dependencies
skinparam componentStyle rectangle

[DIDRegistry] as DID
[BNPLManager] as BNPL
[LoanManager] as Loan
[DAOManager] as DAO
[TokenVault] as Vault

BNPL --> DAO : getBnplTerms()\ncreditTreasury()
DAO --> Vault : withdraw()\nvia executeProposal()
Loan ..> DID : risk checks (via CRE)
BNPL ..> DID : risk checks (via CRE)

note bottom of DID : Standalone\nno contract deps
note bottom of Loan : Standalone\nno contract deps
note bottom of Vault : Standalone\nDAO calls via VAULT_MANAGER role

@enduml
```

## Authentication Flow

```plantuml
@startuml Auth_Flow
actor User
participant "Flutter App" as App
participant "Privy SDK" as Privy
participant "Backend" as BE

User -> App : enter email
App -> Privy : loginWithOTP(email)
Privy --> User : OTP code (email)
User -> App : enter OTP
App -> Privy : verify OTP
Privy --> App : JWT (RS256) + embedded wallet

App -> BE : any request\nAuthorization: Bearer <JWT>
BE -> BE : verify RS256 signature\nvia Privy JWKS
BE -> BE : extract sub (did:privy:...)\nextract wallet_address
BE -> BE : set ctx values
BE --> App : response

@enduml
```

## Backend Request Lifecycle

```plantuml
@startuml Request_Lifecycle
participant "Flutter" as FE
participant "Chi Router" as R
participant "Privy Middleware" as MW
participant "Controller" as C
participant "Service" as S
participant "Transactor" as TX
participant "Sepolia" as Chain
participant "Store (PG)" as DB

FE -> R : HTTP request
R -> MW : route match
MW -> MW : verify JWT
MW -> C : context with wallet + DID
C -> C : parse & validate body
C -> S : call business method
S -> TX : prepare tx
TX -> Chain : eth_sendRawTransaction
Chain --> TX : receipt
TX --> S : receipt + logs
S -> S : parse event logs (IDs, etc.)
S -> DB : upsert model
S --> C : result
C --> FE : JSON response

@enduml
```
