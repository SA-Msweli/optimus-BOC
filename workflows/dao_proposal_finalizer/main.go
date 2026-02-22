// Workflow: optimus_dao_proposal_finalizer
//
// Trigger : Cron – every hour
// Purpose : Scans open DAO proposals.  Finalizes any that have passed their
//   expiry timestamp by calling DAOManager.finalizeProposal on-chain.
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"time"

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
			onCronFinalizeProposals,
		),
	}
	return workflow, nil
}

func onCronFinalizeProposals(config *shared.Config, runtime cre.Runtime, _ *cron.Payload) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}
	now := time.Now().Unix()

	// ── 1. Get open proposal IDs from config or backend ──
	proposalIDs := config.ActiveProposalIDs
	if len(proposalIDs) == 0 {
		resp, err := shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + "/api/dao/proposals/open-ids", Method: "GET", Owner: config.Owner},
			runtime,
		)
		if err != nil {
			return "", fmt.Errorf("fetch open proposal IDs: %w", err)
		}
		proposalIDs = parseIDList([]byte(resp.Body))
	}

	finalized := 0
	for _, id := range proposalIDs {
		proposalID := big.NewInt(id)

		// ── 2. Fetch proposal details to get expiry timestamp ──
		resp, err := shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + fmt.Sprintf("/api/dao/proposals/%d", id), Method: "GET", Owner: config.Owner},
			runtime,
		)
		if err != nil {
			logger.Warn("fetch proposal details failed", "id", id, "err", err)
			continue
		}

		expiry := extractExpiry([]byte(resp.Body))
		if expiry == 0 || now <= expiry {
			continue // not expired yet
		}

		// ── 3. Finalize proposal on-chain ──
		finalizeCall, err := shared.PackFinalizeProposal(proposalID)
		if err != nil {
			continue
		}
		finalizeReport, err := runtime.GenerateReport(&cre.ReportRequest{
			EncodedPayload: finalizeCall,
		}).Await()
		if err != nil {
			logger.Warn("generateReport finalizeProposal failed", "id", id, "err", err)
			continue
		}
		writeReply, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver: shared.AddrBytes(config.DAOManagerAddr),
			Report:   finalizeReport,
		}).Await()
		if err != nil || writeReply.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
			logger.Warn("finalizeProposal tx failed", "id", id, "err", err)
			continue
		}

		// ── 4. Notify backend ──
		body := fmt.Sprintf(`{"proposalId":"%s","status":"FINALIZED"}`, proposalID)
		_, _ = shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + fmt.Sprintf("/api/dao/proposals/%d/finalized", id), Method: "POST", Body: body, Owner: config.Owner},
			runtime,
		)

		finalized++
		logger.Info("proposal finalized", "id", id)
	}

	logger.Info("proposal finalizer complete", "finalized", finalized)
	return fmt.Sprintf("finalized=%d", finalized), nil
}

func extractExpiry(body []byte) int64 {
	var data struct {
		Expiry int64 `json:"expiry"`
	}
	_ = json.Unmarshal(body, &data)
	return data.Expiry
}

func parseIDList(body []byte) []int64 {
	var ids []int64
	_ = json.Unmarshal(body, &ids)
	return ids
}
