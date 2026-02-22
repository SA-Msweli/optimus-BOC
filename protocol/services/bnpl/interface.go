package bnpl

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IBNPLService defines the contract for BNPL operations.
// Matches the on-chain IBNPLManager interface from OPTIMUS_ETH_SPEC.

type IBNPLService interface {
	// Write operations
	CreateBNPL(opts *bind.TransactOpts, daoId *big.Int, recipient common.Address, totalAmount *big.Int, startTimestamp *big.Int, intervalSeconds *big.Int, metadata []byte) (*types.Transaction, error)
	MakePayment(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error)
	ActivateBNPL(opts *bind.TransactOpts, arrangementId *big.Int) (*types.Transaction, error)
	ApplyLateFee(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error)
	Reschedule(opts *bind.TransactOpts, arrangementId *big.Int, newStartTimestamp *big.Int, newIntervalSeconds *big.Int) (*types.Transaction, error)

	// Read operations
	GetArrangement(ctx context.Context, id *big.Int) (BNPLArrangement, error)

	// Helpers
	WaitForArrangementID(ctx context.Context, tx *types.Transaction) (uint64, error)
}

// BNPLArrangement mirrors the data returned by the contract's getArrangement
// view method so consumers don't have to deal with the generated anonymous struct.
type BNPLArrangement struct {
	ID                 *big.Int
	DaoId              *big.Int
	Payer              common.Address
	Recipient          common.Address
	TotalAmount        *big.Int
	NumInstallments    *big.Int
	InstallmentAmounts []*big.Int
	StartTimestamp     *big.Int
	IntervalSeconds    *big.Int
	LateFeeBps         *big.Int
	Status             uint8
}
