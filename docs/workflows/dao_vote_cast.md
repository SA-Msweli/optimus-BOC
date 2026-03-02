# dao_vote_cast Workflow

**Source:** `workflows/dao_vote_cast/main.go`  
**Trigger:** EVM Log — `VoteCast(uint256 indexed proposalId, address indexed voter, bool support, uint256 weight)`  
**Contract:** DAOManager

## Purpose

When a vote is cast on a DAO proposal:
1. Decodes vote details (direction, weight)
2. Notifies the backend of the individual vote
3. Sends a quorum-check request to the backend (backend tracks cumulative votes)

## Flow

```plantuml
@startuml dao_vote_cast_flow

participant "DAOManager" as DM
participant "CRE DON\ndao_vote_cast" as W
participant "Backend API" as BE

DM -> W : VoteCast event log
activate W

W -> W : decode topics\n(proposalId, voter)
W -> W : decode data\n(support, weight)
W -> W : direction = support ? "FOR" : "AGAINST"

W -> BE : POST /api/dao/votes\n{proposalId, voter,\nsupport, weight, direction}

W -> BE : POST /api/dao/proposals/{id}/quorum-check\n{proposalId, latestVoter,\nlatestWeight, latestSupport}

alt response 200
  W -> W : log "quorum may have been reached"
end

deactivate W

@enduml
```
