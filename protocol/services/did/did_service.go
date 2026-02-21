package did

import (
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/optimus-boc-protocol/bindings"
)

// Did is the production implementation of IDid that talks to a real
// DIDRegistry contract.  The shorter concrete name matches the interface
// prefix ('I').
type Did struct {
    client   bind.ContractBackend
    contract *bindings.DIDRegistry
}

// NewDid constructs a Did instance given an RPC URL and contract address.
// It returns the interface type so callers remain decoupled.
func NewDid(rpcUrl, contractAddr string) (IDid, error) {
    cli, err := ethclient.Dial(rpcUrl)
    if err != nil {
        return nil, err
    }
    addr := common.HexToAddress(contractAddr)
    c, err := bindings.NewDIDRegistry(addr, cli)
    if err != nil {
        return nil, err
    }
    return &Did{client: cli, contract: c}, nil
}

// NewDidFromContract constructs a Did instance from an existing registry
// pointer; it is primarily used in tests.
func NewDidFromContract(client bind.ContractBackend, contract *bindings.DIDRegistry) IDid {
    return &Did{client: client, contract: contract}
}

// CreateDID calls the corresponding contract method.
func (s *Did) CreateDID(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
    return s.contract.CreateDID(opts, owner)
}

// Exists checks whether a DID already exists for the given owner.
func (s *Did) Exists(opts *bind.CallOpts, owner common.Address) (bool, error) {
    return s.contract.Exists(opts, owner)
}

// more helper implementations can be added here
