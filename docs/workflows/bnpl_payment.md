# bnpl_payment Workflow

**Source:** `workflows/bnpl_payment/main.go`  
**Trigger:** EVM Log — `BNPLPaymentMade(uint256 indexed arrangementId, address indexed payer, uint8 installmentNumber, uint256 amount)`  
**Contracts:** BNPLManager, DIDRegistry

## Purpose

When a BNPL installment payment is made on-chain, this workflow:
1. Decodes payment details from the event
2. Reads the full arrangement to check if this is the final payment
3. Reads the payer's current DID risk score
4. Adjusts risk score upward (+50 per payment, +500 if final)
5. Writes the updated risk profile on-chain
6. Notifies the backend

## Risk Adjustments

| Condition | Delta | Reason |
|-----------|-------|--------|
| Regular payment | +50 | `bnpl_payment` |
| Final payment (completes arrangement) | +500 | `bnpl_final_payment` |

## Flow

```plantuml
@startuml bnpl_payment_flow

participant "BNPLManager" as BM
participant "CRE DON\nbnpl_payment" as W
participant "DIDRegistry" as DID
participant "Backend API" as BE

BM -> W : BNPLPaymentMade event log
activate W

W -> W : decode topics\n(arrangementId, payer)
W -> W : decode data\n(installmentNumber, amount)

W -> BM : callContract\ngetArrangement(arrangementId)
BM --> W : Arrangement struct
W -> W : isFinal = (installment >= numInstallments - 1)

W -> DID : callContract\ngetRiskProfileScore(payer)
DID --> W : currentScore

W -> W : compute newScore\n= adjustScore(current, delta)
W -> W : compute profileHash

W -> DID : writeReport\nupdateRiskProfile(payer, newScore, hash)
note right: On-chain risk update\nEmits RiskProfileUpdated

W -> BE : POST /api/bnpl/payment\n{arrangementId, payer,\ninstallmentNumber, amount,\nisFinal, newScore, tier}

deactivate W

@enduml
```
