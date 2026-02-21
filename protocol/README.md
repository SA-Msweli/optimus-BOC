# Optimus Protocol Backend

This directory contains the Go backend service for the Optimus protocol. It
provides a simple HTTP API and interfaces with the on‑chain contracts via
`go-ethereum` bindings.

## Prerequisites

- Go 1.24
- `abigen` (from go-ethereum) – used to generate contract bindings
  ```bash
  go install github.com/ethereum/go-ethereum/cmd/abigen@latest
  ```
- PostgreSQL (or any database supported by `pgx`)

## Generating ABI Bindings

Before building the service you must generate Go bindings from the Solidity
ABIs compiled by Foundry. Run from this directory:

```bash
cd protocol
./generate_bindings.sh
```

The script will look for JSON files under `../contracts/out/*/*.json` and
generate corresponding Go types in `protocol/bindings/`. If `abigen` is not
installed the script will exit with an error and instructions.

## Configuration

Environment variables used by the service:

| Variable               | Description |
|------------------------|-------------|
| `PORT`                 | HTTP listen port (default `8000`) |
| `DATABASE_URL`         | PostgreSQL connection string |
| `CHAIN_RPC_URL`        | Ethereum JSON-RPC endpoint |
| `BNPL_MANAGER_ADDRESS` | Deployed BNPLManager contract address |

Additional contract addresses (DAOManager, LoanManager, etc.) can be added
as needed and passed into their respective services.

## Building and Running

```bash
cd protocol
go build -o optimus
BNPL_MANAGER_ADDRESS=<addr> CHAIN_RPC_URL=<url> \
    DATABASE_URL=postgres://... ./optimus
```

For containerised builds the provided `Dockerfile` installs `abigen` during
the build stage, so you can regenerate bindings by mounting the source
and running the script inside the container if necessary.

## Services

The `services` package contains Go wrappers around each core contract. They
are lightweight and currently only provide read methods; expand them as your
backend logic grows.

Example services include:

- `services.BNPLService`
- `services.DAOService`
- `services.LoanService`

Additional business‑logic layers (repository patterns, CQRS, etc.) can be
built on top of these services.

---

This README provides the basic workflow to get the backend up and running with
contract bindings and simple HTTP endpoints. Feel free to extend as the
project evolves.
