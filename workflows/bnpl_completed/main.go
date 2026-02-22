// Workflow: optimus_bnpl_completed
//
// Trigger : EVM Log – BNPLCompleted(uint256,uint256)
// Purpose : When a BNPL arrangement is marked completed on-chain, this workflow:
//   1. Reads the full arrangement state for final reporting.
//   2. Reads the DAO treasury balance to check for rebalancing needs.
//   3. Notifies the backend of completion.
//   4. Triggers a treasury rebalance check for the DAO.
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
			{Values: [][]byte{shared.SigBNPLCompleted.Bytes()}},
		},
		Confidence: evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
	})

	return cre.Workflow[*shared.Config]{cre.Handler(logTrigger, onBNPLCompleted)}, nil
}

func onBNPLCompleted(config *shared.Config, runtime cre.Runtime, log *evm.Log) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Extract indexed fields ──
	arrangementId := shared.TopicToBigInt(log.Topics[1])

	// ── 2. Decode non-indexed data ──
	completed, err := shared.DecodeBNPLCompletedData(log.Data)
	if err != nil {
		return "", fmt.Errorf("decode BNPLCompleted data: %w", err)
	}

	logger.Info("BNPLCompleted",
		"arrangementId", arrangementId,
		"completedAt", completed.CompletedAt,
	)

	// ── 3. Read full arrangement for reporting ──
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

	// ── 4. Notify backend of completion ──
	body := fmt.Sprintf(
		`{"arrangementId":"%s","daoId":"%s","payer":"%s","totalAmount":"%s","completedAt":"%s","status":"COMPLETED"}`,
		arrangementId, arr.DaoId, arr.Payer.Hex(),
		arr.TotalAmount.String(), completed.CompletedAt.String(),
	)
	_, err = shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/bnpl/completed", Method: "POST", Body: body, Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("backend notification (completed) failed", "err", err)
	}

	// ── 5. Read DAO treasury balance and trigger rebalance check ──
	balCall, err := shared.PackGetTreasuryBalance(arr.DaoId)
	if err != nil {
		return "", fmt.Errorf("pack getTreasuryBalance: %w", err)
	}
	balReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DAOManagerAddr), Data: balCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getTreasuryBalance: %w", err)
	}
	balance, err := shared.UnpackUint256(balReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack treasury balance: %w", err)
	}

	// Check if treasury needs rebalancing – alert backend
	threshold := new(big.Int)
	if config.TreasuryAlertThresholdWei != "" {
		threshold.SetString(config.TreasuryAlertThresholdWei, 10)
	}

	if threshold.Sign() > 0 && balance.Cmp(threshold) < 0 {
		rebalBody := fmt.Sprintf(
			`{"daoId":"%s","currentBalance":"%s","threshold":"%s","trigger":"bnpl_completed","arrangementId":"%s"}`,
			arr.DaoId, balance.String(), threshold.String(), arrangementId,
		)
		_, err = shared.ConfidentialRequest(
			shared.BackendReq{
				URL: config.BackendURL + fmt.Sprintf("/api/dao/%s/treasury-rebalance", arr.DaoId),
				Method: "POST", Body: rebalBody, Owner: config.Owner,
			},
			runtime,
		)
		if err != nil {
			logger.Warn("treasury rebalance notification failed", "err", err)
		}
		logger.Info("treasury below threshold",
			"daoId", arr.DaoId, "balance", shared.FormatWei(balance),
			"threshold", shared.FormatWei(threshold),
		)
	}

	logger.Info("processed BNPLCompleted", "arrangementId", arrangementId, "daoId", arr.DaoId)
	return "OK", nil
}
