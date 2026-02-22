// Workflow: optimus_dao_vote_cast
//
// Trigger : EVM Log – VoteCast(uint256,address,bool,uint256)
// Purpose : When a vote is cast on a DAO proposal, this workflow:
//   1. Decodes the vote details (support direction, voting weight).
//   2. Notifies the backend of the individual vote.
//   3. Optionally notifies the backend when quorum may have been reached.
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
		Addresses: [][]byte{shared.AddrBytes(config.DAOManagerAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigVoteCast.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onVoteCast)}, nil
}

func onVoteCast(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()

	// ── 1. Extract indexed fields ──
	proposalId := shared.TopicToBigInt(log.Topics[1])
	voter := shared.TopicToAddress(log.Topics[2])

	// ── 2. Decode non-indexed vote data ──
	vote, err := shared.DecodeVoteCastData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode VoteCast data: %w", err)
	}

	direction := "AGAINST"
	if vote.Support {
		direction = "FOR"
	}

	logger.Info("VoteCast",
		"proposalId", proposalId,
		"voter", voter.Hex(),
		"direction", direction,
		"weight", vote.Weight,
	)

	// ── 3. Notify backend of the vote ──
	body := fmt.Sprintf(
		`{"proposalId":"%s","voter":"%s","support":%t,"weight":"%s","direction":"%s"}`,
		proposalId, voter.Hex(), vote.Support, vote.Weight.String(), direction,
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/dao/votes", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend vote notification failed", "err", err)
	}

	// ── 4. Check for quorum by notifying backend with proposal context ──
	// The backend tracks cumulative votes and will respond if quorum is reached.
	quorumBody := fmt.Sprintf(
		`{"proposalId":"%s","latestVoter":"%s","latestWeight":"%s","latestSupport":%t}`,
		proposalId, voter.Hex(), vote.Weight.String(), vote.Support,
	)
	quorumResp, err := shared.ConfidentialRequest(
		shared.BackendReq{
			URL:    config.BackendURL + fmt.Sprintf("/api/dao/proposals/%s/quorum-check", proposalId),
			Method: "POST", Body: quorumBody, Owner: config.Owner,
		},
		runtime,
	)
	if err != nil {
		logger.Warn("quorum check failed", "err", err)
	} else if quorumResp.StatusCode == 200 {
		logger.Info("quorum may have been reached", "proposalId", proposalId)
	}

	logger.Info("processed VoteCast", "proposalId", proposalId, "voter", voter.Hex())
	return "OK", nil
}
