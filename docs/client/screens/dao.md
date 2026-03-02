# DAO Screen

**Source:** `client/prime/lib/screens/dao_screen.dart`  
**Service:** `DAOService`  
**Tab:** DAO (index 3)  
**Layout:** 4-tab TabBarView

## Tab Structure

| Tab | Name | Functionality |
|-----|------|---------------|
| 0 | DAO | Create DAO + Register Member |
| 1 | Proposals | Propose, Vote, Finalize, Execute |
| 2 | Terms | View + Set BNPL Terms |
| 3 | Treasury | View Balance + Credit |

## Tab 0: DAO (Create & Join)

| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| Goal | `_selectedGoal` dropdown | `POST /dao/` → `goal: 0/1/2` |
| Voting Period (days) | `_votingPeriodCtrl` | `POST /dao/` → `voting_period_days` |
| DAO ID (join) | `_joinDaoIdCtrl` | `POST /dao/{id}/join` |
| Voting Weight | `_investmentCtrl` | `POST /dao/{id}/join` → `investment` |

Goal dropdown options: Savings (0), Lending (1), Investment (2)

> **Removed:** DAO Name field — The contract has no name field; it uses the goal enum.

## Tab 1: Proposals

The proposal tab supports two modes: **structured** (treasury withdrawal) and **raw hex** (generic proposals).

### Structured Mode (default)
| Field | Controller | Purpose |
|-------|-----------|---------|
| DAO ID | `_proposeDaoIdCtrl` | Target DAO for the proposal |
| Token Address | `_propTokenCtrl` | ERC-20 token to withdraw |
| Amount (wei) | `_propAmountCtrl` | Amount to withdraw |
| Recipient (0x…) | `_propRecipientCtrl` | Who receives the tokens |

The service builds `abi.encode(address token, uint256 amount, address recipient)` as the proposal data (3×32-byte hex-encoded values).

### Raw Mode (toggle)
| Field | Controller | Purpose |
|-------|-----------|---------|
| Proposal Data (hex) | `_propRawDataCtrl` | Arbitrary proposal payload |
| Raw Data toggle | `_useRawProposal` SwitchListTile | Switches between structured and raw mode |

### Vote
| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| Proposal ID | `_voteProposalIdCtrl` | `POST /dao/proposal/{id}/vote` |
| DAO ID | `_voteDaoIdCtrl` | (same) → `dao_id` — **required** with validation |
| Support toggle | `_voteSupport` | (same) → `support: bool` |

### Finalize / Execute
| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| Proposal ID | `_actionProposalIdCtrl` | `POST /dao/proposal/{id}/finalize` or `execute` |

> **Changed:** Vote DAO ID is now **required** (with validation). Previously was labeled "(optional)".
> **Changed:** "Initial Stake (wei)" → **"Voting Weight (accounting units)"** with clarifying hint.

## Tab 2: BNPL Terms

| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| DAO ID | `_termsDaoIdCtrl` | `GET /dao/{id}/bnpl-terms` |
| Num Installments | `_numInstCtrl` | `POST /dao/{id}/bnpl-terms` |
| Min Days | `_minDaysCtrl` | (same) |
| Max Days | `_maxDaysCtrl` | (same) |
| Late Fee (bps) | `_lateFeeBpsCtrl` | (same) |
| Grace Days | `_graceDaysCtrl` | (same) |
| Min Down Payment (bps) | `_minDownBpsCtrl` | (same) |
| Reschedule Allowed | `_rescheduleAllowed` switch | (same) |

## Tab 3: Treasury

| Field | Controller | API Endpoint |
|-------|-----------|-------------|
| DAO ID | `_treasuryDaoIdCtrl` | `GET /dao/{id}/treasury` |
| Amount (wei) | `_creditAmountCtrl` | `POST /dao/{id}/treasury/credit` |

## Full DAO Flow Diagram

```plantuml
@startuml DAO_Screen_Flow
skinparam actorStyle awesome

actor User
participant "DAO Screen" as Scr
participant "DAOService" as Svc
participant "ApiClient" as API
participant "Backend\n/dao" as BE
participant "DAOManager" as DM
participant "TokenVault" as TV
participant "CRE Workflows" as CRE

== Tab 0: Create DAO ==
User -> Scr : select goal, period, tap "Create DAO"
Scr -> Svc : createDAO(goal, votingPeriodDays)
Svc -> API : POST /dao/ {goal, voting_period_days}
API -> BE : HTTP + JWT
BE -> DM : createDAO(from, goal, days)
DM --> BE : tx
BE -> BE : WaitForDaoID(tx)
BE -> BE : SaveDAO(daoId, creator, goal, period)
BE --> API : {tx, dao_id}

== Tab 0: Register Member ==
User -> Scr : enter DAO ID + weight, tap "Register"
Scr -> Svc : joinDAO(daoId, investment)
Svc -> API : POST /dao/{id}/join {investment}
API -> BE : HTTP + JWT
BE -> DM : joinDAO(daoId, from, investment)
DM --> BE : tx
note right of DM: Emits MemberJoined\nmemberCount++, treasury += investment

== Tab 1: Create Treasury Withdrawal Proposal ==
User -> Scr : enter token, amount, recipient
Scr -> Svc : proposeTreasuryWithdrawal(daoId, token, amount, recipient)
Svc -> Svc : ABI-encode (token + amount + recipient) as 96-byte hex
Svc -> API : POST /dao/{daoId}/propose {data: "0x..."}
API -> BE : HTTP + JWT
BE -> DM : propose(daoId, data)
DM --> BE : tx
note right of DM: Emits ProposalOpened\n→ CRE dao_proposal_opened

== Tab 1: Vote ==
User -> Scr : enter proposal ID + DAO ID, toggle FOR/AGAINST, tap "Vote"
Scr -> Svc : vote(proposalId, support, daoId)
Svc -> API : POST /dao/proposal/{id}/vote {dao_id, support}
API -> BE : HTTP + JWT
BE -> DM : vote(proposalId, support)
DM -> DM : weight = memberInvestment (min 1)
note right of DM: Emits VoteCast\n→ CRE dao_vote_cast

== Tab 1: Finalize ==
User -> Scr : tap "Finalize"
Scr -> Svc : finalizeProposal(id)
Svc -> API : POST /dao/proposal/{id}/finalize
note right: Must be past expiry\napproved = yesWeight > noWeight

== Tab 1: Execute ==
User -> Scr : tap "Execute"
Scr -> Svc : executeProposal(id)
Svc -> API : POST /dao/proposal/{id}/execute
alt proposal data = (token, amount, recipient)
  DM -> TV : withdraw(token, amount)
  DM -> DM : transfer to recipient
end

== Tab 2: Set BNPL Terms ==
User -> Scr : fill 7 fields, tap "Set Terms"
Scr -> Svc : setBnplTerms(daoId, ...)
Svc -> API : POST /dao/{id}/bnpl-terms {7 params}
note right of DM: Emits BnplTermsUpdated\nAffects future BNPL arrangements

== Tab 3: Credit Treasury ==
User -> Scr : enter amount, tap "Credit"
Scr -> Svc : creditTreasury(daoId, amount)
Svc -> API : POST /dao/{id}/treasury/credit {amount}

note over CRE: dao_proposal_finalizer (cron 1hr)\nauto-finalizes expired proposals\n\ntreasury_monitor (cron 6hr)\nchecks DAO treasury health

@enduml
```
