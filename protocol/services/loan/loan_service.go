package loan

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/optimus-boc-protocol/bindings"
)

// Service is the production implementation of ILoanService.
type Service struct {
	client   *ethclient.Client
	contract *bindings.LoanManager
}

// NewService constructs a Loan service from an RPC URL and contract address.
func NewService(rpcURL, contractAddr string) (ILoanService, error) {
	cli, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("loan dial: %w", err)
	}
	addr := common.HexToAddress(contractAddr)
	c, err := bindings.NewLoanManager(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("loan binding: %w", err)
	}
	return &Service{client: cli, contract: c}, nil
}

func (s *Service) CreateLoan(opts *bind.TransactOpts, borrower common.Address, daoId *big.Int, principal *big.Int, interestRateBps *big.Int, durationSeconds *big.Int) (*types.Transaction, error) {
	return s.contract.CreateLoan(opts, borrower, daoId, principal, interestRateBps, durationSeconds)
}

func (s *Service) ApproveLoan(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return s.contract.ApproveLoan(opts, loanId)
}

func (s *Service) MakePayment(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return s.contract.MakePayment(opts, loanId)
}

func (s *Service) MarkDefaulted(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return s.contract.MarkDefaulted(opts, loanId)
}

func (s *Service) GetLoan(ctx context.Context, id *big.Int) (LoanInfo, error) {
	out, err := s.contract.GetLoan(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return LoanInfo{}, err
	}
	return LoanInfo{
		Id:              out.Id,
		Borrower:        out.Borrower,
		Principal:       out.Principal,
		InterestRateBps: out.InterestRateBps,
		StartTime:       out.StartTime,
		EndTime:         out.EndTime,
		AmountPaid:      out.AmountPaid,
		Status:          out.Status,
	}, nil
}

func (s *Service) GetAccruedInterest(ctx context.Context, loanId *big.Int) (*big.Int, error) {
	return s.contract.GetAccruedInterest(&bind.CallOpts{Context: ctx}, loanId)
}

func (s *Service) GetAmountOwed(ctx context.Context, loanId *big.Int) (*big.Int, error) {
	return s.contract.GetAmountOwed(&bind.CallOpts{Context: ctx}, loanId)
}

func (s *Service) WaitForLoanID(ctx context.Context, tx *types.Transaction) (uint64, error) {
	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return 0, err
	}
	for _, log := range receipt.Logs {
		ev, err := s.contract.ParseLoanCreated(*log)
		if err == nil {
			return ev.LoanId.Uint64(), nil
		}
	}
	return 0, fmt.Errorf("LoanCreated event not found in receipt %s", tx.Hash().Hex())
}
