# Contracts (Foundry project) 🔧

This folder is the Foundry project for the Optimus smart contracts (Solidity). It contains the core on‑chain components you’ll interact with and deploy: `BNPLManager`, `DAOManager`, `LoanManager`, `TokenVault`, `DIDRegistry` and their interfaces.

## Quick start ✅
- Prerequisites: Foundry (`foundryup`), `anvil`, and `cast` in PATH.
- Build: `cd contracts && forge build --sizes`
- Run tests: `cd contracts && forge test -vvv`

## Local integration / deploy
- Deploy to a local Anvil node (recommended for smoke tests):
  - `cd contracts && forge script script/Deploy.s.sol:Deploy --rpc-url http://127.0.0.1:8545 --private-key <ANVIL_KEY> --broadcast`
- Handy: run the bundled integration check (starts Anvil, deploys, verifies wiring):
  - `./contracts/scripts/integration.sh`

## Useful scripts
- `script/Deploy.s.sol` — Foundry Solidity deploy script (canonical deploy + wiring).
- `scripts/integration.sh` — local Anvil deploy + post‑deploy verifications.
- `scripts/check-comments.sh` — enforces repository comment policy for Solidity sources/tests.

## CI
- Workflow: `contracts/.github/workflows/test.yml` runs the comment check, `forge fmt`, `forge build`, unit tests and an Anvil integration job (deploy + verification).

## Conventions & policy
- Interfaces (`contracts/src/interfaces/*`) are the authoritative API docs and should keep NatSpec where useful.  
- Implementation files: avoid trivial NatSpec; do not use `//` inline comments to explain single lines.  
- Tests must be self‑describing (no `/// @notice` required in tests). Use `scripts/check-comments.sh` locally before opening PRs.

## CRE / events
- CRE workflows consume indexed events emitted by the contracts. Event declarations live in the interfaces (e.g. `IBNPLManager`, `IDAOManager`, `ILoanManager`, `IDIDRegistry`, `ITokenVault`). See `CRE_WORKFLOW_MAPPINGS.md` at the repo root for the event→workflow mapping.

## Troubleshooting ⚠️
- If `forge`/`anvil`/`cast` not found: run `foundryup` to install Foundry.  
- If imports fail, ensure `lib/` contains `forge-std` and `openzeppelin-contracts` (remappings are configured in `foundry.toml`).

---
Keeping this folder self‑contained makes builds, tests and deploys deterministic — `contracts/` is the Foundry project root. If you want, I can add a top‑level `Makefile` shortcut or extend the integration checks. ✨
