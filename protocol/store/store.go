package store

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/optimus-boc-protocol/models"
)

// Store wraps a pgxpool to provide simple persistence helpers used by the
// HTTP controllers.  We only persist the minimal data needed for CRE
// workflows and off-chain lookups; most state is kept on-chain.

type Store struct {
    pool *pgxpool.Pool
}

// New creates a Store backed by the given pool.
func New(pool *pgxpool.Pool) *Store {
    return &Store{pool: pool}
}

// SaveArrangement inserts or updates an arrangement record.  The passed
// models.Arrangement should include a non-zero ArrangementID (on-chain id).
func (s *Store) SaveArrangement(ctx context.Context, arr models.Arrangement) error {
    // upsert on arrangement_id
    query := `
INSERT INTO arrangements (arrangement_id, dao_id, payer, recipient, total_amount, num_installments, start_timestamp, interval_seconds, status)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
ON CONFLICT (arrangement_id) DO UPDATE SET
    dao_id = EXCLUDED.dao_id,
    payer = EXCLUDED.payer,
    recipient = EXCLUDED.recipient,
    total_amount = EXCLUDED.total_amount,
    num_installments = EXCLUDED.num_installments,
    start_timestamp = EXCLUDED.start_timestamp,
    interval_seconds = EXCLUDED.interval_seconds,
    status = EXCLUDED.status
` 
    _, err := s.pool.Exec(ctx, query,
        arr.ArrangementID,
        arr.DaoID,
        arr.Payer,
        arr.Recipient,
        arr.TotalAmount,
        arr.NumInstallments,
        arr.StartTimestamp,
        arr.IntervalSeconds,
        arr.Status,
    )
    return err
}

// UpdateArrangementStatus sets a new status string for an existing record.
func (s *Store) UpdateArrangementStatus(ctx context.Context, arrangementID string, status string) error {
    _, err := s.pool.Exec(ctx, "UPDATE arrangements SET status=$1 WHERE arrangement_id=$2", status, arrangementID)
    if err != nil {
        return fmt.Errorf("update arrangement status: %w", err)
    }
    return nil
}

// SaveDAO records information about a DAO when it is created.
func (s *Store) SaveDAO(ctx context.Context, dao models.DAO) error {
    query := `
INSERT INTO daos (dao_id, creator, goal, voting_period_seconds, treasury_balance, total_investments)
VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (dao_id) DO UPDATE SET
    creator = EXCLUDED.creator,
    goal = EXCLUDED.goal,
    voting_period_seconds = EXCLUDED.voting_period_seconds,
    treasury_balance = EXCLUDED.treasury_balance,
    total_investments = EXCLUDED.total_investments
`
    _, err := s.pool.Exec(ctx, query,
        dao.DaoID,
        dao.Creator,
        dao.Goal,
        dao.VotingPeriodSeconds,
        dao.TreasuryBalance,
        dao.TotalInvestments,
    )
    return err
}

// SaveDIDProfile inserts or updates a DID profile record in the database.
func (s *Store) SaveDIDProfile(ctx context.Context, p models.DIDProfile) error {
    query := `
INSERT INTO did_profiles (owner, risk_score, last_activity)
VALUES ($1,$2,$3)
ON CONFLICT (owner) DO UPDATE SET
    risk_score = EXCLUDED.risk_score,
    last_activity = EXCLUDED.last_activity
`
    _, err := s.pool.Exec(ctx, query, p.Owner, p.RiskScore, p.LastActive)
    return err
}

// SaveLoan inserts or updates a loan record.
func (s *Store) SaveLoan(ctx context.Context, l models.Loan) error {
    query := `
INSERT INTO loans (loan_id, borrower, dao_address, principal, interest_rate_bps, start_time, end_time, amount_paid, status)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
ON CONFLICT (loan_id) DO UPDATE SET
    borrower = EXCLUDED.borrower,
    dao_address = EXCLUDED.dao_address,
    principal = EXCLUDED.principal,
    interest_rate_bps = EXCLUDED.interest_rate_bps,
    start_time = EXCLUDED.start_time,
    end_time = EXCLUDED.end_time,
    amount_paid = EXCLUDED.amount_paid,
    status = EXCLUDED.status
`
    _, err := s.pool.Exec(ctx, query,
        l.LoanID,
        l.Borrower,
        l.DaoAddress,
        l.Principal,
        l.InterestRateBps,
        l.StartTime,
        l.EndTime,
        l.AmountPaid,
        l.Status,
    )
    return err
}

// UpdateLoanStatus sets a new status string for an existing loan record.
func (s *Store) UpdateLoanStatus(ctx context.Context, loanID string, status string) error {
    _, err := s.pool.Exec(ctx, "UPDATE loans SET status=$1 WHERE loan_id=$2", status, loanID)
    if err != nil {
        return fmt.Errorf("update loan status: %w", err)
    }
    return nil
}
