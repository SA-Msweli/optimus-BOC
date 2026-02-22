package did

import (
    "bytes"
    "fmt"
    "math/big"
    "net/http"
    "os"

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


// CreateDID calls the corresponding contract method.
func (s *Did) CreateDID(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
    return s.contract.CreateDID(opts, owner)
}

// Exists checks whether a DID already exists for the given owner.
func (s *Did) Exists(opts *bind.CallOpts, owner common.Address) (bool, error) {
    return s.contract.Exists(opts, owner)
}

// LinkPrivyCredential stores a hashed Privy credential pointer on-chain.
// see OPTIMUS_ETH_SPEC `linkPrivyCredential`.
// It is expected that callers have already verified the credential with
// Privy (e.g. via VerifyPrivyCredential below or in a CRE workflow).
func (s *Did) LinkPrivyCredential(opts *bind.TransactOpts, owner common.Address, privyHash [32]byte) (*types.Transaction, error) {
    return s.contract.LinkPrivyCredential(opts, owner, privyHash)
}

// VerifyPrivyCredential makes a call to the Privy API using the configured
// app secret and API key to ensure the provided handle is valid.  This is a
// simple illustration; real validation may require POSTing a JWT or other
// data according to the Privy spec.
func (s *Did) VerifyPrivyCredential(handle string) error {
    url := fmt.Sprintf("https://api.privy.io/v1/credentials/%s/verify", handle)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }
    appSecret := os.Getenv("PRIVY_APP_SECRET")
    if appSecret != "" {
        req.Header.Set("Authorization", "Bearer "+appSecret)
    }

		resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        var buf bytes.Buffer
        buf.ReadFrom(resp.Body)
        return fmt.Errorf("privy verification failed: %s", buf.String())
    }
    return nil
}

// GetPrivyCredentialHash returns the stored hash (or zero if none).
func (s *Did) GetPrivyCredentialHash(opts *bind.CallOpts, owner common.Address) ([32]byte, error) {
    return s.contract.GetPrivyCredentialHash(opts, owner)
}

// Client returns the underlying contract backend, useful for event filtering
// or waiting for receipts.
func (s *Did) Client() bind.ContractBackend {
    return s.client
}

// UpdateRiskProfile calls the on-chain updateRiskProfile function.
func (s *Did) UpdateRiskProfile(opts *bind.TransactOpts, owner common.Address, newScore *big.Int, riskProfileHash [32]byte) (*types.Transaction, error) {
    return s.contract.UpdateRiskProfile(opts, owner, newScore, riskProfileHash)
}

// GetRiskProfileScore returns the on-chain risk profile score for an owner.
func (s *Did) GetRiskProfileScore(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
    return s.contract.GetRiskProfileScore(opts, owner)
}
