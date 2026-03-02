# bnpl_created Workflow

**Source:** `workflows/bnpl_created/main.go`  
**Trigger:** EVM Log — `BNPLCreated(uint256 indexed arrangementId, uint256 indexed daoId, address indexed payer, address recipient, uint256 totalAmount, uint256 numInstallments)`  
**Contract:** BNPLManager

## Purpose

When a new BNPL arrangement is created on-chain, this workflow:
1. Reads the full arrangement state from BNPLManager
2. Reads the DAO's BNPL terms for policy validation
3. Validates arrangement parameters against DAO policy
4. Computes the installment schedule
5. Notifies the backend via confidential HTTP

## Flow

```plantuml
@startuml bnpl_created_flow

participant "BNPLManager\nContract" as BM
participant "CRE DON\nbnpl_created" as W
participant "DAOManager\nContract" as DM
participant "Backend API" as BE

BM -> W : BNPLCreated event log
activate W

W -> W : decode topics\n(arrangementId, daoId, payer)
W -> W : decode data\n(recipient, totalAmount, numInstallments)

W -> BM : callContract\ngetArrangement(arrangementId)
BM --> W : Arrangement struct

W -> DM : callContract\ngetBnplTerms(daoId)
DM --> W : BnplTerms struct

W -> W : validateArrangement\n(check installment count,\ninterval min/max)

W -> W : compute installment\nschedule (amounts, due dates)

W -> BE : POST /api/bnpl/created\n{arrangementId, daoId, payer,\nrecipient, totalAmount,\nnumInstallments, violations}

deactivate W

@enduml
```

## Policy Validation

The workflow checks the arrangement against the DAO's BNPL terms:

| Check | Rule |
|-------|------|
| Installment count | Must match `terms.NumInstallments` (if set) |
| Interval min | `intervalDays >= terms.AllowedIntervalMinDays` |
| Interval max | `intervalDays <= terms.AllowedIntervalMaxDays` |

Violations are logged as warnings and reported to the backend.
