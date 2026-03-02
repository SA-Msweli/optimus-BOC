# state_reconciler Workflow

**Source:** `workflows/state_reconciler/main.go`  
**Trigger:** Cron — every 2 hours  
**Contracts:** BNPLManager, LoanManager, DAOManager

## Purpose

Cross-checks on-chain state with the backend database for all active entities. Reports mismatches for reconciliation.

## Reconciliation Sections

1. **BNPL Arrangements** — reads on-chain status/amounts, sends to `/api/reconcile/bnpl`
2. **Loans** — reads on-chain loan state + amount owed, sends to `/api/reconcile/loan`
3. **DAO Treasuries** — reads on-chain treasury balances, sends to `/api/reconcile/dao-treasury`

## Flow

```plantuml
@startuml state_reconciler_flow

participant "Cron\n(2 hours)" as Cron
participant "CRE DON\nstate_reconciler" as W
participant "Backend API" as BE
participant "BNPLManager" as BM
participant "LoanManager" as LM
participant "DAOManager" as DM

Cron -> W : trigger
activate W

== Section 1: BNPL ==
W -> BE : GET /api/bnpl/active-ids
BE --> W : [1, 2, ...]

loop for each arrangement
  W -> BM : callContract getArrangement(id)
  BM --> W : on-chain state
  W -> BE : POST /api/reconcile/bnpl\n{arrangementId, status, totalAmount, numInstallments}
  BE --> W : {mismatch: true/false}
end

== Section 2: Loans ==
W -> BE : GET /api/loans/active-ids
BE --> W : [1, 2, ...]

loop for each loan
  W -> LM : callContract getLoan(id)
  LM --> W : on-chain state
  W -> LM : callContract getAmountOwed(id)
  LM --> W : amountOwed
  W -> BE : POST /api/reconcile/loan\n{loanId, status, principal, amountPaid, amountOwed}
  BE --> W : {mismatch: true/false}
end

== Section 3: DAO Treasuries ==
W -> BE : GET /api/dao/active-ids
BE --> W : [1, 2, ...]

loop for each DAO
  W -> DM : callContract getTreasuryBalance(id)
  DM --> W : balance
  W -> BE : POST /api/reconcile/dao-treasury\n{daoId, treasuryBalance}
  BE --> W : {mismatch: true/false}
end

W -> W : log "mismatches=N"

deactivate W

@enduml
```
