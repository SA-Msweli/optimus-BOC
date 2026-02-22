// Workflow: optimus_treasury_monitor
//
// Trigger : Cron – every 6 hours
// Purpose : Monitors the TokenVault ETH balance and each DAO treasury balance.
//   Alerts the backend when the vault balance drops below a configured threshold.
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
			onCronTreasuryCheck,
		),
	}
	return workflow, nil
}

func onCronTreasuryCheck(config *shared.Config, runtime cre.Runtime, _ *cron.Payload) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}

	// ── 1. Check TokenVault ETH balance against threshold ──
	token := common.HexToAddress(config.ETHTokenAddress)
	balCall, err := shared.PackGetBalance(token)
	if err != nil {
		return "", fmt.Errorf("pack getBalance: %w", err)
	}
	balReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.TokenVaultAddr), Data: balCall},
	}).Await()
	if err != nil {
		return "", fmt.Errorf("callContract getBalance: %w", err)
	}
	balance, err := shared.UnpackUint256(balReply.Data)
	if err != nil {
		return "", fmt.Errorf("unpack balance: %w", err)
	}

	threshold, ok := new(big.Int).SetString(config.TreasuryAlertThresholdWei, 10)
	if !ok {
		threshold = big.NewInt(1_000_000_000_000_000_000) // 1 ETH fallback
	}

	if balance.Cmp(threshold) < 0 {
		alertBody := fmt.Sprintf(`{"vaultBalance":"%s","threshold":"%s"}`, balance, threshold)
		_, _ = shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + "/api/treasury/alert", Method: "POST", Body: alertBody, Owner: config.Owner},
			runtime,
		)
		logger.Warn("vault balance below threshold", "balance", shared.FormatWei(balance), "threshold", shared.FormatWei(threshold))
	}

	// ── 2. Check each DAO treasury ──
	resp, err := shared.ConfidentialRequest(
		shared.BackendReq{URL: config.BackendURL + "/api/dao/active-ids", Method: "GET", Owner: config.Owner},
		runtime,
	)
	if err != nil {
		logger.Warn("fetch DAO active-ids failed", "err", err)
		return fmt.Sprintf("vault_balance=%s", balance), nil
	}

	var daoIDs []int64
	_ = json.Unmarshal([]byte(resp.Body), &daoIDs)

	for _, id := range daoIDs {
		daoID := big.NewInt(id)
		tCall, _ := shared.PackGetTreasuryBalance(daoID)
		tReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
			Call: &evm.CallMsg{To: shared.AddrBytes(config.DAOManagerAddr), Data: tCall},
		}).Await()
		if err != nil {
			continue
		}
		daoBalance, err := shared.UnpackUint256(tReply.Data)
		if err != nil {
			continue
		}

		statusBody := fmt.Sprintf(`{"daoId":"%s","treasuryBalance":"%s"}`, daoID, daoBalance)
		_, _ = shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + fmt.Sprintf("/api/dao/%d/treasury-status", id), Method: "POST", Body: statusBody, Owner: config.Owner},
			runtime,
		)
	}

	logger.Info("treasury monitor complete", "vault_balance", shared.FormatWei(balance), "daos_checked", len(daoIDs))
	return fmt.Sprintf("vault_balance=%s", balance), nil
}
