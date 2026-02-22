package loan

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"

	"github.com/optimus-boc-protocol/models"
	loansvc "github.com/optimus-boc-protocol/services/loan"
	"github.com/optimus-boc-protocol/store"
)

// Controller handles Loan HTTP routes. It depends on the ILoanService
// interface and the Store for off-chain CRE persistence.
type Controller struct {
	svc   loansvc.ILoanService
	store *store.Store
	auth  *bind.TransactOpts
}

// NewController creates a Loan controller.
func NewController(svc loansvc.ILoanService, store *store.Store, auth *bind.TransactOpts) *Controller {
	return &Controller{svc: svc, store: store, auth: auth}
}

// Routes returns the chi.Router for loan endpoints.
func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()
	// Write
	r.Post("/", c.createLoan)
	r.Post("/{id}/approve", c.approveLoan)
	r.Post("/{id}/payment", c.makePayment)
	r.Post("/{id}/default", c.markDefaulted)
	// Read
	r.Get("/{id}", c.getLoan)
	r.Get("/{id}/interest", c.getAccruedInterest)
	r.Get("/{id}/owed", c.getAmountOwed)
	return r
}

// ---------- handlers ----------

func (c *Controller) createLoan(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Borrower        string `json:"borrower"`
		DaoID           string `json:"dao_id"`
		Principal       string `json:"principal"`
		InterestRateBps string `json:"interest_rate_bps"`
		DurationSeconds string `json:"duration_seconds"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	borrower := common.HexToAddress(req.Borrower)
	daoId := new(big.Int)
	if _, ok := daoId.SetString(req.DaoID, 10); !ok {
		writeError(w, "invalid dao_id", http.StatusBadRequest)
		return
	}
	principal := new(big.Int)
	if _, ok := principal.SetString(req.Principal, 10); !ok {
		writeError(w, "invalid principal", http.StatusBadRequest)
		return
	}
	rateBps := new(big.Int)
	if _, ok := rateBps.SetString(req.InterestRateBps, 10); !ok {
		writeError(w, "invalid interest_rate_bps", http.StatusBadRequest)
		return
	}
	dur := new(big.Int)
	if _, ok := dur.SetString(req.DurationSeconds, 10); !ok {
		writeError(w, "invalid duration_seconds", http.StatusBadRequest)
		return
	}

	tx, err := c.svc.CreateLoan(c.auth, borrower, daoId, principal, rateBps, dur)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	loanID, err := c.svc.WaitForLoanID(r.Context(), tx)
	if err != nil {
		writeError(w, "tx mined but could not parse event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_ = c.store.SaveLoan(r.Context(), models.Loan{
		LoanID:          fmt.Sprintf("%d", loanID),
		Borrower:        req.Borrower,
		DaoAddress:      req.DaoID,
		Principal:       req.Principal,
		InterestRateBps: int(rateBps.Int64()),
		Status:          "PENDING",
	})

	writeJSON(w, map[string]interface{}{"tx": tx.Hash().Hex(), "loan_id": loanID})
}

func (c *Controller) approveLoan(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	idStr := chi.URLParam(r, "id")
	if _, ok := loanId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.ApproveLoan(c.auth, loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateLoanStatus(r.Context(), idStr, "APPROVED")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) makePayment(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	idStr := chi.URLParam(r, "id")
	if _, ok := loanId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.MakePayment(c.auth, loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateLoanStatus(r.Context(), idStr, "PAYMENT_MADE")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) markDefaulted(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	idStr := chi.URLParam(r, "id")
	if _, ok := loanId.SetString(idStr, 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	tx, err := c.svc.MarkDefaulted(c.auth, loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = c.store.UpdateLoanStatus(r.Context(), idStr, "DEFAULTED")
	writeJSON(w, map[string]string{"tx": tx.Hash().Hex()})
}

func (c *Controller) getLoan(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	if _, ok := loanId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	info, err := c.svc.GetLoan(r.Context(), loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, info)
}

func (c *Controller) getAccruedInterest(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	if _, ok := loanId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	interest, err := c.svc.GetAccruedInterest(r.Context(), loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"accrued_interest": interest.String()})
}

func (c *Controller) getAmountOwed(w http.ResponseWriter, r *http.Request) {
	loanId := new(big.Int)
	if _, ok := loanId.SetString(chi.URLParam(r, "id"), 10); !ok {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}
	owed, err := c.svc.GetAmountOwed(r.Context(), loanId)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"amount_owed": owed.String()})
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
