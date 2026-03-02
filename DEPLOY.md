# Deployment Guide

## 1. Smart Contracts (Foundry)

The contracts are deployed via Foundry from the `contracts/` directory.

### Prerequisites

- Foundry (`forge`, `cast`, `anvil`) installed
- `.env` with `DEPLOYER_KEY`, `RPC_URL`, optionally `ETHERSCAN_API_KEY`

### What the deploy script does

Deploys all 5 contracts and performs wiring:

| Contract | Role |
|----------|------|
| `TokenVault` | On-chain vault for treasury funds |
| `DAOManager` | Governance + treasury accounting |
| `BNPLManager` | BNPL lifecycle (wired to DAOManager) |
| `LoanManager` | Loan lifecycle |
| `DIDRegistry` | Identity registry |

Wiring steps:
- `DAOManager.setTokenVault()` → TokenVault address
- `BNPLManager.setDaoManager()` → DAOManager address
- `TokenVault.grantRole(VAULT_MANAGER_ROLE, DAOManager)`
- `DAOManager.grantRole(TREASURY_FUNDER_ROLE, BNPLManager)`

### Deploy to Sepolia

```bash
cd contracts
set -o allexport; source ../.env; set +o allexport
forge script script/Deploy.s.sol:Deploy --rpc-url $RPC_URL --broadcast
```

### Deploy to local Anvil

```bash
anvil --port 8545 &
cd contracts
forge script script/Deploy.s.sol:Deploy \
  --rpc-url http://127.0.0.1:8545 \
  --private-key <ANVIL_KEY> \
  --broadcast
```

### Verify on Etherscan

```bash
forge verify-contract --chain sepolia <ADDRESS> contracts/src/<Contract>.sol:<Contract> $ETHERSCAN_API_KEY
```

### Current Sepolia Deployment

| Contract | Address |
|----------|---------|
| DIDRegistry | `0x0E9D33b0cA7c7e93F4F5e413C35cE0E518040296` |
| BNPLManager | `0x4d99F01530F9b5Ee58C13E52a114d79cF397c1f4` |
| LoanManager | `0xbB0D4067488edf4a007822407e2486412dC8D39D` |
| DAOManager | `0x5612dB018f17E82D76E79a77aA75Bbf42E40C9B1` |
| TokenVault | `0x4C70E0Ae715d04e3a6e65a70191276a5D0E5B66a` |
| Deployer | `0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780` |

---

## 2. Backend (Go API Server on EC2)

The Go backend is deployed to an EC2 instance as a systemd service.

### One-time EC2 setup

```bash
./scripts/setup-ec2.sh ubuntu@13.60.166.148 optimus.pem
```

This installs Go, creates the `optimus` system user, sets up PostgreSQL, installs the `optimus.service` systemd unit, and provisions `/etc/optimus/backend.env` from the template.

After setup, SSH in and fill the env file:

```bash
ssh -i optimus.pem ubuntu@13.60.166.148
sudo nano /etc/optimus/backend.env
# Fill in: CHAIN_RPC_URL, PRIVATE_KEY, PRIVY_APP_ID, PRIVY_APP_SECRET
```

### Deploy / redeploy

```bash
./scripts/deploy-ec2.sh ubuntu@13.60.166.148 optimus.pem
```

This syncs protocol source, runs DB migrations, builds the Go binary on EC2, and restarts the service. It runs a health check at the end.

### What's on EC2

```
/home/optimus/
├── optimus-backend          # compiled Go binary (active)
├── optimus-backend-new      # newly built binary (swapped on deploy)
└── optimus-BOC/protocol/    # synced source (for remote build)

/etc/optimus/backend.env     # environment variables (chmod 600)
/etc/systemd/system/optimus.service
/var/log/optimus/backend.log
```

### Backend management

```bash
sudo systemctl status optimus     # check status
sudo systemctl restart optimus    # restart
sudo journalctl -u optimus -f     # follow logs
curl http://localhost:8000/health  # health check
```

### Environment variables

See `scripts/backend.env.example` for the full list. Required:

| Variable | Description |
|----------|-------------|
| `DATABASE_URL` | PostgreSQL connection string |
| `CHAIN_RPC_URL` | Sepolia RPC endpoint |
| `PRIVATE_KEY` | Deployer private key for signing transactions |
| `BNPL_MANAGER_ADDRESS` | BNPLManager contract address |
| `LOAN_MANAGER_ADDRESS` | LoanManager contract address |
| `DAO_MANAGER_ADDRESS` | DAOManager contract address |
| `DID_REGISTRY_ADDRESS` | DIDRegistry contract address |
| `TOKEN_VAULT_ADDRESS` | TokenVault contract address |
| `PRIVY_APP_ID` | Privy dashboard app identifier |
| `PRIVY_APP_SECRET` | Privy server-side API secret |

---

## 3. CRE Workflows (Chainlink DON)

CRE workflows run as WASM binaries on the Chainlink DON — they are **not deployed to EC2**.

### Build

```bash
cd workflows
GOOS=wasip1 GOARCH=wasm go build ./...
```

### Deploy

Workflows are deployed via the CRE CLI using the spec in `workflows/project.yaml`. See Chainlink CRE documentation for deployment commands.

### Simulate (local testing)

```bash
cre workflow simulate <workflow-name> --target staging-settings
```

---

## Scripts Reference

| Script | Purpose |
|--------|---------|
| `scripts/setup-ec2.sh` | One-time EC2 provisioning (Go, PostgreSQL, systemd, env) |
| `scripts/deploy-ec2.sh` | Deploy/redeploy backend to EC2 |
| `scripts/optimus.service` | systemd unit file for the backend |
| `scripts/backend.env.example` | Environment variable template |

---

## Post-deploy checks

- `curl http://<host>:8000/health` returns `OK`
- `TokenVault` has `VAULT_MANAGER_ROLE` granted to `DAOManager`
- `DAOManager.tokenVault()` points to the deployed vault
- `BNPLManager.daoManager()` points to the deployed DAOManager

Run the invariant test:

```bash
cd contracts && forge test --match-contract DeploymentTest
```

## Notes

- The deployer receives `DEFAULT_ADMIN_ROLE` for all contracts (constructor behavior)
- `.env` is gitignored — never commit private keys
- Always test on Sepolia before mainnet
