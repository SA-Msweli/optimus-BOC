package models

import (
    "time"
)

// Arrangement represents a BNPL arrangement stored off-chain.
// Note: many fields duplicated from on-chain state for fast queries.

type Arrangement struct {
    ID              int64     `db:"id"`
    ArrangementID   string    `db:"arrangement_id"`
    DaoID           string    `db:"dao_id"`
    Payer           string    `db:"payer"`
    Recipient       string    `db:"recipient"`
    TotalAmount     string    `db:"total_amount"`
    NumInstallments int       `db:"num_installments"`
    StartTimestamp  int64     `db:"start_timestamp"`
    IntervalSeconds int64     `db:"interval_seconds"`
    Status          string    `db:"status"`
    CreatedAt       time.Time `db:"created_at"`
}
