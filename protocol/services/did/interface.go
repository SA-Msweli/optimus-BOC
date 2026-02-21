package did

import (
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

// IDid defines the operations our application needs from the on-chain DID
// registry.  It is intentionally short so callers don't have to import
// concrete types; higher layers can depend on this abstraction.
// The long name 'IDid' follows Go convention of prefixing interfaces with
// 'I' when the concrete implementation is named `Did`.
type IDid interface {
    CreateDID(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)
    Exists(opts *bind.CallOpts, owner common.Address) (bool, error)
}
