package services

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/optimus-boc-protocol/bindings"
)

// TokenVaultService wraps the TokenVault contract.

type TokenVaultService struct {
    client *ethclient.Client
    contract *bindings.TokenVault
}

func NewTokenVaultService(rpcUrl, contractAddr string) (*TokenVaultService, error) {
    cli, err := ethclient.Dial(rpcUrl)
    if err != nil {
        return nil, err
    }
    addr := common.HexToAddress(contractAddr)
    c, err := bindings.NewTokenVault(addr, cli)
    if err != nil {
        return nil, err
    }
    return &TokenVaultService{client: cli, contract: c}, nil
}

// Add TokenVault-specific methods here
