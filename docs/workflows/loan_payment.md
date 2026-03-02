# loan_payment Workflow

**Source:** `workflows/loan_payment/main.go`  
**Trigger:** EVM Log — `PaymentMade(uint256 indexed loanId, address indexed payer, uint256 amount, uint256 amountPaid, uint256 remaining)`  
**Contracts:** LoanManager, DIDRegistry

## Purpose

When a loan payment is made on-chain:
1. Decodes payment event data
2. Reads borrower's current risk score
3. Checks remaining interest to determine if loan is fully repaid
4. Adjusts risk score upward
5. Writes updated risk profile on-chain
6. Notifies the backend

## Risk Adjustments

| Condition | Delta | Reason |
|-----------|-------|--------|
| Regular payment | +100 | `loan_payment` |
| Full repayment (remaining=0, interest=0) | +700 | `loan_payoff` |

## Flow

```plantuml
@startuml loan_payment_flow

participant "LoanManager" as LM
participant "CRE DON\nloan_payment" as W
participant "DIDRegistry" as DID
participant "Backend API" as BE

LM -> W : PaymentMade event log
activate W

W -> W : decode topics\n(loanId, payer)
W -> W : decode data\n(amount, totalPaid, remaining)

W -> DID : callContract\ngetRiskProfileScore(payer)
DID --> W : currentScore

W -> LM : callContract\ngetAccruedInterest(loanId)
LM --> W : interest

W -> W : isPayoff = (remaining==0 && interest==0)
W -> W : delta = isPayoff ? +700 : +100
W -> W : newScore = adjustScore(current, delta)

W -> DID : writeReport\nupdateRiskProfile(payer, newScore, hash)
note right: On-chain risk update\nEmits RiskProfileUpdated

W -> BE : POST /api/loans/payment\n{loanId, payer, amount,\ntotalPaid, remaining,\nisPayoff, newScore, tier}

deactivate W

@enduml
```
