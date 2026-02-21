package models

import (
    "time"
)

type Loan struct {
    ID              int64     `db:"id"`
    LoanID          string    `db:"loan_id"`
    Borrower        string    `db:"borrower"`
    DaoAddress      string    `db:"dao_address"`
    Principal       string    `db:"principal"`
    InterestRateBps int       `db:"interest_rate_bps"`
    StartTime       int64     `db:"start_time"`
    EndTime         int64     `db:"end_time"`
    AmountPaid      string    `db:"amount_paid"`
    Status          string    `db:"status"`
    CreatedAt       time.Time `db:"created_at"`
}