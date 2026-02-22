// Workflow: optimus_loan_approved
//
// Trigger : EVM Log – LoanApproved(uint256,address)
// Purpose : When a loan is approved on-chain, this workflow:
//   1. Reads the full loan state including principal, rates, and terms.
//   2. Reads the current amount owed and accrued interest.
//   3. Assembles a comprehensive loan snapshot for the backend.
//   4. Notifies the backend via confidential HTTP.
package main

import (
	"fmt"
	"log/slog"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"

	"optimus-cre-workflows/shared"
)

func main() {
	wasm.NewRunner(cre.ParseJSON[shared.Config]).Run(InitWorkflow)
}

func InitWorkflow(config *shared.Config, logger *slog.Logger, _ cre.SecretsProvider) (cre.Workflow[*shared.Config], error) {
	logTrigger := evm.LogTrigger(evm.EthereumTestnetSepolia, &evm.FilterLogTriggerRequest{
		Addresses: [][]byte{shared.AddrBytes(config.LoanManagerAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigLoanApproved.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onLoanApproved)}, nil
}

func onLoanApproved(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields ──
	loanId := shared.TopicToBigInt(log.Topics[1])
	approver := shared.TopicToAddress(log.Topics[2])

	logger.Info("LoanApproved", "loanId", loanId, "approver", approver.Hex())

	// ── 2. Read full loan state ──
	loanCall, err := shared.PackGetLoan(loanId)
	if err != nil {
		return "", fmt.Errorf("pack getLoan: %w", err)
	}
	loanReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: loanCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getLoan: %w", err)
	}
	loan, err := shared.UnpackLoan(loanReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack loan: %w", err)
	}

	// ── 3. Read amount owed ──
	owedCall, err := shared.PackGetAmountOwed(loanId)
	if err != nil {
		return "", fmt.Errorf("pack getAmountOwed: %w", err)
	}
	owedReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: owedCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getAmountOwed: %w", err)
	}
	owed, err := shared.UnpackUint256(owedReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack amountOwed: %w", err)
	}

	// ── 4. Read accrued interest ──
	interestCall, err := shared.PackGetAccruedInterest(loanId)
	if err != nil {
		return "", fmt.Errorf("pack getAccruedInterest: %w", err)
	}
	interestReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.LoanManagerAddr), Data: interestCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getAccruedInterest: %w", err)
	}
	interest, err := shared.UnpackUint256(interestReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack accruedInterest: %w", err)
	}

	logger.Info("loan snapshot",
		"loanId", loanId,
		"borrower", loan.Borrower.Hex(),
		"principal", shared.FormatWei(loan.Principal),
		"amountOwed", shared.FormatWei(owed),
		"accruedInterest", shared.FormatWei(interest),
		"status", loan.Status,
	)

	// ── 5. Notify backend with full loan snapshot ──
	body := fmt.Sprintf(
		`{"loanId":"%s","borrower":"%s","approver":"%s","principal":"%s","interestRateBps":"%s","startTime":"%s","endTime":"%s","amountOwed":"%s","accruedInterest":"%s","amountPaid":"%s","status":"ACTIVE"}`,
		loanId, loan.Borrower.Hex(), approver.Hex(),
		loan.Principal.String(), loan.InterestRateBps.String(),
		loan.StartTime.String(), loan.EndTime.String(),
		owed.String(), interest.String(), loan.AmountPaid.String(),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/loans/approved", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed LoanApproved", "loanId", loanId)
	return "OK", nil
}
