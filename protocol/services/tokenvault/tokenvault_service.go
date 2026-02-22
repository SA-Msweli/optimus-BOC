package tokenvault

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

// Service is the production implementation of ITokenVaultService.
type Service struct {
	client   *ethclient.Client
	contract *bindings.TokenVault
}

// NewService constructs a TokenVault service from an RPC URL and contract address.
func NewService(rpcURL, contractAddr string) (ITokenVaultService, error) {
	cli, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("tokenvault dial: %w", err)
	}
	addr := common.HexToAddress(contractAddr)
	c, err := bindings.NewTokenVault(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("tokenvault binding: %w", err)
	}
	return &Service{client: cli, contract: c}, nil
}

func (s *Service) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return s.contract.Deposit(opts, token, amount)
}

func (s *Service) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return s.contract.Withdraw(opts, token, amount)
}

func (s *Service) GetBalance(ctx context.Context, token common.Address) (*big.Int, error) {
	return s.contract.GetBalance(&bind.CallOpts{Context: ctx}, token)
}
