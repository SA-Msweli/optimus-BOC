// Package shared – Config is the canonical CRE workflow configuration
// struct shared by all Optimus workflows.  The CRE runtime loads the
// matching config.*.json at deploy time and injects it via the Runner
// generic parameter.
package shared

// Config is loaded from the per-workflow config.staging.json or
// config.production.json.  All workflows use the same struct so
// they can share a single configuration schema.
type Config struct {
	// Vault DON secret owner.  Empty string for local simulation;
	// set to the workflow owner address in production.
	Owner string `json:"owner"`

	// Chain name for evm.ChainSelectorFromName (e.g. "ethereum-testnet-sepolia").
	ChainSelector string `json:"chainSelector"`

	// Cron schedule (6-field with seconds prefix).  Only used by cron-triggered workflows.
	Schedule string `json:"schedule,omitempty"`

	// Contract addresses on the target chain
	BNPLManagerAddr string `json:"bnplManagerAddress"`
	LoanManagerAddr string `json:"loanManagerAddress"`
	DAOManagerAddr  string `json:"daoManagerAddress"`
	DIDRegistryAddr string `json:"didRegistryAddress"`
	TokenVaultAddr  string `json:"tokenVaultAddress"`

	// Backend API for offchain state and notifications
	BackendURL string `json:"backendUrl"`

	// ── Workflow-specific parameters ──

	// GracePeriodSeconds is added to each installment due date before
	// considering it overdue.  Used by bnpl_late_fee workflow.
	GracePeriodSeconds int64 `json:"gracePeriodSeconds,omitempty"`

	// TreasuryAlertThresholdWei triggers a low-balance alert when the
	// TokenVault balance falls below this value.
	TreasuryAlertThresholdWei string `json:"treasuryAlertThresholdWei,omitempty"`

	// ActiveArrangementIDs is the list of BNPL arrangement IDs the cron
	// workflows should check.  Populated from the backend at deploy time
	// or refreshed via HTTP within the workflow.
	ActiveArrangementIDs []int64 `json:"activeArrangementIds,omitempty"`

	// ActiveLoanIDs is the list of Loan IDs the cron workflows should check.
	ActiveLoanIDs []int64 `json:"activeLoanIds,omitempty"`

	// ActiveProposalIDs is the list of DAO proposal IDs to check for expiry.
	ActiveProposalIDs []int64 `json:"activeProposalIds,omitempty"`

	// ETHTokenAddress is the zero address for native ETH balance queries.
	ETHTokenAddress string `json:"ethTokenAddress,omitempty"`
}
