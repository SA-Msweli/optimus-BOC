# Deployment — quick guide ✅

This project uses Foundry. The `contracts/script/Deploy.s.sol` script (inside the Foundry project at `contracts/`) deploys the core contracts and wires roles required for DAO ↔ TokenVault ↔ BNPL interactions.

## What the script deploys
- `TokenVault` — on‑chain vault for treasury funds
- `DAOManager` — governance + treasury accounting
- `BNPLManager` — BNPL lifecycle (wired to DAOManager)
- `LoanManager` — loan lifecycle
- `DIDRegistry` — identity registry

It also performs these wiring steps:
- `DAOManager.setTokenVault()` → TokenVault address
- `BNPLManager.setDaoManager()` → DAOManager address
- `TokenVault.grantRole(VAULT_MANAGER_ROLE, DAOManager)`
- `DAOManager.grantRole(TREASURY_FUNDER_ROLE, BNPLManager)`

## How to run (example)

1) Prepare your environment
- Copy `.env.example` -> `.env` and fill values (recommended):
  - `cp .env.example .env`
  - Edit `.env` and set `DEPLOYER_KEY`, `RPC_URL`, and optionally `ETHERSCAN_API_KEY`.
  - Note: `.env.example` is tracked in the repo; `.env` remains ignored (keep secrets out of git).

2) Run the Foundry deploy script (recommended from `contracts/`)
- Change into the Foundry project and run the deploy script:
  - `cd contracts`
  - `forge script script/Deploy.s.sol:Deploy --rpc-url $RPC_URL --private-key $DEPLOYER_KEY --broadcast`
- Key note: the Solidity script reads `DEPLOYER_KEY` via `vm.envUint("DEPLOYER_KEY")` so you can either set `DEPLOYER_KEY` in your environment/.env or pass `--private-key` on the CLI.
- Example (using `.env`):
  - `set -o allexport; source ../.env; set +o allexport; cd contracts && forge script script/Deploy.s.sol:Deploy --rpc-url $RPC_URL --broadcast`

3) Local / integration deploy (Anvil)
- Start Anvil: `anvil --port 8545 &`
- Deploy locally: `cd contracts && forge script script/Deploy.s.sol:Deploy --rpc-url http://127.0.0.1:8545 --private-key <ANVIL_KEY> --broadcast`
- Or run the bundled smoke check: `./contracts/scripts/integration.sh`

4) Verify on-chain (Etherscan)
- Set `ETHERSCAN_API_KEY` in `.env` or your environment.
- Use Foundry to verify sources (example for `DAOManager`):
  - `forge verify-contract --chain sepolia <DAO_ADDRESS> contracts/src/DAOManager.sol:DAOManager $ETHERSCAN_API_KEY`
- Repeat for the other contracts. See `forge verify-contract --help` for constructor/fully-qualified-name details.

5) Artifacts & broadcast logs
- Forge saves broadcast traces to: `contracts/broadcast/<ScriptName>/<chainId>/run-*.json` (contains tx hashes).  
- ABIs / compiled artifacts: `contracts/out/<Contract>.json`  
- Cache (sensitive values redacted): `contracts/cache/` — useful when debugging.

Notes & troubleshooting
- Run the deploy from the `contracts/` directory so remappings in `contracts/foundry.toml` resolve imports (`forge-std`, `openzeppelin-contracts`).  
- If you see "Source ... not found" or missing `forge-std`/`openzeppelin-contracts`, run `git submodule update --init --recursive` or `forge install` to populate `lib/`.  
- If RPC/import errors persist, verify `RPC_URL`, `DEPLOYER_KEY` and that `anvil`/`forge`/`cast` are installed and in PATH.
- `.env` is ignored by git — never commit private keys. Rotate keys immediately if leaked.

## Post-deploy checks
- Verify `TokenVault` has `VAULT_MANAGER_ROLE` granted to `DAOManager`.
- Verify `DAOManager.tokenVault()` points to the deployed vault.
- Verify `BNPLManager.daoManager()` points to the deployed DAOManager.

You can run the unit test that asserts these invariants:
  - forge test --match-contract DeploymentTest

## Notes / safety
- The deployer receives DEFAULT_ADMIN_ROLE for all contracts (constructor behavior).
- Adjust role grants in the script if you need different admin/ownership semantics.
- Always test on a devnet/testnet (Sepolia) before mainnet.
