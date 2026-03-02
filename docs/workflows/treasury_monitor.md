# treasury_monitor Workflow

**Source:** `workflows/treasury_monitor/main.go`  
**Trigger:** Cron — every 6 hours  
**Contracts:** TokenVault, DAOManager

## Purpose

Monitors the TokenVault ETH balance and each DAO treasury balance. Alerts the backend when the vault balance drops below a configured threshold.

## Flow

```plantuml
@startuml treasury_monitor_flow

participant "Cron\n(6 hours)" as Cron
participant "CRE DON\ntreasury_monitor" as W
participant "TokenVault" as TV
participant "DAOManager" as DM
participant "Backend API" as BE

Cron -> W : trigger
activate W

== TokenVault Balance ==
W -> TV : callContract\ngetBalance(ethTokenAddress)
TV --> W : vaultBalance

alt vaultBalance < threshold
  W -> BE : POST /api/treasury/alert\n{vaultBalance, threshold}
  W -> W : WARN "vault balance below threshold"
end

== DAO Treasuries ==
W -> BE : GET /api/dao/active-ids
BE --> W : [1, 2, ...]

loop for each DAO
  W -> DM : callContract\ngetTreasuryBalance(daoId)
  DM --> W : balance

  W -> BE : POST /api/dao/{id}/treasury-status\n{daoId, treasuryBalance}
end

W -> W : log "vault_balance=X, daos_checked=N"

deactivate W

@enduml
```
