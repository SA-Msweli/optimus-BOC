// Workflow: optimus_risk_profile_updater
//
// Trigger : EVM Log – RiskProfileUpdated(address,uint256,bytes32)
// Purpose : When a risk profile is updated on-chain, this workflow:
//   1. Decodes the new score from the event.
//   2. Computes the credit tier (EXCELLENT/GOOD/FAIR/POOR).
//   3. Computes the max BNPL and loan amounts for the new tier.
//   4. Reads the previous tier from the backend to detect tier changes.
//   5. Notifies the backend of the score update and any tier transition.
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
		Addresses: [][]byte{shared.AddrBytes(config.DIDRegistryAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigRiskProfileUpdated.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onRiskProfileUpdated)}, nil
}

func onRiskProfileUpdated(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()

	// ── 1. Extract indexed fields ──
	owner := shared.TopicToAddress(log.Topics[1])
	// Topics[2] is the profileHash (bytes32) – logged but not specifically used

	// ── 2. Decode non-indexed data (new score) ──
	data, err := shared.DecodeRiskProfileUpdatedData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode RiskProfileUpdated data: %w", err)
	}

	newScore := data.NewScore
	newTier := shared.CreditTier(newScore)
	maxBNPL := shared.MaxBNPLAmount(newScore)
	maxLoan := shared.MaxLoanPrincipal(newScore)

	logger.Info("RiskProfileUpdated",
		"owner", owner.Hex(),
		"newScore", newScore,
		"tier", newTier,
		"maxBNPL", shared.FormatWei(maxBNPL),
		"maxLoan", shared.FormatWei(maxLoan),
	)

	// ── 3. Notify backend of the score update ──
	body := fmt.Sprintf(
		`{"owner":"%s","newScore":"%s","tier":"%s","maxBnplWei":"%s","maxLoanWei":"%s"}`,
		owner.Hex(), newScore.String(), newTier,
		maxBNPL.String(), maxLoan.String(),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/did/risk-updated", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend risk-updated notification failed", "err", err)
	}

	// ── 4. Ask backend for the previous tier to detect tier changes ──
	prevResp, err := shared.ConfidentialRequest(
		shared.BackendReq{
			URL:    config.BackendURL + fmt.Sprintf("/api/did/%s/previous-tier", owner.Hex()),
			Method: "GET", Owner: config.Owner,
		},
		runtime,
	)
	if err != nil {
		logger.Warn("failed to fetch previous tier", "err", err)
	} else {
		previousTier := prevResp.Body
		if previousTier != "" && previousTier != newTier {
			tierBody := fmt.Sprintf(
				`{"owner":"%s","previousTier":"%s","newTier":"%s","newScore":"%s","maxBnplWei":"%s","maxLoanWei":"%s"}`,
				owner.Hex(), previousTier, newTier, newScore.String(),
				maxBNPL.String(), maxLoan.String(),
			)
			_, err = shared.ConfidentialRequest(
				shared.BackendReq{
					URL: config.BackendURL + "/api/did/tier-changed", Method: "POST",
					Body: tierBody, Owner: config.Owner,
				},
				runtime,
			)
			if err != nil {
				logger.Warn("tier-changed notification failed", "err", err)
			}
			logger.Info("tier changed",
				"owner", owner.Hex(),
				"from", previousTier, "to", newTier,
			)
		}
	}

	logger.Info("processed RiskProfileUpdated", "owner", owner.Hex(), "score", newScore, "tier", newTier)
	return "OK", nil
}
