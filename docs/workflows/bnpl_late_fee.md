# bnpl_late_fee Workflow

**Source:** `workflows/bnpl_late_fee/main.go`  
**Trigger:** Cron — every 30 minutes  
**Contracts:** BNPLManager, DAOManager, DIDRegistry

## Purpose

Scans active BNPL arrangements for overdue installments past the grace period. For each overdue installment:
1. Calls `BNPLManager.applyLateFee` on-chain (credits DAO treasury)
2. Reads the payer's risk score and applies -300 penalty
3. Writes updated risk profile on-chain

## Risk Adjustments

| Condition | Delta | Reason |
|-----------|-------|--------|
| Late fee applied per installment | -300 | `bnpl_late_fee` |

## Flow

```plantuml
@startuml bnpl_late_fee_flow

participant "Cron\n(30 min)" as Cron
participant "CRE DON\nbnpl_late_fee" as W
participant "Backend API" as BE
participant "BNPLManager" as BM
participant "DAOManager" as DM
participant "DIDRegistry" as DID

Cron -> W : trigger
activate W

alt no static IDs in config
  W -> BE : GET /api/bnpl/active-ids
  BE --> W : [1, 2, 3, ...]
end

loop for each arrangement ID
  W -> BM : callContract\ngetArrangement(id)
  BM --> W : Arrangement struct
  
  alt status != ACTIVE
    W -> W : skip
  end

  W -> DM : callContract\ngetBnplTerms(daoId)
  DM --> W : BnplTerms (grace period)

  loop for each installment
    W -> W : check: now > dueTime + gracePeriod?
    
    alt overdue
      W -> BM : writeReport\napplyLateFee(id, installmentNo)
      note right: Credits DAO treasury\nwith late fee amount
      
      W -> DID : callContract\ngetRiskProfileScore(payer)
      DID --> W : currentScore
      W -> W : newScore = adjustScore(current, -300)
      W -> DID : writeReport\nupdateRiskProfile(payer, newScore, hash)
    end
  end
end

W -> W : log "applied=N"

deactivate W

@enduml
```
