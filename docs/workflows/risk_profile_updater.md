# risk_profile_updater Workflow

**Source:** `workflows/risk_profile_updater/main.go`  
**Trigger:** EVM Log — `RiskProfileUpdated(address indexed owner, uint256 newScore, bytes32 indexed profileHash)`  
**Contract:** DIDRegistry

## Purpose

When a risk profile is updated on-chain (by any of the other workflows):
1. Decodes the new score from the event
2. Computes credit tier, max BNPL amount, max loan principal
3. Reads the previous tier from the backend to detect tier transitions
4. Notifies the backend of the update and any tier change

## Credit Tiers

| Tier | Score Range |
|------|-------------|
| EXCELLENT | ≥ 800 |
| GOOD | ≥ 600 |
| FAIR | ≥ 400 |
| POOR | < 400 |

## Flow

```plantuml
@startuml risk_profile_updater_flow

participant "DIDRegistry" as DID
participant "CRE DON\nrisk_profile_updater" as W
participant "Backend API" as BE

DID -> W : RiskProfileUpdated event log
activate W

W -> W : decode topics\n(owner, profileHash)
W -> W : decode data\n(newScore)

W -> W : tier = CreditTier(newScore)
W -> W : maxBNPL = MaxBNPLAmount(newScore)
W -> W : maxLoan = MaxLoanPrincipal(newScore)

W -> BE : POST /api/did/risk-updated\n{owner, newScore, tier,\nmaxBnplWei, maxLoanWei}

W -> BE : GET /api/did/{owner}/previous-tier
BE --> W : previousTier

alt previousTier != tier
  W -> BE : POST /api/did/tier-changed\n{owner, previousTier, newTier,\nnewScore, maxBnplWei, maxLoanWei}
  W -> W : log "tier changed from X to Y"
end

deactivate W

@enduml
```
