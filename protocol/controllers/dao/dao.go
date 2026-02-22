package dao

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"

	"github.com/optimus-boc-protocol/models"
	daosvc "github.com/optimus-boc-protocol/services/dao"
	"github.com/optimus-boc-protocol/store"
)

// Controller handles DAO management HTTP APIs. It depends on the IDAOService
// interface (Dependency Inversion), a transactor for signing, and the store
// for persisting minimal CRE-relevant data.
type Controller struct {
	svc   daosvc.IDAOService
	store *store.Store
	auth  *bind.TransactOpts
}

// NewController creates a DAO controller.
func NewController(svc daosvc.IDAOService, store *store.Store, auth *bind.TransactOpts) *Controller {
	return &Controller{svc: svc, store: store, auth: auth}
}

// Routes returns a chi.Router with all DAO endpoints.
func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	// DAO lifecycle
	r.Post("/", c.createDAO)
	r.Post("/{id}/join", c.joinDAO)

	// Governance
	r.Post("/{daoId}/propose", c.propose)
	r.Post("/proposal/{id}/vote", c.vote)
	r.Post("/proposal/{id}/finalize", c.finalize)
	r.Post("/proposal/{id}/execute", c.execute)

	// BNPL policy
	r.Post("/{id}/bnpl-terms", c.setBnplTerms)
	r.Get("/{id}/bnpl-terms", c.getBnplTerms)

	// Treasury reads
	r.Get("/{id}/treasury", c.getTreasuryBalance)
	r.Post("/{id}/treasury/credit", c.creditTreasury)

	return r
}

// ---------- DAO lifecycle ----------

func (c *Controller) createDAO(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Goal             uint8  `json:"goal"`
		VotingPeriodDays uint64 `json:"voting_period_days"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := c.svc.CreateDAO(c.auth, c.auth.From, req.Goal, req.VotingPeriodDays)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	daoId, err := c.svc.WaitForDaoID(r.Context(), tx)
	if err != nil {
		writeError(w, "unable to determine dao id: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.SaveDAO(r.Context(), models.DAO{
		DaoID:               fmt.Sprintf("%d", daoId),
		Creator:             c.auth.From.Hex(),
		Goal:                fmt.Sprintf("%d", req.Goal),
		VotingPeriodSeconds: int(req.VotingPeriodDays * 24 * 3600),
	})
	writeJSON(w, map[string]interface{}{"tx": tx.Hash().Hex(), "dao_id": daoId})
}

func (c *Controller) joinDAO(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Investment string `json:"investment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	inv := new(big.Int)
	if _, ok := inv.SetString(req.Investment, 10); !ok {
		writeError(w, "invalid investment", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.JoinDAO(c.auth, daoId, c.auth.From, inv)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

// ---------- Governance ----------

func (c *Controller) propose(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "daoId"), 10); !ok {
		writeError(w, "invalid daoId", http.StatusBadRequest)
		return
	}
	// Authorization: caller must be a member
	if ok, _ := c.svc.IsMember(r.Context(), daoId, c.auth.From); !ok {
		writeError(w, "not a member of this DAO", http.StatusForbidden)
		return
	}
	var req struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := c.svc.Propose(c.auth, daoId, common.FromHex(req.Data))
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) vote(w http.ResponseWriter, r *http.Request) {
	propId := new(big.Int)
	if _, ok := propId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid proposal id", http.StatusBadRequest)
		return
	}
	var req struct {
		DaoID   string `json:"dao_id"`
		Support bool   `json:"support"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Authorization: non-members cannot vote
	if req.DaoID != "" {
		daoId := new(big.Int)
		if _, ok := daoId.SetString(req.DaoID, 10); ok {
			if isMember, _ := c.svc.IsMember(r.Context(), daoId, c.auth.From); !isMember {
				writeError(w, "not a member of this DAO", http.StatusForbidden)
				return
			}
		}
	}
	tx, err := c.svc.Vote(c.auth, propId, req.Support)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) finalize(w http.ResponseWriter, r *http.Request) {
	propId := new(big.Int)
	if _, ok := propId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid proposal id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.FinalizeProposal(c.auth, propId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) execute(w http.ResponseWriter, r *http.Request) {
	propId := new(big.Int)
	if _, ok := propId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid proposal id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.ExecuteProposal(c.auth, propId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

// ---------- BNPL policy ----------

func (c *Controller) setBnplTerms(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		NumInstallments   string `json:"num_installments"`
		MinDays           string `json:"min_days"`
		MaxDays           string `json:"max_days"`
		LateFeeBps        string `json:"late_fee_bps"`
		GraceDays         string `json:"grace_days"`
		RescheduleAllowed bool   `json:"reschedule_allowed"`
		MinDownBps        string `json:"min_down_bps"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	toBig := func(s string) *big.Int {
		v := new(big.Int)
		v.SetString(s, 10)
		return v
	}
	tx, err := c.svc.SetBnplTerms(c.auth, daoId,
		toBig(req.NumInstallments), toBig(req.MinDays), toBig(req.MaxDays),
		toBig(req.LateFeeBps), toBig(req.GraceDays), req.RescheduleAllowed, toBig(req.MinDownBps))
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) getBnplTerms(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	terms, err := c.svc.GetBnplTerms(r.Context(), daoId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, terms)
}

// ---------- Treasury ----------

func (c *Controller) getTreasuryBalance(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	bal, err := c.svc.GetTreasuryBalance(r.Context(), daoId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"balance": bal.String()})
}

func (c *Controller) creditTreasury(w http.ResponseWriter, r *http.Request) {
	daoId := new(big.Int)
	if _, ok := daoId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Amount string `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	amt := new(big.Int)
	if _, ok := amt.SetString(req.Amount, 10); !ok {
		writeError(w, "invalid amount", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.CreditTreasury(c.auth, daoId, amt)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
