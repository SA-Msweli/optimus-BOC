// Workflow: optimus_bnpl_created
//
// Trigger : EVM Log – BNPLCreated(uint256,uint256,address,address,uint256,uint256,uint256)
// Purpose : When a new BNPL arrangement is created on-chain, this workflow:
//   1. Reads the full arrangement state from BNPLManager.
//   2. Reads the DAO's BNPL terms to validate the arrangement parameters.
//   3. Notifies the backend via confidential HTTP.
//   4. Reports any policy violations for human review.
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
		Addresses: [][]byte{shared.AddrBytes(config.BNPLManagerAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigBNPLCreated.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	workflow := cre.Workflow[*shared.Config]{
		cre.Handler(logTrigger, onBNPLCreated),
	}
	return workflow, nil
}

func onBNPLCreated(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields from topics ──
	arrangementId := shared.TopicToBigInt(log.Topics[1])
	daoId := shared.TopicToBigInt(log.Topics[2])
	payer := shared.TopicToAddress(log.Topics[3])

	// ── 2. Decode non-indexed event data ──
	created, err := shared.DecodeBNPLCreatedData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode BNPLCreated data: %w", err)
	}

	logger.Info("BNPLCreated",
		"arrangementId", arrangementId,
		"daoId", daoId,
		"payer", payer.Hex(),
		"recipient", created.Recipient.Hex(),
		"total", shared.FormatWei(created.TotalAmount),
		"installments", created.NumInstallments,
	)

	// ── 3. Read full arrangement state ──
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

	// ── 4. Read DAO BNPL terms for policy validation ──
	termsCall, err := shared.PackGetBnplTerms(daoId)
	if err != nil {
		return "", fmt.Errorf("pack getBnplTerms: %w", err)
	}
	termsReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DAOManagerAddr), Data: termsCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getBnplTerms: %w", err)
	}
	terms, err := shared.UnpackBnplTerms(termsReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack bnplTerms: %w", err)
	}

	// ── 5. Validate arrangement against DAO terms ──
	violations := validateArrangement(arr, terms)
	if len(violations) > 0 {
		logger.Warn("policy violations", "arrangementId", arrangementId, "violations", violations)
	}

	// ── 6. Compute installment schedule ──
	numInst := int(arr.NumInstallments.Int64())
	type installment struct {
		Number    int    `json:"number"`
		DueTime   int64  `json:"dueTime"`
		AmountWei string `json:"amountWei"`
	}
	schedule := make([]installment, numInst)
	for i := 0; i < numInst; i++ {
		dueTime := new(big.Int).Add(
			arr.StartTimestamp,
			new(big.Int).Mul(big.NewInt(int64(i+1)), arr.IntervalSeconds),
		).Int64()
		amount := big.NewInt(0)
		if i < len(arr.InstallmentAmounts) {
			amount = arr.InstallmentAmounts[i]
		}
		schedule[i] = installment{Number: i, DueTime: dueTime, AmountWei: amount.String()}
	}

	// ── 7. Notify backend via confidential HTTP ──
	body := fmt.Sprintf(
		`{"arrangementId":"%s","daoId":"%s","payer":"%s","recipient":"%s","totalAmount":"%s","numInstallments":%d,"status":"PENDING","violations":%d}`,
		arrangementId, daoId, payer.Hex(), created.Recipient.Hex(),
		created.TotalAmount.String(), numInst, len(violations),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/bnpl/created", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed BNPLCreated", "arrangementId", arrangementId, "installments", numInst)
	return "OK", nil
}

func validateArrangement(arr *shared.Arrangement, terms *shared.BnplTerms) []string {
	var violations []string

	if terms.NumInstallments.Sign() > 0 {
		if arr.NumInstallments.Cmp(terms.NumInstallments) != 0 {
			violations = append(violations, fmt.Sprintf(
				"installment count %s does not match DAO policy %s",
				arr.NumInstallments, terms.NumInstallments))
		}
	}

	intervalDays := new(big.Int).Div(arr.IntervalSeconds, big.NewInt(86400))
	if terms.AllowedIntervalMinDays.Sign() > 0 && intervalDays.Cmp(terms.AllowedIntervalMinDays) < 0 {
		violations = append(violations, fmt.Sprintf(
			"interval %s days below minimum %s days",
			intervalDays, terms.AllowedIntervalMinDays))
	}
	if terms.AllowedIntervalMaxDays.Sign() > 0 && intervalDays.Cmp(terms.AllowedIntervalMaxDays) > 0 {
		violations = append(violations, fmt.Sprintf(
			"interval %s days exceeds maximum %s days",
			intervalDays, terms.AllowedIntervalMaxDays))
	}

	return violations
}
