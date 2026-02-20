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

## Deployments — Sepolia 🚀

- These addresses come from the latest Sepolia deploy (see `contracts/broadcast/Deploy.s.sol/11155111/run-latest.json`).
- Deployer / admin: `0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780` (default admin on all contracts)
- Block: `10300036` — deployed Feb 20, 2026

### Deployed contracts (Sepolia)
- `TokenVault`  — `0x4C704D51fc47cfe582F8c5477de3AE398B344907` — multi‑asset vault for DAO treasury (VAULT_MANAGER_ROLE → `DAOManager`).
- `DAOManager`  — `0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b` — governance & treasury accounting (TREASURY_FUNDER_ROLE → `BNPLManager`).
- `BNPLManager` — `0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba` — BNPL lifecycle (create/makePayment/applyLateFee/reschedule).
- `LoanManager` — `0xbB0D4067488edf4a007822407e2486412dC8D39D` — loan lifecycle and repayments.
- `DIDRegistry` — `0x0E9D8959bCD99e7AFD7C693e51781058A998b756` — DID storage for on‑chain identities.

Artifacts / verification
- Broadcast JSON: `contracts/broadcast/Deploy.s.sol/11155111/run-latest.json`
- ABIs / build: `contracts/out/<Contract>.json`
- Verification: source verification on Etherscan is *not* performed by the deploy script; use `forge verify-contract` (requires `ETHERSCAN_API_KEY`).

### Quick interaction examples (cast)
- Read a role constant:
  - `cast call --rpc-url $RPC_URL 0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b "TREASURY_FUNDER_ROLE()(bytes32)"`
- Check a role membership (example checks BNPLManager):
  - `TREASURY_HASH=$(cast call --rpc-url $RPC_URL 0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b "TREASURY_FUNDER_ROLE()(bytes32)")`
  - `cast call --rpc-url $RPC_URL 0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b "hasRole(bytes32,address)(bool)" $TREASURY_HASH 0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba`
- Create a DAO (example):
  - `cast send --rpc-url $RPC_URL --private-key $DEPLOYER_KEY 0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b "createDAO(address,uint8,uint64)" 0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780 1 30`
- Check DAO treasury balance:
  - `cast call --rpc-url $RPC_URL 0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b "getTreasuryBalance(uint256)(uint256)" 0`
- Make a BNPL payment (payable):
  - `cast send --rpc-url $RPC_URL --private-key $USER_KEY --value 1ether 0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba "makePayment(uint256,uint256)" <arrangementId> 0`
- Deposit / withdraw from `TokenVault`:
  - `cast send --rpc-url $RPC_URL --private-key $SOME_KEY 0x4C704D51fc47cfe582F8c5477de3AE398B344907 "deposit(address,uint256)" <tokenAddr> <amount>`
  - `cast call --rpc-url $RPC_URL 0x4C704D51fc47cfe582F8c5477de3AE398B344907 "getBalance(address)(uint256)" <tokenAddr>`

Notes & links
- Deployer is the `DEFAULT_ADMIN_ROLE` for all contracts — treat this account as highly privileged. ⚠️
- Role / policy wiring performed by the deploy script (see `contracts/script/Deploy.s.sol`).
- For complete usage examples, see the unit tests under `contracts/test/` (they are the canonical examples).  

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
