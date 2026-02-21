package services

import (
    "context"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/optimus-boc-protocol/bindings"
)

// LoanService wraps interactions with LoanManager contract.

type LoanService struct {
    client *ethclient.Client
    contract *bindings.LoanManager
}

func NewLoanService(rpcUrl, contractAddr string) (*LoanService, error) {
    cli, err := ethclient.Dial(rpcUrl)
    if err != nil {
        return nil, err
    }
    addr := common.HexToAddress(contractAddr)
    c, err := bindings.NewLoanManager(addr, cli)
    if err != nil {
        return nil, err
    }
    return &LoanService{client: cli, contract: c}, nil
}

// Example method, fleshed out later
func (s *LoanService) GetLoan(ctx context.Context, id *big.Int) (struct {
    Id *big.Int
    Borrower common.Address
    Principal *big.Int
    InterestRateBps *big.Int
    StartTime *big.Int
    EndTime *big.Int
    AmountPaid *big.Int
    Status uint8
}, error) {
    return s.contract.GetLoan(&bind.CallOpts{Context: ctx}, id)
}
