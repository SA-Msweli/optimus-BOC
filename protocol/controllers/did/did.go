package did

import (
    "encoding/json"
    "net/http"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/go-chi/chi/v5"

    "github.com/optimus-boc-protocol/services/did"
)

// Controller handles HTTP routes related to DIDs.  It depends only on the
// DIDService interface, allowing the implementation to be swapped in tests.
type Controller struct {
    svc did.IDid
}

// NewController constructs a DID controller.
func NewController(svc did.IDid) *Controller {
    return &Controller{svc: svc}
}

// Routes returns a router mounting the controller's endpoints.
func (c *Controller) Routes() chi.Router {
    r := chi.NewRouter()
    r.Post("/did", c.createDID)
    r.Get("/did/{owner}", c.lookupDID)
    return r
}

// createDID POST /did body {"owner":"0x..."}
func (c *Controller) createDID(w http.ResponseWriter, r *http.Request) {
    var req struct{ Owner string `json:"owner"` }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    addr := common.HexToAddress(req.Owner)
    // transaction options would normally be obtained from a signer
    tx, err := c.svc.CreateDID(nil, addr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(tx.Hash().Hex())
}

// lookupDID GET /did/{owner}
func (c *Controller) lookupDID(w http.ResponseWriter, r *http.Request) {
    owner := chi.URLParam(r, "owner")
    addr := common.HexToAddress(owner)
    exists, err := c.svc.Exists(&bind.CallOpts{Context: r.Context()}, addr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]bool{"exists": exists})
}
