package tokenvault

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ITokenVaultService defines the contract for TokenVault operations.
// Matches the on-chain ITokenVault interface from OPTIMUS_ETH_SPEC.

type ITokenVaultService interface {
	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)
	Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)
	GetBalance(ctx context.Context, token common.Address) (*big.Int, error)
}
