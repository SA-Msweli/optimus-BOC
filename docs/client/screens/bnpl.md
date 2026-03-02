# BNPL Screen

**Source:** `client/prime/lib/screens/bnpl_screen.dart`  
**Service:** `BNPLService`  
**Tab:** BNPL (index 1)

## UI Elements

### Lookup Section
| Element | Controller | API Endpoint |
|---------|-----------|-------------|
| Arrangement ID | `_arrIdCtrl` | `GET /bnpl/arrangements/{id}` |
| Fetch | Button | Triggers fetch |

### Arrangement Detail (when loaded)
Displays: arrangementId, daoId, payer, recipient, totalAmount, numInstallments, startTimestamp, intervalSeconds, lateFeeBps, status

Status labels: `PENDING` (0), `ACTIVE` (1), `COMPLETED` (2)

### Payment Section
| Field | Controller | Purpose |
|-------|-----------|---------|
| Installment # | `_installmentCtrl` | Which installment to pay |
| Amount (wei) | `_payAmountCtrl` | ETH value to send with the payable call |
| Pay | Button | `POST /bnpl/arrangements/{id}/payment` |

### Reschedule Section
| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| New Start (unix) | `_newStartCtrl` | `POST /bnpl/arrangements/{id}/reschedule` |
| New Interval (sec) | `_newIntervalCtrl` | (same endpoint) |

### Create Form (toggle)
| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| DAO ID | `_daoIdCtrl` | `POST /bnpl/arrangements` |
| Recipient Address | `_recipientCtrl` | (same) — auto-filled from Privy wallet |
| Total Amount (wei) | `_totalCtrl` | (same) |
| Start Date (unix) | `_startCtrl` | (same) |
| Payment Interval (sec) | `_intervalCtrl` | (same) |

> **Removed** (not user-facing):
> - **Activate** button — Contract auto-activates on first payment. Manual activation is redundant.
> - **Late Fee** button — No access control on contract = griefable. Applied exclusively by `bnpl_late_fee` CRE cron workflow.

## Screen → API → Contract → Workflow Flow

```plantuml
@startuml BNPL_Screen_Flow
skinparam actorStyle awesome

actor User
participant "BNPL Screen" as Scr
participant "BNPLService" as Svc
participant "ApiClient" as API
participant "Backend\n/bnpl" as BE
participant "BNPLManager" as BM
participant "DAOManager" as DM
participant "CRE Workflows" as CRE

== Create Arrangement ==
User -> Scr : fill form, tap "Create"
Scr -> Svc : createArrangement(daoId, recipient, total, start, interval)
Svc -> API : POST /bnpl/arrangements
API -> BE : HTTP + JWT
BE -> BM : createBNPL(daoId, recipient, total, start, interval, "")
BM -> DM : getBnplTerms(daoId)
DM --> BM : terms (validates)
BM --> BE : tx
BE -> BE : WaitForArrangementID(tx)
BE -> BE : SaveArrangement(PENDING)
BE --> API : {tx, arrangement_id}
API --> Svc : update current
Svc --> Scr : notifyListeners()
note right of BM: Emits BNPLCreated\n→ CRE bnpl_created

== Make Payment ==
User -> Scr : enter installment # + amount, tap "Pay"
Scr -> Svc : makePayment(id, installment, amount)
Svc -> API : POST /bnpl/arrangements/{id}/payment\n{installment, amount}
API -> BE : HTTP + JWT
BE -> BE : parse amount, set msg.value
BE -> BM : makePayment(id, installment) {value: amount}
BM -> BM : auto-activate if first payment
BM -> BM : auto-complete if all paid
BE --> API : {tx}
note right of BM: Emits BNPLPaymentMade\n→ CRE bnpl_payment\n(risk +50 / +500 final)

== Reschedule ==
User -> Scr : enter new start/interval, tap "Reschedule"
Scr -> Svc : reschedule(id, newStart, newInterval)
Svc -> API : POST /bnpl/arrangements/{id}/reschedule
API -> BE : HTTP + JWT
BE -> BM : reschedule(id, newStart, newInterval)
BM -> DM : getBnplTerms(daoId)
note right: Validates against DAO policy

note over CRE: Late fees applied by\nbnpl_late_fee (cron 30min)\n→ risk -300 per violation

@enduml
```
