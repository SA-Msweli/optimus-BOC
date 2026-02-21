package models

import "time"

// DAO represents minimal off-chain record of a DAO entity.

type DAO struct {
    ID                 int64     `db:"id"`
    DaoID              string    `db:"dao_id"`
    Creator            string    `db:"creator"`
    Goal               string    `db:"goal"`
    VotingPeriodSeconds int      `db:"voting_period_seconds"`
    TreasuryBalance    string    `db:"treasury_balance"`
    TotalInvestments   string    `db:"total_investments"`
    CreatedAt          time.Time `db:"created_at"`
}