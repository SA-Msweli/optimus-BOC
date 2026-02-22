package eth

import (
    "context"
    "errors"
    "os"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

// NewTransactor constructs a TransactOpts object based on the PRIVATE_KEY
// environment variable and the chain ID obtained from the RPC endpoint.
// It also returns a corresponding *ethclient.Client which the caller may use
// for additional on-chain operations (e.g. waiting for receipts).
func NewTransactor(ctx context.Context, rpcURL string) (*bind.TransactOpts, *ethclient.Client, error) {
    priv := os.Getenv("PRIVATE_KEY")
    if priv == "" {
        return nil, nil, errors.New("PRIVATE_KEY environment variable required")
    }
    // parse key (allow with or without 0x prefix)
    key, err := crypto.HexToECDSA(drainHexPrefix(priv))
    if err != nil {
        return nil, nil, err
    }
    client, err := ethclient.DialContext(ctx, rpcURL)
    if err != nil {
        return nil, nil, err
    }
    chainID, err := client.ChainID(ctx)
    if err != nil {
        return nil, nil, err
    }
    auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
    if err != nil {
        return nil, nil, err
    }
    return auth, client, nil
}

func drainHexPrefix(s string) string {
    if len(s) >= 2 && s[0:2] == "0x" {
        return s[2:]
    }
    return s
}
