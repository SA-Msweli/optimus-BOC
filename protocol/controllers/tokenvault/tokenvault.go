package tokenvault

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"

	tvsvc "github.com/optimus-boc-protocol/services/tokenvault"
)

// Controller handles TokenVault HTTP routes. It depends on the
// ITokenVaultService interface (Dependency Inversion).
type Controller struct {
	svc  tvsvc.ITokenVaultService
	auth *bind.TransactOpts
}

// NewController creates a TokenVault controller.
func NewController(svc tvsvc.ITokenVaultService, auth *bind.TransactOpts) *Controller {
	return &Controller{svc: svc, auth: auth}
}

// Routes returns the chi.Router for token vault endpoints.
func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/deposit", c.deposit)
	r.Post("/withdraw", c.withdraw)
	r.Get("/balance/{token}", c.getBalance)
	return r
}

// ---------- handlers ----------

func (c *Controller) deposit(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	token := common.HexToAddress(req.Token)
	amt := new(big.Int)
	if _, ok := amt.SetString(req.Amount, 10); !ok {
		writeError(w, "invalid amount", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.Deposit(c.auth, token, amt)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) withdraw(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	token := common.HexToAddress(req.Token)
	amt := new(big.Int)
	if _, ok := amt.SetString(req.Amount, 10); !ok {
		writeError(w, "invalid amount", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.Withdraw(c.auth, token, amt)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) getBalance(w http.ResponseWriter, r *http.Request) {
	token := common.HexToAddress(chi.URLParam(r, "token"))
	bal, err := c.svc.GetBalance(r.Context(), token)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"balance": bal.String()})
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
