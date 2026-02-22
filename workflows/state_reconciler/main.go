// Workflow: optimus_state_reconciler
//
// Trigger : Cron – every 2 hours
// Purpose : Cross-checks on-chain state with the backend database for all
//   active entities (BNPL arrangements, loans, DAO treasuries).  Reports
//   any mismatches to the backend for reconciliation.
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"

	"optimus-cre-workflows/shared"
)

func main() {
	wasm.NewRunner(cre.ParseJSON[shared.Config]).Run(InitWorkflow)
}

func InitWorkflow(config *shared.Config, logger *slog.Logger, _ cre.SecretsProvider) (cre.Workflow[*shared.Config], error) {
	workflow := cre.Workflow[*shared.Config]{
		cre.Handler(
			cron.Trigger(&cron.Config{Schedule: config.Schedule}),
			onCronReconcile,
		),
	}
	return workflow, nil
}

func onCronReconcile(config *shared.Config, runtime cre.Runtime, _ *cron.Payload) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}
	mismatches := 0

	// ═══════════════════════════════════════════
	// Section 1: BNPL Arrangements
	// ═══════════════════════════════════════════
	bnplResp, err := shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/bnpl/active-ids", Method: "GET", Owner: config.Owner},
		runtime,
	)
	if err == nil {
		arrIDs := parseIDList([]byte(bnplResp.Body))
		for _, id := range arrIDs {
			arrID := big.NewInt(id)
			arrCall, _ := shared.PackGetArrangement(arrID)
			arrReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
				Call: &evm.CallMsg{To: shared.AddrBytes(config.BNPLManagerAddr), Data: arrCall},
			}).Await()
			if err != nil {
				continue
			}
			arr, err := shared.UnpackArrangement(arrReply.Data)
			if err != nil {
				continue
			}

			body := fmt.Sprintf(`{"arrangementId":"%s","status":%d,"totalAmount":"%s","numInstallments":"%s"}`,
				arrID, arr.Status, arr.TotalAmount, arr.NumInstallments)
			resp, _ := shared.ConfidentialRequest(
				shared.BackendReq{URL: config.BackendURL + "/api/reconcile/bnpl", Method: "POST", Body: body, Owner: config.Owner},
				runtime,
			)
			if containsMismatch([]byte(resp.Body)) {
				mismatches++
			}
		}
	}

	// ═══════════════════════════════════════════
	// Section 2: Loans
	// ═══════════════════════════════════════════
	loanResp, err := shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/loans/active-ids", Method: "GET", Owner: config.Owner},
		runtime,
	)
	if err == nil {
		loanIDs := parseIDList([]byte(loanResp.Body))
		for _, id := range loanIDs {
			loanID := big.NewInt(id)

			loanCall, _ := shared.PackGetLoan(loanID)
			loanReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
				Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: loanCall},
			}).Await()
			if err != nil {
				continue
			}
			loan, err := shared.UnpackLoan(loanReply.Data)
			if err != nil {
				continue
			}

			owedCall, _ := shared.PackGetAmountOwed(loanID)
			owedReply, _ := evmClient.CallContract(runtime, &evm.CallContractRequest{
				Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: owedCall},
			}).Await()
			amountOwed := big.NewInt(0)
			if owedReply != nil {
				val, err := shared.UnpackUint256(owedReply.Data)
				if err == nil {
					amountOwed = val
				}
			}

			body := fmt.Sprintf(`{"loanId":"%s","status":%d,"principal":"%s","amountPaid":"%s","amountOwed":"%s"}`,
				loanID, loan.Status, loan.Principal, loan.AmountPaid, amountOwed)
			resp, _ := shared.ConfidentialRequest(
				shared.BackendReq{URL: config.BackendURL + "/api/reconcile/loan", Method: "POST", Body: body, Owner: config.Owner},
				runtime,
			)
			if containsMismatch([]byte(resp.Body)) {
				mismatches++
			}
		}
	}

	// ═══════════════════════════════════════════
	// Section 3: DAO Treasuries
	// ═══════════════════════════════════════════
	daoResp, err := shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/dao/active-ids", Method: "GET", Owner: config.Owner},
		runtime,
	)
	if err == nil {
		daoIDs := parseIDList([]byte(daoResp.Body))
		for _, id := range daoIDs {
			daoID := big.NewInt(id)
			tCall, _ := shared.PackGetTreasuryBalance(daoID)
			tReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
				Call: &evm.CallMsg{To: shared.AddrBytes(config.DAOManagerAddr), Data: tCall},
			}).Await()
			if err != nil {
				continue
			}
			daoBalance, err := shared.UnpackUint256(tReply.Data)
			if err != nil {
				continue
			}

			body := fmt.Sprintf(`{"daoId":"%s","treasuryBalance":"%s"}`, daoID, daoBalance)
			resp, _ := shared.ConfidentialRequest(
				shared.BackendReq{URL: config.BackendURL + "/api/reconcile/dao-treasury", Method: "POST", Body: body, Owner: config.Owner},
				runtime,
			)
			if containsMismatch([]byte(resp.Body)) {
				mismatches++
			}
		}
	}

	logger.Info("state reconciliation complete", "mismatches", mismatches)
	return fmt.Sprintf("mismatches=%d", mismatches), nil
}

func containsMismatch(body []byte) bool {
	var data struct {
		Mismatch bool `json:"mismatch"`
	}
	_ = json.Unmarshal(body, &data)
	return data.Mismatch
}

func parseIDList(body []byte) []int64 {
	var ids []int64
	_ = json.Unmarshal(body, &ids)
	return ids
}
