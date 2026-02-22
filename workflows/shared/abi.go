// Package shared provides ABI encoding/decoding helpers and common types
// used across all Optimus CRE workflows.  The helpers wrap go-ethereum's
// accounts/abi package so that each workflow can call Pack*/Unpack* without
// duplicating low-level ABI logic.
package shared

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ──────────────────────────────────────────────
// Reusable ABI types
// ──────────────────────────────────────────────

var (
	Uint256Ty, _    = abi.NewType("uint256", "", nil)
	Uint8Ty, _      = abi.NewType("uint8", "", nil)
	AddressTy, _    = abi.NewType("address", "", nil)
	BoolTy, _       = abi.NewType("bool", "", nil)
	Bytes32Ty, _    = abi.NewType("bytes32", "", nil)
	BytesTy, _      = abi.NewType("bytes", "", nil)
	Uint256ArrTy, _ = abi.NewType("uint256[]", "", nil)
	Uint64Ty, _     = abi.NewType("uint64", "", nil)
)

// selector returns the first 4 bytes of keccak256(sig).
func selector(sig string) []byte {
	return crypto.Keccak256([]byte(sig))[:4]
}

// ──────────────────────────────────────────────
// Domain structs – mirror Solidity return types
// ──────────────────────────────────────────────

// Arrangement mirrors BNPLManager.getArrangement return values.
type Arrangement struct {
	Id                 *big.Int
	DaoId              *big.Int
	Payer              common.Address
	Recipient          common.Address
	TotalAmount        *big.Int
	NumInstallments    *big.Int
	InstallmentAmounts []*big.Int
	StartTimestamp     *big.Int
	IntervalSeconds    *big.Int
	LateFeeBps         *big.Int
	Status             uint8 // 0=PENDING, 1=ACTIVE, 2=COMPLETED
}

// Loan mirrors LoanManager.getLoan return values.
type Loan struct {
	Id              *big.Int
	Borrower        common.Address
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	AmountPaid      *big.Int
	Status          uint8 // 0=PENDING, 1=ACTIVE, 2=REPAID, 3=DEFAULTED
}

// BnplTerms mirrors DAOManager.getBnplTerms return values.
type BnplTerms struct {
	NumInstallments       *big.Int
	AllowedIntervalMinDays *big.Int
	AllowedIntervalMaxDays *big.Int
	LateFeeBps            *big.Int
	GracePeriodDays       *big.Int
	RescheduleAllowed     bool
	MinDownPaymentBps     *big.Int
}

// ──────────────────────────────────────────────
// Status constants
// ──────────────────────────────────────────────

const (
	BNPLStatusPending   uint8 = 0
	BNPLStatusActive    uint8 = 1
	BNPLStatusCompleted uint8 = 2

	LoanStatusPending   uint8 = 0
	LoanStatusActive    uint8 = 1
	LoanStatusRepaid    uint8 = 2
	LoanStatusDefaulted uint8 = 3
)

// ══════════════════════════════════════════════
// BNPLManager helpers
// ══════════════════════════════════════════════

var getArrangementSel = selector("getArrangement(uint256)")

var getArrangementInputs = abi.Arguments{
	{Name: "arrangementId", Type: Uint256Ty},
}

var getArrangementOutputs = abi.Arguments{
	{Name: "id", Type: Uint256Ty},
	{Name: "daoId", Type: Uint256Ty},
	{Name: "payer", Type: AddressTy},
	{Name: "recipient", Type: AddressTy},
	{Name: "totalAmount", Type: Uint256Ty},
	{Name: "numInstallments", Type: Uint256Ty},
	{Name: "installmentAmounts", Type: Uint256ArrTy},
	{Name: "startTimestamp", Type: Uint256Ty},
	{Name: "intervalSeconds", Type: Uint256Ty},
	{Name: "lateFeeBps", Type: Uint256Ty},
	{Name: "status", Type: Uint8Ty},
}

// PackGetArrangement encodes calldata for BNPLManager.getArrangement(uint256).
func PackGetArrangement(arrangementId *big.Int) ([]byte, error) {
	data, err := getArrangementInputs.Pack(arrangementId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getArrangementSel...), data...), nil
}

// UnpackArrangement decodes the ABI-encoded return data from getArrangement.
func UnpackArrangement(data []byte) (*Arrangement, error) {
	vals, err := getArrangementOutputs.Unpack(data)
	if err != nil {
		return nil, err
	}
	amounts := vals[6].([]*big.Int)
	cp := make([]*big.Int, len(amounts))
	copy(cp, amounts)
	return &Arrangement{
		Id:                 vals[0].(*big.Int),
		DaoId:              vals[1].(*big.Int),
		Payer:              vals[2].(common.Address),
		Recipient:          vals[3].(common.Address),
		TotalAmount:        vals[4].(*big.Int),
		NumInstallments:    vals[5].(*big.Int),
		InstallmentAmounts: cp,
		StartTimestamp:     vals[7].(*big.Int),
		IntervalSeconds:    vals[8].(*big.Int),
		LateFeeBps:         vals[9].(*big.Int),
		Status:             vals[10].(uint8),
	}, nil
}

var applyLateFeeSel = selector("applyLateFee(uint256,uint8)")

var applyLateFeeInputs = abi.Arguments{
	{Name: "arrangementId", Type: Uint256Ty},
	{Name: "installmentNumber", Type: Uint8Ty},
}

// PackApplyLateFee encodes calldata for BNPLManager.applyLateFee.
func PackApplyLateFee(arrangementId *big.Int, installmentNumber uint8) ([]byte, error) {
	data, err := applyLateFeeInputs.Pack(arrangementId, installmentNumber)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, applyLateFeeSel...), data...), nil
}

// ══════════════════════════════════════════════
// DAOManager helpers
// ══════════════════════════════════════════════

var getBnplTermsSel = selector("getBnplTerms(uint256)")

var getBnplTermsInputs = abi.Arguments{
	{Name: "daoId", Type: Uint256Ty},
}

var getBnplTermsOutputs = abi.Arguments{
	{Name: "numInstallments", Type: Uint256Ty},
	{Name: "allowedIntervalMinDays", Type: Uint256Ty},
	{Name: "allowedIntervalMaxDays", Type: Uint256Ty},
	{Name: "lateFeeBps", Type: Uint256Ty},
	{Name: "gracePeriodDays", Type: Uint256Ty},
	{Name: "rescheduleAllowed", Type: BoolTy},
	{Name: "minDownPaymentBps", Type: Uint256Ty},
}

// PackGetBnplTerms encodes calldata for DAOManager.getBnplTerms.
func PackGetBnplTerms(daoId *big.Int) ([]byte, error) {
	data, err := getBnplTermsInputs.Pack(daoId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getBnplTermsSel...), data...), nil
}

// UnpackBnplTerms decodes the return data from getBnplTerms.
func UnpackBnplTerms(data []byte) (*BnplTerms, error) {
	vals, err := getBnplTermsOutputs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &BnplTerms{
		NumInstallments:       vals[0].(*big.Int),
		AllowedIntervalMinDays: vals[1].(*big.Int),
		AllowedIntervalMaxDays: vals[2].(*big.Int),
		LateFeeBps:            vals[3].(*big.Int),
		GracePeriodDays:       vals[4].(*big.Int),
		RescheduleAllowed:     vals[5].(bool),
		MinDownPaymentBps:     vals[6].(*big.Int),
	}, nil
}

var getTreasuryBalanceSel = selector("getTreasuryBalance(uint256)")

var getTreasuryBalanceInputs = abi.Arguments{
	{Name: "daoId", Type: Uint256Ty},
}

// PackGetTreasuryBalance encodes calldata for DAOManager.getTreasuryBalance.
func PackGetTreasuryBalance(daoId *big.Int) ([]byte, error) {
	data, err := getTreasuryBalanceInputs.Pack(daoId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getTreasuryBalanceSel...), data...), nil
}

var finalizeProposalSel = selector("finalizeProposal(uint256)")

var finalizeProposalInputs = abi.Arguments{
	{Name: "proposalId", Type: Uint256Ty},
}

// PackFinalizeProposal encodes calldata for DAOManager.finalizeProposal.
func PackFinalizeProposal(proposalId *big.Int) ([]byte, error) {
	data, err := finalizeProposalInputs.Pack(proposalId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, finalizeProposalSel...), data...), nil
}

// ══════════════════════════════════════════════
// LoanManager helpers
// ══════════════════════════════════════════════

var getLoanSel = selector("getLoan(uint256)")

var getLoanInputs = abi.Arguments{
	{Name: "loanId", Type: Uint256Ty},
}

var getLoanOutputs = abi.Arguments{
	{Name: "id", Type: Uint256Ty},
	{Name: "borrower", Type: AddressTy},
	{Name: "principal", Type: Uint256Ty},
	{Name: "interestRateBps", Type: Uint256Ty},
	{Name: "startTime", Type: Uint256Ty},
	{Name: "endTime", Type: Uint256Ty},
	{Name: "amountPaid", Type: Uint256Ty},
	{Name: "status", Type: Uint8Ty},
}

// PackGetLoan encodes calldata for LoanManager.getLoan.
func PackGetLoan(loanId *big.Int) ([]byte, error) {
	data, err := getLoanInputs.Pack(loanId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getLoanSel...), data...), nil
}

// UnpackLoan decodes the return data from getLoan.
func UnpackLoan(data []byte) (*Loan, error) {
	vals, err := getLoanOutputs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &Loan{
		Id:              vals[0].(*big.Int),
		Borrower:        vals[1].(common.Address),
		Principal:       vals[2].(*big.Int),
		InterestRateBps: vals[3].(*big.Int),
		StartTime:       vals[4].(*big.Int),
		EndTime:         vals[5].(*big.Int),
		AmountPaid:      vals[6].(*big.Int),
		Status:          vals[7].(uint8),
	}, nil
}

var getAmountOwedSel = selector("getAmountOwed(uint256)")

var getAmountOwedInputs = abi.Arguments{
	{Name: "loanId", Type: Uint256Ty},
}

// PackGetAmountOwed encodes calldata for LoanManager.getAmountOwed.
func PackGetAmountOwed(loanId *big.Int) ([]byte, error) {
	data, err := getAmountOwedInputs.Pack(loanId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getAmountOwedSel...), data...), nil
}

var getAccruedInterestSel = selector("getAccruedInterest(uint256)")

var getAccruedInterestInputs = abi.Arguments{
	{Name: "loanId", Type: Uint256Ty},
}

// PackGetAccruedInterest encodes calldata for LoanManager.getAccruedInterest.
func PackGetAccruedInterest(loanId *big.Int) ([]byte, error) {
	data, err := getAccruedInterestInputs.Pack(loanId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getAccruedInterestSel...), data...), nil
}

var markDefaultedSel = selector("markDefaulted(uint256)")

var markDefaultedInputs = abi.Arguments{
	{Name: "loanId", Type: Uint256Ty},
}

// PackMarkDefaulted encodes calldata for LoanManager.markDefaulted.
func PackMarkDefaulted(loanId *big.Int) ([]byte, error) {
	data, err := markDefaultedInputs.Pack(loanId)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, markDefaultedSel...), data...), nil
}

// ══════════════════════════════════════════════
// DIDRegistry helpers
// ══════════════════════════════════════════════

var getRiskProfileScoreSel = selector("getRiskProfileScore(address)")

var getRiskProfileScoreInputs = abi.Arguments{
	{Name: "owner", Type: AddressTy},
}

// PackGetRiskProfileScore encodes calldata for DIDRegistry.getRiskProfileScore.
func PackGetRiskProfileScore(owner common.Address) ([]byte, error) {
	data, err := getRiskProfileScoreInputs.Pack(owner)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getRiskProfileScoreSel...), data...), nil
}

var updateRiskProfileSel = selector("updateRiskProfile(address,uint256,bytes32)")

var updateRiskProfileInputs = abi.Arguments{
	{Name: "owner", Type: AddressTy},
	{Name: "newScore", Type: Uint256Ty},
	{Name: "riskProfileHash", Type: Bytes32Ty},
}

// PackUpdateRiskProfile encodes calldata for DIDRegistry.updateRiskProfile.
func PackUpdateRiskProfile(owner common.Address, newScore *big.Int, profileHash [32]byte) ([]byte, error) {
	data, err := updateRiskProfileInputs.Pack(owner, newScore, profileHash)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, updateRiskProfileSel...), data...), nil
}

var didExistsSel = selector("exists(address)")

var didExistsInputs = abi.Arguments{
	{Name: "owner", Type: AddressTy},
}

// PackDIDExists encodes calldata for DIDRegistry.exists.
func PackDIDExists(owner common.Address) ([]byte, error) {
	data, err := didExistsInputs.Pack(owner)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, didExistsSel...), data...), nil
}

// ══════════════════════════════════════════════
// TokenVault helpers
// ══════════════════════════════════════════════

var getBalanceSel = selector("getBalance(address)")

var getBalanceInputs = abi.Arguments{
	{Name: "token", Type: AddressTy},
}

// PackGetBalance encodes calldata for TokenVault.getBalance.
func PackGetBalance(token common.Address) ([]byte, error) {
	data, err := getBalanceInputs.Pack(token)
	if err != nil {
		return nil, err
	}
	return append(append([]byte{}, getBalanceSel...), data...), nil
}

// ══════════════════════════════════════════════
// Generic unpackers
// ══════════════════════════════════════════════

var uint256Output = abi.Arguments{{Name: "val", Type: Uint256Ty}}
var boolOutput = abi.Arguments{{Name: "val", Type: BoolTy}}

// UnpackUint256 decodes a single uint256 return value.
func UnpackUint256(data []byte) (*big.Int, error) {
	vals, err := uint256Output.Unpack(data)
	if err != nil {
		return nil, err
	}
	return vals[0].(*big.Int), nil
}

// UnpackBool decodes a single bool return value.
func UnpackBool(data []byte) (bool, error) {
	vals, err := boolOutput.Unpack(data)
	if err != nil {
		return false, err
	}
	return vals[0].(bool), nil
}
