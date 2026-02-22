package did

import (
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"

	"github.com/optimus-boc-protocol/models"
	"github.com/optimus-boc-protocol/services/did"
	"github.com/optimus-boc-protocol/store"
)

// Controller handles HTTP routes related to DIDs. It depends only on the
// IDid interface, allowing the implementation to be swapped.
type Controller struct {
	svc   did.IDid
	store *store.Store
	auth  *bind.TransactOpts
}

// NewController constructs a DID controller.
func NewController(svc did.IDid, store *store.Store, auth *bind.TransactOpts) *Controller {
	return &Controller{svc: svc, store: store, auth: auth}
}

// Routes returns a router mounting the controller's endpoints.
func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.createDID)
	r.Get("/{owner}", c.lookupDID)
	r.Post("/{owner}/link", c.linkPrivy)
	r.Get("/{owner}/privy", c.getPrivyHash)
	r.Post("/{owner}/risk", c.updateRiskProfile)
	r.Get("/{owner}/risk", c.getRiskProfileScore)
	return r
}

func (c *Controller) createDID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Owner string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	addr := common.HexToAddress(req.Owner)
	tx, err := c.svc.CreateDID(c.auth, addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if c.store != nil {
		_ = c.store.SaveDIDProfile(r.Context(), models.DIDProfile{Owner: req.Owner, RiskScore: 0, LastActive: time.Now().Unix()})
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) lookupDID(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "owner")
	addr := common.HexToAddress(owner)
	exists, err := c.svc.Exists(&bind.CallOpts{Context: r.Context()}, addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]bool{"exists": exists})
}

func (c *Controller) linkPrivy(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "owner")
	addr := common.HexToAddress(owner)
	var req struct {
		Hash string `json:"hash"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	var h [32]byte
	copy(h[:], common.Hex2Bytes(req.Hash))
	if err := c.svc.VerifyPrivyCredential(req.Hash); err != nil {
		writeError(w, "invalid privy credential: "+err.Error(), http.StatusBadRequest)
		return
	}
	if c.store != nil {
		_ = c.store.SaveDIDProfile(r.Context(), models.DIDProfile{Owner: owner, LastActive: time.Now().Unix()})
	}
	tx, err := c.svc.LinkPrivyCredential(c.auth, addr, h)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) getPrivyHash(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "owner")
	addr := common.HexToAddress(owner)
	hash, err := c.svc.GetPrivyCredentialHash(&bind.CallOpts{Context: r.Context()}, addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"hash": common.Bytes2Hex(hash[:])})
}

func (c *Controller) updateRiskProfile(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "owner")
	addr := common.HexToAddress(owner)
	var req struct {
		NewScore        string `json:"new_score"`
		RiskProfileHash string `json:"risk_profile_hash"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	score := new(big.Int)
	if _, ok := score.SetString(req.NewScore, 10); !ok {
		writeError(w, "invalid new_score", http.StatusBadRequest)
		return
	}
	var h [32]byte
	copy(h[:], common.Hex2Bytes(req.RiskProfileHash))
	tx, err := c.svc.UpdateRiskProfile(c.auth, addr, score, h)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if c.store != nil {
		_ = c.store.SaveDIDProfile(r.Context(), models.DIDProfile{Owner: owner, RiskScore: int(score.Int64()), LastActive: time.Now().Unix()})
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) getRiskProfileScore(w http.ResponseWriter, r *http.Request) {
	owner := chi.URLParam(r, "owner")
	addr := common.HexToAddress(owner)
	score, err := c.svc.GetRiskProfileScore(&bind.CallOpts{Context: r.Context()}, addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"risk_score": score.String()})
}

// ---------- helpers ----------

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
