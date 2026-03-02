# loan_default_monitor Workflow

**Source:** `workflows/loan_default_monitor/main.go`  
**Trigger:** Cron — every hour  
**Contracts:** LoanManager, DIDRegistry

## Purpose

Scans active loans past their end time. For each defaultable loan:
1. Verifies outstanding balance > 0 via `getAmountOwed`
2. Calls `LoanManager.markDefaulted` on-chain
3. Slashes the borrower's DID risk score (-2000)
4. Notifies the backend

## Risk Adjustments

| Condition | Delta | Reason |
|-----------|-------|--------|
| Loan defaulted | -2000 | `loan_defaulted` |

## Flow

```plantuml
@startuml loan_default_monitor_flow

participant "Cron\n(1 hour)" as Cron
participant "CRE DON\nloan_default_monitor" as W
participant "Backend API" as BE
participant "LoanManager" as LM
participant "DIDRegistry" as DID

Cron -> W : trigger
activate W

alt no static IDs in config
  W -> BE : GET /api/loans/active-ids
  BE --> W : [1, 2, 3, ...]
end

loop for each loan ID
  W -> LM : callContract\ngetLoan(loanId)
  LM --> W : Loan struct
  
  alt status != ACTIVE or now <= endTime
    W -> W : skip
  end

  W -> LM : callContract\ngetAmountOwed(loanId)
  LM --> W : amountOwed

  alt amountOwed == 0
    W -> W : skip (already repaid)
  end

  W -> LM : writeReport\nmarkDefaulted(loanId)
  note right: Emits LoanDefaulted

  W -> DID : callContract\ngetRiskProfileScore(borrower)
  DID --> W : currentScore
  W -> W : newScore = adjustScore(current, -2000)
  W -> DID : writeReport\nupdateRiskProfile(borrower, newScore, hash)

  W -> BE : POST /api/loans/{id}/defaulted\n{loanId, borrower, amountOwed}
end

W -> W : log "defaulted=N"

deactivate W

@enduml
```
