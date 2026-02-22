// Workflow: optimus_bnpl_payment
//
// Trigger : EVM Log – BNPLPaymentMade(uint256,uint8,address,uint256,uint256)
// Purpose : When a BNPL installment payment is made on-chain, this workflow:
//   1. Decodes the payment details from the event.
//   2. Reads the full arrangement to determine if this is the final payment.
//   3. Reads the payer's current DID risk score.
//   4. Adjusts the risk score upward (+50 per payment, +500 if final).
//   5. Writes the updated risk profile on-chain.
//   6. Notifies the backend via confidential HTTP.
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
		Addresses: [][]byte{shared.AddrBytes(config.BNPLManagerAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigBNPLPaymentMade.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onBNPLPayment)}, nil
}

func onBNPLPayment(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields ──
	arrangementId := shared.TopicToBigInt(log.Topics[1])
	payer := shared.TopicToAddress(log.Topics[2])

	// ── 2. Decode non-indexed payment data ──
	pay, err := shared.DecodeBNPLPaymentData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode BNPLPaymentMade data: %w", err)
	}
	logger.Info("BNPLPaymentMade",
		"arrangementId", arrangementId,
		"payer", payer.Hex(),
		"installment", pay.InstallmentNumber,
		"amount", shared.FormatWei(pay.Amount),
	)

	// ── 3. Read full arrangement to check if this is the final payment ──
	arrCall, err := shared.PackGetArrangement(arrangementId)
	if err != nil {
		return "", fmt.Errorf("pack getArrangement: %w", err)
	}
	arrReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.BNPLManagerAddr), Data: arrCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getArrangement: %w", err)
	}
	arr, err := shared.UnpackArrangement(arrReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack arrangement: %w", err)
	}

	isFinal := int(pay.InstallmentNumber) >= int(arr.NumInstallments.Int64())-1

	// ── 4. Read current risk score ──
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

	// ── 5. Compute new risk score ──
	var delta int64
	var reason string
	if isFinal {
		delta = shared.AdjBNPLCompletedOnTime
		reason = "bnpl_final_payment"
	} else {
		delta = shared.AdjBNPLPaymentOnTime
		reason = "bnpl_payment"
	}
	newScore := shared.AdjustScore(currentScore, delta)
	profileHash := shared.ComputeProfileHash(payer, newScore, reason)

	logger.Info("adjusting risk",
		"payer", payer.Hex(),
		"current", currentScore, "new", newScore,
		"delta", delta, "reason", reason,
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

	// ── 7. Notify backend via confidential HTTP ──
	body := fmt.Sprintf(
		`{"arrangementId":"%s","payer":"%s","installmentNumber":%d,"amountWei":"%s","isFinal":%t,"newScore":"%s","tier":"%s"}`,
		arrangementId, payer.Hex(), pay.InstallmentNumber,
		pay.Amount.String(), isFinal, newScore.String(), shared.CreditTier(newScore),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/bnpl/payment", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed BNPLPaymentMade", "arrangementId", arrangementId, "isFinal", isFinal)
	return "OK", nil
}
