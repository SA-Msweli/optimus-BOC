package did

import (
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

// IDid defines the operations our application needs from the on-chain DID
// registry.  It is intentionally short so callers don't have to import
// concrete types; higher layers can depend on this abstraction.
// The methods below correspond to functions/events documented in the
// OPTIMUS_ETH_SPEC under the `IDIDRegistry` interface.

type IDid interface {
    CreateDID(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)
    Exists(opts *bind.CallOpts, owner common.Address) (bool, error)

    // Privy credential
    LinkPrivyCredential(opts *bind.TransactOpts, owner common.Address, privyHash [32]byte) (*types.Transaction, error)
    GetPrivyCredentialHash(opts *bind.CallOpts, owner common.Address) ([32]byte, error)
    VerifyPrivyCredential(handle string) error

    // Risk profile
    UpdateRiskProfile(opts *bind.TransactOpts, owner common.Address, newScore *big.Int, riskProfileHash [32]byte) (*types.Transaction, error)
    GetRiskProfileScore(opts *bind.CallOpts, owner common.Address) (*big.Int, error)
}
