// Workflow: optimus_loan_default_monitor
//
// Trigger : Cron – every hour
// Purpose : Scans active loans past their end time.  For each defaultable loan:
//   1. Verifies outstanding balance > 0 via LoanManager.getAmountOwed.
//   2. Calls LoanManager.markDefaulted on-chain.
//   3. Slashes the borrower's DID risk score (-2000).
//   4. Notifies the backend via confidential HTTP.
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
			onCronDefaultCheck,
		),
	}
	return workflow, nil
}

func onCronDefaultCheck(config *shared.Config, runtime cre.Runtime, _ *cron.Payload) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}
	now := time.Now().Unix()

	// ── 1. Resolve active loan IDs ──
	loanIDs := config.ActiveLoanIDs
	if len(loanIDs) == 0 {
		resp, err := shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + "/api/loans/active-ids", Method: "GET", Owner: config.Owner},
			runtime,
		)
		if err != nil {
			return "", fmt.Errorf("fetch active loan IDs: %w", err)
		}
		loanIDs = parseIDList([]byte(resp.Body))
	}

	defaulted := 0
	for _, id := range loanIDs {
		loanID := big.NewInt(id)

		// ── 2. Read loan state ──
		loanCall, err := shared.PackGetLoan(loanID)
		if err != nil {
			continue
		}
		loanReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
			Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: loanCall},
		}).Await()
		if err != nil {
			logger.Warn("getLoan failed", "id", id, "err", err)
			continue
		}
		loan, err := shared.UnpackLoan(loanReply.Data)
		if err != nil || loan.Status != shared.LoanStatusActive {
			continue
		}

		// Only default if past end time
		if now <= loan.EndTime.Int64() {
			continue
		}

		// ── 3. Read outstanding amount ──
		owedCall, _ := shared.PackGetAmountOwed(loanID)
		owedReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
			Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: owedCall},
		}).Await()
		if err != nil {
			continue
		}
		amountOwed, err := shared.UnpackUint256(owedReply.Data)
		if err != nil || amountOwed.Sign() == 0 {
			continue // already repaid
		}

		// ── 4. Mark defaulted on-chain ──
		defaultCall, _ := shared.PackMarkDefaulted(loanID)
		defaultReport, err := runtime.GenerateReport(&cre.ReportRequest{
			EncodedPayload: defaultCall,
		}).Await()
		if err != nil {
			logger.Warn("generateReport markDefaulted failed", "loan", id, "err", err)
			continue
		}
		writeReply, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver: shared.AddrBytes(config.LoanManagerAddr),
			Report:   defaultReport,
		}).Await()
		if err != nil || writeReply.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
			logger.Warn("markDefaulted tx failed", "loan", id, "err", err)
			continue
		}

		// ── 5. Slash risk score ──
		adjustRisk(evmClient, runtime, config, loan.Borrower, shared.AdjLoanDefaulted, "loan_defaulted", logger)

		// ── 6. Notify backend ──
		body := fmt.Sprintf(`{"loanId":"%s","borrower":"%s","amountOwed":"%s"}`,
			loanID, loan.Borrower.Hex(), amountOwed)
		_, _ = shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + fmt.Sprintf("/api/loans/%d/defaulted", id), Method: "POST", Body: body, Owner: config.Owner},
			runtime,
		)

		defaulted++
	}

	logger.Info("default monitor complete", "defaulted", defaulted)
	return fmt.Sprintf("defaulted=%d", defaulted), nil
}

func adjustRisk(evmClient *evm.Client, runtime cre.Runtime, config *shared.Config, owner common.Address, delta int64, reason string, logger *slog.Logger) {
	scoreCall, err := shared.PackGetRiskProfileScore(owner)
	if err != nil {
		return
	}
	scoreReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DIDRegistryAddr), Data: scoreCall},
	}).Await()
	if err != nil {
		return
	}
	currentScore, err := shared.UnpackUint256(scoreReply.Data)
	if err != nil {
		return
	}

	newScore := shared.AdjustScore(currentScore, delta)
	profileHash := shared.ComputeProfileHash(owner, newScore, reason)
	updateCall, err := shared.PackUpdateRiskProfile(owner, newScore, profileHash)
	if err != nil {
		return
	}

	riskReport, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: updateCall,
	}).Await()
	if err != nil {
		logger.Warn("generateReport updateRiskProfile failed", "owner", owner.Hex(), "err", err)
		return
	}

	_, err = evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: shared.AddrBytes(config.DIDRegistryAddr),
		Report:   riskReport,
	}).Await()
	if err != nil {
		logger.Warn("updateRiskProfile failed", "owner", owner.Hex(), "err", err)
	}
}

func parseIDList(body []byte) []int64 {
	var ids []int64
	_ = json.Unmarshal(body, &ids)
	return ids
}
