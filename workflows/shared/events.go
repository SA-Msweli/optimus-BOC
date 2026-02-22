package shared

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ──────────────────────────────────────────────
// Event data decoders
// ──────────────────────────────────────────────
// For each event, indexed parameters live in Topics[1..3]
// (Topics[0] = event signature hash).  Non-indexed params are
// ABI-encoded in the Data field.

// ── BNPLManager events ──

// BNPLPaymentData holds the non-indexed fields of BNPLPaymentMade.
type BNPLPaymentData struct {
	InstallmentNumber uint8
	Amount            *big.Int
	Timestamp         *big.Int
}

var bnplPaymentDataArgs = abi.Arguments{
	{Name: "installmentNumber", Type: Uint8Ty},
	{Name: "amount", Type: Uint256Ty},
	{Name: "timestamp", Type: Uint256Ty},
}

// DecodeBNPLPaymentData decodes the Data field of a BNPLPaymentMade log.
// Indexed: arrangementId (Topics[1]), payer (Topics[2]).
func DecodeBNPLPaymentData(data []byte) (*BNPLPaymentData, error) {
	vals, err := bnplPaymentDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &BNPLPaymentData{
		InstallmentNumber: vals[0].(uint8),
		Amount:            vals[1].(*big.Int),
		Timestamp:         vals[2].(*big.Int),
	}, nil
}

// BNPLCreatedData holds the non-indexed fields of BNPLCreated.
type BNPLCreatedData struct {
	Recipient       common.Address
	TotalAmount     *big.Int
	NumInstallments *big.Int
	CreatedAt       *big.Int
}

var bnplCreatedDataArgs = abi.Arguments{
	{Name: "recipient", Type: AddressTy},
	{Name: "totalAmount", Type: Uint256Ty},
	{Name: "numInstallments", Type: Uint256Ty},
	{Name: "createdAt", Type: Uint256Ty},
}

// DecodeBNPLCreatedData decodes the Data field of a BNPLCreated log.
// Indexed: arrangementId (Topics[1]), daoId (Topics[2]), payer (Topics[3]).
func DecodeBNPLCreatedData(data []byte) (*BNPLCreatedData, error) {
	vals, err := bnplCreatedDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &BNPLCreatedData{
		Recipient:       vals[0].(common.Address),
		TotalAmount:     vals[1].(*big.Int),
		NumInstallments: vals[2].(*big.Int),
		CreatedAt:       vals[3].(*big.Int),
	}, nil
}

// BNPLLateFeeData holds the non-indexed fields of BNPLLateFeeApplied.
type BNPLLateFeeData struct {
	InstallmentNumber uint8
	FeeAmount         *big.Int
	Timestamp         *big.Int
}

var bnplLateFeeDataArgs = abi.Arguments{
	{Name: "installmentNumber", Type: Uint8Ty},
	{Name: "feeAmount", Type: Uint256Ty},
	{Name: "timestamp", Type: Uint256Ty},
}

// DecodeBNPLLateFeeData decodes the Data field of BNPLLateFeeApplied.
// Indexed: arrangementId (Topics[1]).
func DecodeBNPLLateFeeData(data []byte) (*BNPLLateFeeData, error) {
	vals, err := bnplLateFeeDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &BNPLLateFeeData{
		InstallmentNumber: vals[0].(uint8),
		FeeAmount:         vals[1].(*big.Int),
		Timestamp:         vals[2].(*big.Int),
	}, nil
}

// BNPLCompletedData holds the non-indexed field of BNPLCompleted.
type BNPLCompletedData struct {
	CompletedAt *big.Int
}

var bnplCompletedDataArgs = abi.Arguments{
	{Name: "completedAt", Type: Uint256Ty},
}

// DecodeBNPLCompletedData decodes the Data field of BNPLCompleted.
// Indexed: arrangementId (Topics[1]).
func DecodeBNPLCompletedData(data []byte) (*BNPLCompletedData, error) {
	vals, err := bnplCompletedDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &BNPLCompletedData{CompletedAt: vals[0].(*big.Int)}, nil
}

// ── LoanManager events ──

// LoanCreatedData holds the non-indexed fields of LoanCreated.
type LoanCreatedData struct {
	Principal       *big.Int
	InterestRateBps *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
}

var loanCreatedDataArgs = abi.Arguments{
	{Name: "principal", Type: Uint256Ty},
	{Name: "interestRateBps", Type: Uint256Ty},
	{Name: "startTime", Type: Uint256Ty},
	{Name: "endTime", Type: Uint256Ty},
}

// DecodeLoanCreatedData decodes the Data field of LoanCreated.
// Indexed: loanId (Topics[1]), borrower (Topics[2]).
func DecodeLoanCreatedData(data []byte) (*LoanCreatedData, error) {
	vals, err := loanCreatedDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &LoanCreatedData{
		Principal:       vals[0].(*big.Int),
		InterestRateBps: vals[1].(*big.Int),
		StartTime:       vals[2].(*big.Int),
		EndTime:         vals[3].(*big.Int),
	}, nil
}

// LoanPaymentData holds the non-indexed fields of PaymentMade (LoanManager).
type LoanPaymentData struct {
	Amount     *big.Int
	AmountPaid *big.Int
	Remaining  *big.Int
}

var loanPaymentDataArgs = abi.Arguments{
	{Name: "amount", Type: Uint256Ty},
	{Name: "amountPaid", Type: Uint256Ty},
	{Name: "remaining", Type: Uint256Ty},
}

// DecodeLoanPaymentData decodes the Data field of PaymentMade.
// Indexed: loanId (Topics[1]), payer (Topics[2]).
func DecodeLoanPaymentData(data []byte) (*LoanPaymentData, error) {
	vals, err := loanPaymentDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &LoanPaymentData{
		Amount:     vals[0].(*big.Int),
		AmountPaid: vals[1].(*big.Int),
		Remaining:  vals[2].(*big.Int),
	}, nil
}

// ── DAOManager events ──

// ProposalOpenedData holds the non-indexed fields of ProposalOpened.
type ProposalOpenedData struct {
	Expiry *big.Int
	Data   []byte
}

var proposalOpenedDataArgs = abi.Arguments{
	{Name: "expiry", Type: Uint256Ty},
	{Name: "data", Type: BytesTy},
}

// DecodeProposalOpenedData decodes the Data field of ProposalOpened.
// Indexed: proposalId (Topics[1]), daoId (Topics[2]).
func DecodeProposalOpenedData(data []byte) (*ProposalOpenedData, error) {
	vals, err := proposalOpenedDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &ProposalOpenedData{
		Expiry: vals[0].(*big.Int),
		Data:   vals[1].([]byte),
	}, nil
}

// VoteCastData holds the non-indexed fields of VoteCast.
type VoteCastData struct {
	Support bool
	Weight  *big.Int
}

var voteCastDataArgs = abi.Arguments{
	{Name: "support", Type: BoolTy},
	{Name: "weight", Type: Uint256Ty},
}

// DecodeVoteCastData decodes the Data field of VoteCast.
// Indexed: proposalId (Topics[1]), voter (Topics[2]).
func DecodeVoteCastData(data []byte) (*VoteCastData, error) {
	vals, err := voteCastDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &VoteCastData{
		Support: vals[0].(bool),
		Weight:  vals[1].(*big.Int),
	}, nil
}

// ── DIDRegistry events ──

// RiskProfileUpdatedData holds the non-indexed fields of RiskProfileUpdated.
type RiskProfileUpdatedData struct {
	NewScore *big.Int
}

var riskProfileUpdatedDataArgs = abi.Arguments{
	{Name: "newScore", Type: Uint256Ty},
}

// DecodeRiskProfileUpdatedData decodes the Data field of RiskProfileUpdated.
// Indexed: owner (Topics[1]), profileHash (Topics[2]).
func DecodeRiskProfileUpdatedData(data []byte) (*RiskProfileUpdatedData, error) {
	vals, err := riskProfileUpdatedDataArgs.Unpack(data)
	if err != nil {
		return nil, err
	}
	return &RiskProfileUpdatedData{NewScore: vals[0].(*big.Int)}, nil
}

// ── Helpers to extract indexed parameters from topics ──

// TopicToAddress extracts an address from a 32-byte topic ([]byte from evm.Log).
func TopicToAddress(topic []byte) common.Address {
	return common.BytesToAddress(topic)
}

// TopicToBigInt extracts a *big.Int from a 32-byte topic ([]byte from evm.Log).
func TopicToBigInt(topic []byte) *big.Int {
	return new(big.Int).SetBytes(topic)
}
