package bnpl

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"

	"github.com/optimus-boc-protocol/models"
	bnplsvc "github.com/optimus-boc-protocol/services/bnpl"
	"github.com/optimus-boc-protocol/store"
)

// Controller handles BNPL HTTP routes. It depends on the IBNPLService
// interface (Dependency Inversion) and the Store for off-chain persistence
// needed by CRE workflows.
type Controller struct {
	svc   bnplsvc.IBNPLService
	store *store.Store
	auth  *bind.TransactOpts
}

// NewController creates a BNPL controller.
func NewController(svc bnplsvc.IBNPLService, store *store.Store, auth *bind.TransactOpts) *Controller {
	return &Controller{svc: svc, store: store, auth: auth}
}

// Routes returns the HTTP routes handled by the BNPL controller.
func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/arrangements", c.createArrangement)
	r.Get("/arrangements/{id}", c.getArrangement)
	r.Post("/arrangements/{id}/payment", c.makePayment)
	r.Post("/arrangements/{id}/activate", c.activate)
	r.Post("/arrangements/{id}/latefee", c.applyLateFee)
	r.Post("/arrangements/{id}/reschedule", c.reschedule)
	return r
}

// ---------- handlers ----------

func (c *Controller) createArrangement(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DaoID           string `json:"dao_id"`
		Recipient       string `json:"recipient"`
		TotalAmount     string `json:"total_amount"`
		StartTimestamp  int64  `json:"start_timestamp"`
		IntervalSeconds int64  `json:"interval_seconds"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	daoId := new(big.Int)
	if _, ok := daoId.SetString(req.DaoID, 10); !ok {
		writeError(w, "invalid dao_id", http.StatusBadRequest)
		return
	}
	tot := new(big.Int)
	if _, ok := tot.SetString(req.TotalAmount, 10); !ok {
		writeError(w, "invalid total_amount", http.StatusBadRequest)
		return
	}
	recipient := common.HexToAddress(req.Recipient)

	tx, err := c.svc.CreateBNPL(c.auth, daoId, recipient, tot, big.NewInt(req.StartTimestamp), big.NewInt(req.IntervalSeconds), nil)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	arrID, err := c.svc.WaitForArrangementID(r.Context(), tx)
	if err != nil {
		writeError(w, "tx mined but could not parse event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_ = c.store.SaveArrangement(r.Context(), models.Arrangement{
		ArrangementID:   fmt.Sprintf("%d", arrID),
		DaoID:           req.DaoID,
		Payer:           c.auth.From.Hex(),
		Recipient:       req.Recipient,
		TotalAmount:     req.TotalAmount,
		NumInstallments: 0,
		StartTimestamp:  req.StartTimestamp,
		IntervalSeconds: req.IntervalSeconds,
		Status:          "PENDING",
	})

	writeJSON(w, map[string]interface{}{"tx": tx.Hash().Hex(), "arrangement_id": arrID})
}

func (c *Controller) getArrangement(w http.ResponseWriter, r *http.Request) {
	id := new(big.Int)
	if _, ok := id.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	arr, err := c.svc.GetArrangement(r.Context(), id)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, arr)
}

func (c *Controller) makePayment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	arrId := new(big.Int)
	if _, ok := arrId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Installment uint8 `json:"installment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := c.svc.MakePayment(c.auth, arrId, req.Installment)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateArrangementStatus(r.Context(), idStr, "PAYMENT_MADE")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) activate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	arrId := new(big.Int)
	if _, ok := arrId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.ActivateBNPL(c.auth, arrId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateArrangementStatus(r.Context(), idStr, "ACTIVE")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) applyLateFee(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	arrId := new(big.Int)
	if _, ok := arrId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Installment uint8 `json:"installment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := c.svc.ApplyLateFee(c.auth, arrId, req.Installment)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateArrangementStatus(r.Context(), idStr, "LATE_FEE_APPLIED")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) reschedule(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	arrId := new(big.Int)
	if _, ok := arrId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		NewStart    int64 `json:"new_start_timestamp"`
		NewInterval int64 `json:"new_interval_seconds"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := c.svc.Reschedule(c.auth, arrId, big.NewInt(req.NewStart), big.NewInt(req.NewInterval))
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateArrangementStatus(r.Context(), idStr, "RESCHEDULED")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
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
