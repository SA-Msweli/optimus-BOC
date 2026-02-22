// Workflow: optimus_loan_payment
//
// Trigger : EVM Log – PaymentMade(uint256,address,uint256,uint256,uint256)
// Purpose : When a loan payment is made on-chain, this workflow:
//   1. Decodes the payment event data (amount, total paid, remaining).
//   2. Reads the borrower's current DID risk score.
//   3. Checks remaining interest to determine if loan is fully repaid.
//   4. Adjusts risk score (+100 per payment, +700 if full repayment).
//   5. Writes the updated risk profile on-chain.
//   6. Notifies the backend via confidential HTTP.
package main

import (
	"fmt"
	"log/slog"
	"math/big"

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
			{Values: [][]byte{shared.SigPaymentMade.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onLoanPayment)}, nil
}

func onLoanPayment(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields ──
	loanId := shared.TopicToBigInt(log.Topics[1])
	payer := shared.TopicToAddress(log.Topics[2])

	// ── 2. Decode non-indexed payment data ──
	pay, err := shared.DecodeLoanPaymentData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode PaymentMade data: %w", err)
	}

	logger.Info("PaymentMade",
		"loanId", loanId,
		"payer", payer.Hex(),
		"amount", shared.FormatWei(pay.Amount),
		"totalPaid", shared.FormatWei(pay.AmountPaid),
		"remaining", shared.FormatWei(pay.Remaining),
	)

	// ── 3. Read borrower's current risk score ──
	scoreCall, err := shared.PackGetRiskProfileScore(payer)
	if err != nil {
		return "", fmt.Errorf("pack getRiskProfileScore: %w", err)
	}
	scoreReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DIDRegistryAddr), Data: scoreCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getRiskProfileScore: %w", err)
	}
	currentScore, err := shared.UnpackUint256(scoreReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack risk score: %w", err)
	}

	// ── 4. Check accrued interest to detect full repayment ──
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

	// Loan is fully repaid when remaining == 0 and no accrued interest
	isPayoff := pay.Remaining.Sign() == 0 && interest.Cmp(big.NewInt(0)) == 0

	// ── 5. Compute new risk score ──
	var delta int64
	var reason string
	if isPayoff {
		delta = shared.AdjLoanRepaidFull
		reason = "loan_payoff"
	} else {
		delta = shared.AdjLoanPaymentOnTime
		reason = "loan_payment"
	}
	newScore := shared.AdjustScore(currentScore, delta)
	profileHash := shared.ComputeProfileHash(payer, newScore, reason)

	logger.Info("adjusting risk",
		"payer", payer.Hex(),
		"current", currentScore, "new", newScore,
		"delta", delta, "isPayoff", isPayoff,
	)

	// ── 6. Write updated risk profile on-chain ──
	updateCall, err := shared.PackUpdateRiskProfile(payer, newScore, profileHash)
	if err != nil {
		return "", fmt.Errorf("pack updateRiskProfile: %w", err)
	}
	report, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: updateCall,
	}).Await()
	if err != nil {
		return "", fmt.Errorf("generateReport updateRiskProfile: %w", err)
	}
	_, err = evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: shared.AddrBytes(config.DIDRegistryAddr),
		Report:   report,
	}).Await()
	if err != nil {
		return "", fmt.Errorf("writeReport updateRiskProfile: %w", err)
	}

	// ── 7. Notify backend ──
	body := fmt.Sprintf(
		`{"loanId":"%s","payer":"%s","amountWei":"%s","totalPaid":"%s","remaining":"%s","isPayoff":%t,"newScore":"%s","tier":"%s"}`,
		loanId, payer.Hex(),
		pay.Amount.String(), pay.AmountPaid.String(), pay.Remaining.String(),
		isPayoff, newScore.String(), shared.CreditTier(newScore),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/loans/payment", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed PaymentMade", "loanId", loanId, "isPayoff", isPayoff)
	return "OK", nil
}
