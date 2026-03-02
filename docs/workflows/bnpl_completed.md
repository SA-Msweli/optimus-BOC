# bnpl_completed Workflow

**Source:** `workflows/bnpl_completed/main.go`  
**Trigger:** EVM Log — `BNPLCompleted(uint256 indexed arrangementId, uint256 completedAt)`  
**Contracts:** BNPLManager, DAOManager

## Purpose

When a BNPL arrangement is marked completed on-chain, this workflow:
1. Reads the full arrangement state for final reporting
2. Notifies the backend of completion
3. Reads the DAO treasury balance
4. Triggers a treasury rebalance check if balance falls below threshold

## Flow

```plantuml
@startuml bnpl_completed_flow

participant "BNPLManager" as BM
participant "CRE DON\nbnpl_completed" as W
participant "DAOManager" as DM
participant "Backend API" as BE

BM -> W : BNPLCompleted event log
activate W

W -> W : decode\n(arrangementId, completedAt)

W -> BM : callContract\ngetArrangement(arrangementId)
BM --> W : Arrangement struct

W -> BE : POST /api/bnpl/completed\n{arrangementId, daoId,\npayer, totalAmount,\ncompletedAt, status:COMPLETED}

W -> DM : callContract\ngetTreasuryBalance(daoId)
DM --> W : balance

alt balance < threshold
  W -> BE : POST /api/dao/{daoId}/treasury-rebalance\n{daoId, currentBalance,\nthreshold, trigger:bnpl_completed}
  note right: Low treasury alert
end

deactivate W

@enduml
```
