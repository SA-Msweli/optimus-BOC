package loan

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ILoanService defines the contract for Loan operations.
// Matches the on-chain ILoanManager interface from OPTIMUS_ETH_SPEC.

type ILoanService interface {
	// Write operations
	CreateLoan(opts *bind.TransactOpts, borrower common.Address, daoId *big.Int, principal *big.Int, interestRateBps *big.Int, durationSeconds *big.Int) (*types.Transaction, error)
	ApproveLoan(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error)
	MakePayment(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error)
	MarkDefaulted(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error)

	// Read operations
	GetLoan(ctx context.Context, id *big.Int) (LoanInfo, error)
	GetAccruedInterest(ctx context.Context, loanId *big.Int) (*big.Int, error)
	GetAmountOwed(ctx context.Context, loanId *big.Int) (*big.Int, error)

	// Helpers
	WaitForLoanID(ctx context.Context, tx *types.Transaction) (uint64, error)
}

// LoanInfo is a friendlier representation of the on-chain Loan struct.
type LoanInfo struct {
	Id              *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	AmountPaid      *big.Int
	Status          uint8
}
