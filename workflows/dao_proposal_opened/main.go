// Workflow: optimus_dao_proposal_opened
//
// Trigger : EVM Log – ProposalOpened(uint256,uint256,uint256,bytes)
// Purpose : When a new DAO proposal is opened on-chain, this workflow:
//   1. Decodes the proposal event with expiry and attached data.
//   2. Attempts to parse the data payload as a treasury operation.
//   3. Notifies the backend via confidential HTTP with the full proposal details.
package main

import (
	"encoding/hex"
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
		Addresses: [][]byte{shared.AddrBytes(config.DAOManagerAddr)},
		Topics: []*evm.TopicValues{
			{Values: [][]byte{shared.SigProposalOpened.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onProposalOpened)}, nil
}

func onProposalOpened(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()

	// ── 1. Extract indexed fields ──
	proposalId := shared.TopicToBigInt(log.Topics[1])
	daoId := shared.TopicToBigInt(log.Topics[2])

	// ── 2. Decode non-indexed proposal data ──
	proposal, err := shared.DecodeProposalOpenedData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode ProposalOpened data: %w", err)
	}

	logger.Info("ProposalOpened",
		"proposalId", proposalId,
		"daoId", daoId,
		"expiry", proposal.Expiry,
		"dataLen", len(proposal.Data),
	)

	// ── 3. Classify the proposal based on data payload ──
	proposalType := classifyProposal(proposal.Data)

	// ── 4. Check if this is a treasury operation and extract details ──
	var treasuryAmountStr string
	if proposalType == "TREASURY_TRANSFER" && len(proposal.Data) >= 32 {
		amount := new(big.Int).SetBytes(proposal.Data[:32])
		treasuryAmountStr = amount.String()
	}

	// ── 5. Notify backend via confidential HTTP ──
	body := fmt.Sprintf(
		`{"proposalId":"%s","daoId":"%s","expiry":"%s","proposalType":"%s","dataHex":"%s","treasuryAmount":"%s"}`,
		proposalId, daoId, proposal.Expiry.String(),
		proposalType, hex.EncodeToString(proposal.Data), treasuryAmountStr,
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/dao/proposals/opened", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification failed", "err", err)
	}

	logger.Info("processed ProposalOpened", "proposalId", proposalId, "type", proposalType)
	return "OK", nil
}

// classifyProposal inspects the data payload to determine the proposal type.
// Solidity function selectors (first 4 bytes) are used to identify treasury operations.
func classifyProposal(data []byte) string {
	if len(data) < 4 {
		return "GENERIC"
	}

	// Known treasury operation selectors:
	sel := hex.EncodeToString(data[:4])
	switch sel {
	case "a9059cbb": // transfer(address,uint256)
		return "TREASURY_TRANSFER"
	case "095ea7b3": // approve(address,uint256)
		return "TREASURY_APPROVE"
	case "23b872dd": // transferFrom(address,address,uint256)
		return "TREASURY_TRANSFER_FROM"
	default:
		return "CUSTOM"
	}
}
