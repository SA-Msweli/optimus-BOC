// Package shared – Pre-computed Keccak256 event signature hashes used by
// EVM log trigger filters across all Optimus workflows.
package shared

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Event signature hashes – these are Topics[0] values for filtering.
var (
	SigBNPLCreated        = crypto.Keccak256Hash([]byte("BNPLCreated(uint256,uint256,address,address,uint256,uint256,uint256)"))
	SigBNPLPaymentMade    = crypto.Keccak256Hash([]byte("BNPLPaymentMade(uint256,uint8,address,uint256,uint256)"))
	SigBNPLCompleted      = crypto.Keccak256Hash([]byte("BNPLCompleted(uint256,uint256)"))
	SigLoanCreated        = crypto.Keccak256Hash([]byte("LoanCreated(uint256,address,uint256,uint256,uint256,uint256)"))
	SigLoanApproved       = crypto.Keccak256Hash([]byte("LoanApproved(uint256,address)"))
	SigPaymentMade        = crypto.Keccak256Hash([]byte("PaymentMade(uint256,address,uint256,uint256,uint256)"))
	SigProposalOpened     = crypto.Keccak256Hash([]byte("ProposalOpened(uint256,uint256,uint256,bytes)"))
	SigVoteCast           = crypto.Keccak256Hash([]byte("VoteCast(uint256,address,bool,uint256)"))
	SigRiskProfileUpdated = crypto.Keccak256Hash([]byte("RiskProfileUpdated(address,uint256,bytes32)"))
)

// AddrBytes converts a hex address string to a 20-byte slice.
func AddrBytes(hex string) []byte {
	return common.HexToAddress(hex).Bytes()
}
