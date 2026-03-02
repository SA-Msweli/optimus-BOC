# Loan Screen

**Source:** `client/prime/lib/screens/loan_screen.dart`  
**Service:** `LoanService`  
**Tab:** Loans (index 2)

## UI Elements

### Lookup Section
| Element | Controller | API Endpoint |
|---------|-----------|-------------|
| Loan ID | `_idCtrl` | `GET /loan/{id}` + interest + owed |

### Loan Detail (when loaded)
Displays: loanId, borrower, principal (wei + ETH), interestRate %, start, end, amountPaid, status, accruedInterest, totalOwed

Status labels: `PENDING` (0), `APPROVED` (1), `REPAID` (2), `DEFAULTED` (3)

### Action Buttons
| Button | API Endpoint | Notes |
|--------|-------------|-------|
| Approve | `POST /loan/{id}/approve` | Changes status to APPROVED |
| Make Payment | `POST /loan/{id}/payment` | Sends ETH amount; auto-marks REPAID when owed=0 |
| Refresh Amounts | `GET /loan/{id}/interest` + `GET /loan/{id}/owed` | Updates display |

### Payment Section
| Field | Controller | Purpose |
|-------|-----------|---------|
| Amount (wei) | `_payAmountCtrl` | ETH value to send with the payable `makePayment` call |

### Create Form (toggle)
| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| Borrower (0x…) | `_borrowerCtrl` | `POST /loan/` — auto-filled from Privy wallet |
| Principal (wei) | `_principalCtrl` | (same) |
| Interest Rate (bps) | `_rateBpsCtrl` | (same) |
| Duration (seconds) | `_durationCtrl` | (same) |

> **Removed** (not user-facing):
> - **DAO ID** field — Contract's second parameter is unused (reserved for future DAO-linked features). The API hardcodes `dao_id: '0'`.
> - **Mark Defaulted** button — No access control on the contract = griefable. Handled exclusively by `loan_default_monitor` CRE cron workflow.

**Note:** Borrower field is auto-filled from Privy wallet address.

## Screen → API → Contract → Workflow Flow

```plantuml
@startuml Loan_Screen_Flow
skinparam actorStyle awesome

actor User
participant "Loan Screen" as Scr
participant "LoanService" as Svc
participant "ApiClient" as API
participant "Backend\n/loan" as BE
participant "LoanManager" as LM
participant "CRE Workflows" as CRE

== Create Loan ==
User -> Scr : fill form, tap "Create Loan"
Scr -> Svc : createLoan(borrower, principal, rateBps, duration)
Svc -> API : POST /loan/ {borrower, dao_id: '0', principal, rate, duration}
API -> BE : HTTP + JWT
BE -> LM : createLoan(borrower, 0, principal, rateBps, duration)
LM --> BE : tx
BE -> BE : WaitForLoanID(tx)
BE -> BE : SaveLoan(PENDING)
BE --> API : {tx, loan_id}
note right of LM: Emits LoanCreated\n→ CRE loan_created\n(checks risk tier,\nflags if over-limit)

== Approve Loan ==
User -> Scr : tap "Approve"
Scr -> Svc : approveLoan(id)
Svc -> API : POST /loan/{id}/approve
API -> BE : HTTP + JWT
BE -> LM : approveLoan(id)
LM --> BE : tx
BE -> BE : UpdateLoanStatus(APPROVED)
note right of LM: Emits LoanApproved\n→ CRE loan_approved\n(reads full snapshot)

== Make Payment ==
User -> Scr : enter amount, tap "Make Payment"
Scr -> Svc : makePayment(id, amount)
Svc -> API : POST /loan/{id}/payment {amount}
API -> BE : HTTP + JWT
BE -> BE : parse amount, set msg.value
BE -> LM : makePayment(id) {value: amount}
LM -> LM : if paid >= owed → REPAID
LM --> BE : tx
BE -> BE : UpdateLoanStatus(PAYMENT_MADE)
note right of LM: Emits PaymentMade\n→ CRE loan_payment\n(risk +100 / +700 payoff)

note over CRE: Loan default is automated by\nloan_default_monitor (cron 1hr)\n→ markDefaulted + risk -2000

@enduml
```
