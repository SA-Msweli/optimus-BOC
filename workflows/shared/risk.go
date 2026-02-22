package shared

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// ──────────────────────────────────────────────
// Risk score constants and computation
// ──────────────────────────────────────────────

const (
	MaxRiskScore = 10000 // basis-point ceiling
	MinRiskScore = 0

	// Credit tier thresholds (basis points)
	TierExcellent = 7000
	TierGood      = 5000
	TierFair      = 3000
)

// Risk score adjustment deltas (basis points).
const (
	AdjBNPLCompletedOnTime   int64 = 500  // +5% for on-time BNPL completion
	AdjBNPLLateFee           int64 = -300 // -3% when a late fee is applied
	AdjLoanRepaidFull        int64 = 700  // +7% for full loan repayment
	AdjLoanDefaulted         int64 = -2000 // -20% for loan default
	AdjLoanPaymentOnTime     int64 = 100  // +1% per on-time loan payment
	AdjBNPLPaymentOnTime     int64 = 50   // +0.5% per on-time BNPL installment
)

// CreditTier returns a human-readable tier label for the given score.
func CreditTier(score *big.Int) string {
	s := score.Int64()
	switch {
	case s >= TierExcellent:
		return "EXCELLENT"
	case s >= TierGood:
		return "GOOD"
	case s >= TierFair:
		return "FAIR"
	default:
		return "POOR"
	}
}

// ClampScore ensures a risk score stays within [0, 10000].
func ClampScore(score int64) *big.Int {
	if score > MaxRiskScore {
		score = MaxRiskScore
	}
	if score < MinRiskScore {
		score = MinRiskScore
	}
	return big.NewInt(score)
}

// AdjustScore applies a delta to a current risk score and clamps the result.
func AdjustScore(current *big.Int, delta int64) *big.Int {
	updated := new(big.Int).Add(current, big.NewInt(delta))
	return ClampScore(updated.Int64())
}

// ComputeProfileHash computes the keccak256 hash used as the profileHash
// parameter for DIDRegistry.updateRiskProfile.  It combines the owner
// address, the new score, and a descriptive reason string.
func ComputeProfileHash(owner [20]byte, newScore *big.Int, reason string) [32]byte {
	data := append(owner[:], newScore.Bytes()...)
	data = append(data, []byte(reason)...)
	hash := crypto.Keccak256(data)
	var out [32]byte
	copy(out[:], hash)
	return out
}

// MaxBNPLAmount returns the maximum BNPL arrangement value for a credit tier.
func MaxBNPLAmount(score *big.Int) *big.Int {
	tier := CreditTier(score)
	switch tier {
	case "EXCELLENT":
		v, _ := new(big.Int).SetString("10000000000000000000", 10)
		return v // 10 ETH
	case "GOOD":
		return big.NewInt(5_000_000_000_000_000_000) // 5 ETH
	case "FAIR":
		return big.NewInt(1_000_000_000_000_000_000) // 1 ETH
	default:
		return big.NewInt(500_000_000_000_000_000) // 0.5 ETH
	}
}

// MaxLoanPrincipal returns the maximum loan principal for a credit tier.
func MaxLoanPrincipal(score *big.Int) *big.Int {
	tier := CreditTier(score)
	switch tier {
	case "EXCELLENT":
		v, _ := new(big.Int).SetString("50000000000000000000", 10)
		return v // 50 ETH
	case "GOOD":
		v, _ := new(big.Int).SetString("20000000000000000000", 10)
		return v // 20 ETH
	case "FAIR":
		return big.NewInt(5_000_000_000_000_000_000) // 5 ETH
	default:
		return big.NewInt(1_000_000_000_000_000_000) // 1 ETH
	}
}

// FormatWei returns a human-readable string for a wei amount (approximate ETH).
func FormatWei(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), new(big.Float).SetInt64(1e18))
	return fmt.Sprintf("%.4f ETH", eth)
}
