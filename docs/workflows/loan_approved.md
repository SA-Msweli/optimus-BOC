# loan_approved Workflow

**Source:** `workflows/loan_approved/main.go`  
**Trigger:** EVM Log — `LoanApproved(uint256 indexed loanId, address indexed approver)`  
**Contract:** LoanManager

## Purpose

When a loan is approved on-chain:
1. Reads full loan state (principal, rates, terms)
2. Reads current amount owed and accrued interest
3. Assembles a comprehensive loan snapshot
4. Notifies the backend with the full snapshot

## Flow

```plantuml
@startuml loan_approved_flow

participant "LoanManager" as LM
participant "CRE DON\nloan_approved" as W
participant "Backend API" as BE

LM -> W : LoanApproved event log
activate W

W -> W : decode topics\n(loanId, approver)

W -> LM : callContract\ngetLoan(loanId)
LM --> W : Loan struct

W -> LM : callContract\ngetAmountOwed(loanId)
LM --> W : amountOwed

W -> LM : callContract\ngetAccruedInterest(loanId)
LM --> W : accruedInterest

W -> BE : POST /api/loans/approved\n{loanId, borrower, approver,\nprincipal, interestRateBps,\nstartTime, endTime,\namountOwed, accruedInterest,\namountPaid, status:ACTIVE}

deactivate W

@enduml
```
