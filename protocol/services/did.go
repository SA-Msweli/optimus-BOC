package services

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/optimus-boc-protocol/bindings"
)

// DIDService wraps DIDRegistry interactions.

type DIDService struct {
    client *ethclient.Client
    contract *bindings.DIDRegistry
}

func NewDIDService(rpcUrl, contractAddr string) (*DIDService, error) {
    cli, err := ethclient.Dial(rpcUrl)
    if err != nil {
        return nil, err
    }
    addr := common.HexToAddress(contractAddr)
    c, err := bindings.NewDIDRegistry(addr, cli)
    if err != nil {
        return nil, err
    }
    return &DIDService{client: cli, contract: c}, nil
}

// Add DID-specific helper methods here
