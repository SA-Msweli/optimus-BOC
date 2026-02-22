// Workflow: optimus_loan_created
//
// Trigger : EVM Log – LoanCreated(uint256,address,uint256,uint256,uint256,uint256)
// Purpose : When a new loan is created on-chain, this workflow:
//   1. Decodes the loan creation event data (principal, rates, terms).
//   2. Reads the borrower's current DID risk score to assess eligibility.
//   3. Computes max allowed principal for the borrower's tier.
//   4. Flags over-limit loans for review.
//   5. Notifies the backend via confidential HTTP.
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
			{Values: [][]byte{shared.SigLoanCreated.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onLoanCreated)}, nil
}

func onLoanCreated(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields ──
	loanId := shared.TopicToBigInt(log.Topics[1])
	borrower := shared.TopicToAddress(log.Topics[2])

	// ── 2. Decode non-indexed loan creation data ──
	loanData, err := shared.DecodeLoanCreatedData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode LoanCreated data: %w", err)
	}

	logger.Info("LoanCreated",
		"loanId", loanId,
		"borrower", borrower.Hex(),
		"principal", shared.FormatWei(loanData.Principal),
		"interestRateBps", loanData.InterestRateBps,
		"startTime", loanData.StartTime,
		"endTime", loanData.EndTime,
	)

	// ── 3. Read borrower's risk profile score ──
	scoreCall, err := shared.PackGetRiskProfileScore(borrower)
	if err != nil {
		return "", fmt.Errorf("pack getRiskProfileScore: %w", err)
	}
	scoreReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DIDRegistryAddr), Data: scoreCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getRiskProfileScore: %w", err)
	}
	score, err := shared.UnpackUint256(scoreReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack risk score: %w", err)
	}

	tier := shared.CreditTier(score)
	maxPrincipal := shared.MaxLoanPrincipal(score)
	overLimit := loanData.Principal.Cmp(maxPrincipal) > 0

	if overLimit {
		logger.Warn("loan exceeds tier limit",
			"loanId", loanId, "borrower", borrower.Hex(),
			"principal", shared.FormatWei(loanData.Principal),
			"maxAllowed", shared.FormatWei(maxPrincipal),
			"tier", tier, "score", score,
		)
	}

	// ── 4. Notify backend via confidential HTTP ──
	body := fmt.Sprintf(
		`{"loanId":"%s","borrower":"%s","principal":"%s","interestRateBps":"%s","startTime":"%s","endTime":"%s","riskScore":"%s","tier":"%s","maxPrincipal":"%s","overLimit":%t}`,
		loanId, borrower.Hex(),
		loanData.Principal.String(), loanData.InterestRateBps.String(),
		loanData.StartTime.String(), loanData.EndTime.String(),
		score.String(), tier, maxPrincipal.String(), overLimit,
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/loans/created", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed LoanCreated", "loanId", loanId, "tier", tier, "overLimit", overLimit)
	return "OK", nil
}
