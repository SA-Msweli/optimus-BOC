// Workflow: optimus_bnpl_late_fee
//
// Trigger : Cron – every 30 minutes
// Purpose : Scans active BNPL arrangements for overdue installments past the
//   grace period.  For each overdue installment it:
//   1. Calls BNPLManager.applyLateFee on-chain.
//   2. Adjusts the payer's DID risk score downward (-300).
//   3. Notifies the backend via confidential HTTP.
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"time"

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
			onCronLateFee,
		),
	}
	return workflow, nil
}

func onCronLateFee(config *shared.Config, runtime cre.Runtime, _ *cron.Payload) (string, error) {
	logger := runtime.Logger()
	evmClient := &evm.Client{ChainSelector: evm.EthereumTestnetSepolia}
	now := time.Now().Unix()

	// ── 1. Resolve active arrangement IDs ──
	arrIDs := config.ActiveArrangementIDs
	if len(arrIDs) == 0 {
		resp, err := shared.ConfidentialRequest(
			shared.BackendReq{URL: config.BackendURL + "/api/bnpl/active-ids", Method: "GET", Owner: config.Owner},
			runtime,
		)
		if err != nil {
			return "", fmt.Errorf("fetch active arrangement IDs: %w", err)
		}
		arrIDs = parseIDList([]byte(resp.Body))
	}

	applied := 0
	for _, id := range arrIDs {
		arrID := big.NewInt(id)

		// ── 2. Read arrangement state ──
		arrCall, err := shared.PackGetArrangement(arrID)
		if err != nil {
			continue
		}
		arrReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
			Call: &evm.CallMsg{To: shared.AddrBytes(config.BNPLManagerAddr), Data: arrCall},
		}).Await()
		if err != nil {
			logger.Warn("getArrangement failed", "id", id, "err", err)
			continue
		}
		arr, err := shared.UnpackArrangement(arrReply.Data)
		if err != nil || arr.Status != shared.BNPLStatusActive {
			continue
		}

		// ── 3. Read DAO BNPL terms for grace period ──
		termsCall, _ := shared.PackGetBnplTerms(arr.DaoId)
		termsReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
			Call: &evm.CallMsg{To: shared.AddrBytes(config.DAOManagerAddr), Data: termsCall},
		}).Await()
		if err != nil {
			continue
		}
		terms, err := shared.UnpackBnplTerms(termsReply.Data)
		if err != nil {
			continue
		}

		graceSec := config.GracePeriodSeconds
		if terms.GracePeriodDays.Sign() > 0 {
			graceSec = terms.GracePeriodDays.Int64() * 86400
		}

		// ── 4. Check each installment for overdue status ──
		numInst := int(arr.NumInstallments.Int64())
		for i := 0; i < numInst; i++ {
			dueTime := new(big.Int).Add(
				arr.StartTimestamp,
				new(big.Int).Mul(big.NewInt(int64(i)+1), arr.IntervalSeconds),
			).Int64()

			if now <= dueTime+graceSec {
				continue // not overdue yet
			}

			// ── 5. Apply late fee on-chain ──
			feeCall, err := shared.PackApplyLateFee(arrID, uint8(i))
			if err != nil {
				continue
			}
			feeReport, err := runtime.GenerateReport(&cre.ReportRequest{
				EncodedPayload: feeCall,
			}).Await()
			if err != nil {
				logger.Warn("generateReport applyLateFee failed", "arr", id, "inst", i, "err", err)
				continue
			}
			writeReply, err := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
				Receiver: shared.AddrBytes(config.BNPLManagerAddr),
				Report:   feeReport,
			}).Await()
			if err != nil || writeReply.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
				logger.Warn("applyLateFee tx failed", "arr", id, "inst", i, "err", err)
				continue
			}

			// ── 6. Adjust risk score ──
			adjustRisk(evmClient, runtime, config, arr.Payer, shared.AdjBNPLLateFee, "bnpl_late_fee", logger)
			applied++
		}
	}

	logger.Info("late fee scan complete", "applied", applied)
	return fmt.Sprintf("applied=%d", applied), nil
}

// adjustRisk reads the payer's current risk score, applies a delta, and writes
// the updated profile back on-chain.
func adjustRisk(evmClient *evm.Client, runtime cre.Runtime, config *shared.Config, owner common.Address, delta int64, reason string, logger *slog.Logger) {
	scoreCall, err := shared.PackGetRiskProfileScore(owner)
	if err != nil {
		return
	}
	scoreReply, err := evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call: &evm.CallMsg{To: shared.AddrBytes(config.DIDRegistryAddr), Data: scoreCall},
	}).Await()
	if err != nil {
		return
	}
	currentScore, err := shared.UnpackUint256(scoreReply.Data)
	if err != nil {
		return
	}

	newScore := shared.AdjustScore(currentScore, delta)
	profileHash := shared.ComputeProfileHash(owner, newScore, reason)
	updateCall, err := shared.PackUpdateRiskProfile(owner, newScore, profileHash)
	if err != nil {
		return
	}

	riskReport, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: updateCall,
	}).Await()
	if err != nil {
		logger.Warn("generateReport updateRiskProfile failed", "owner", owner.Hex(), "err", err)
		return
	}

	_, err = evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver: shared.AddrBytes(config.DIDRegistryAddr),
		Report:   riskReport,
	}).Await()
	if err != nil {
		logger.Warn("updateRiskProfile failed", "owner", owner.Hex(), "err", err)
	}
}

func parseIDList(body []byte) []int64 {
	var ids []int64
	_ = json.Unmarshal(body, &ids)
	return ids
}
