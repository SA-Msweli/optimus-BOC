package models

import "time"

// DIDProfile stores basic DID information off-chain.

type DIDProfile struct {
    ID         int64     `db:"id"`
    Owner      string    `db:"owner"`
    RiskScore  int       `db:"risk_score"`
    LastActive int64     `db:"last_activity"`
    CreatedAt  time.Time `db:"created_at"`
}