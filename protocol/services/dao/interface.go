package dao

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IDAOService defines the contract for DAO operations.
// Matches the on-chain IDAOManager interface from OPTIMUS_ETH_SPEC.

type IDAOService interface {
	// DAO lifecycle
	CreateDAO(opts *bind.TransactOpts, creator common.Address, goal uint8, votingPeriodDays uint64) (*types.Transaction, error)
	JoinDAO(opts *bind.TransactOpts, daoId *big.Int, member common.Address, investment *big.Int) (*types.Transaction, error)

	// Governance
	Propose(opts *bind.TransactOpts, daoId *big.Int, data []byte) (*types.Transaction, error)
	Vote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error)
	FinalizeProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error)
	ExecuteProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error)

	// BNPL policy
	SetBnplTerms(opts *bind.TransactOpts, daoId *big.Int, numInstallments *big.Int, minDays *big.Int, maxDays *big.Int, lateFeeBps *big.Int, graceDays *big.Int, rescheduleAllowed bool, minDownBps *big.Int) (*types.Transaction, error)
	GetBnplTerms(ctx context.Context, daoId *big.Int) (BnplTerms, error)

	// Treasury
	GetTreasuryBalance(ctx context.Context, daoId *big.Int) (*big.Int, error)
	CreditTreasury(opts *bind.TransactOpts, daoId *big.Int, amount *big.Int) (*types.Transaction, error)

	// Query helpers
	IsMember(ctx context.Context, daoId *big.Int, member common.Address) (bool, error)
	WaitForDaoID(ctx context.Context, tx *types.Transaction) (uint64, error)
}

// BnplTerms is a friendlier representation of the on-chain BNPL policy.
type BnplTerms struct {
	NumInstallments        *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps             *big.Int
	GracePeriodDays        *big.Int
	RescheduleAllowed      bool
	MinDownPaymentBps      *big.Int
}
