# dao_proposal_finalizer Workflow

**Source:** `workflows/dao_proposal_finalizer/main.go`  
**Trigger:** Cron — every hour  
**Contract:** DAOManager

## Purpose

Scans open DAO proposals and finalizes any that have passed their expiry timestamp by calling `DAOManager.finalizeProposal` on-chain.

## Flow

```plantuml
@startuml dao_proposal_finalizer_flow

participant "Cron\n(1 hour)" as Cron
participant "CRE DON\ndao_proposal_finalizer" as W
participant "Backend API" as BE
participant "DAOManager" as DM

Cron -> W : trigger
activate W

alt no static IDs in config
  W -> BE : GET /api/dao/proposals/open-ids
  BE --> W : [1, 2, 3, ...]
end

loop for each proposal ID
  W -> BE : GET /api/dao/proposals/{id}
  BE --> W : {expiry: ...}

  alt now > expiry
    W -> DM : writeReport\nfinalizeProposal(proposalId)
    note right: Sets approved = (yesWeight > noWeight)\nEmits ProposalFinalized

    W -> BE : POST /api/dao/proposals/{id}/finalized\n{proposalId, status:FINALIZED}
  end
end

W -> W : log "finalized=N"

deactivate W

@enduml
```
