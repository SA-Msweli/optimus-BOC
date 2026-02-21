-- initial schema for Optimus backend
-- run with your preferred migration tool (e.g. goose, migrate)

CREATE TABLE IF NOT EXISTS arrangements (
    id SERIAL PRIMARY KEY,
    arrangement_id NUMERIC NOT NULL UNIQUE,
    dao_id NUMERIC NOT NULL,
    payer TEXT NOT NULL,
    recipient TEXT NOT NULL,
    total_amount NUMERIC NOT NULL,
    num_installments INT NOT NULL,
    start_timestamp BIGINT NOT NULL,
    interval_seconds BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    loan_id NUMERIC NOT NULL UNIQUE,
    borrower TEXT NOT NULL,
    dao_address TEXT NOT NULL,
    principal NUMERIC NOT NULL,
    interest_rate_bps INT NOT NULL,
    start_time BIGINT NOT NULL,
    end_time BIGINT NOT NULL,
    amount_paid NUMERIC DEFAULT 0,
    status TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS daos (
    id SERIAL PRIMARY KEY,
    dao_id NUMERIC NOT NULL UNIQUE,
    creator TEXT NOT NULL,
    goal TEXT NOT NULL,
    voting_period_seconds INT NOT NULL,
    treasury_balance NUMERIC DEFAULT 0,
    total_investments NUMERIC DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS did_profiles (
    id SERIAL PRIMARY KEY,
    owner TEXT NOT NULL UNIQUE,
    risk_score INT DEFAULT 0,
    last_activity BIGINT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
