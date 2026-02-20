#!/usr/bin/env bash
set -euo pipefail

# Local integration helper: start Anvil, run Deploy.s.sol, verify basic post-deploy invariants.
# Requires: anvil, forge, cast available in PATH.
# Usage: ./contracts/scripts/integration.sh

RPC_URL="http://127.0.0.1:8545"
ANVIL_LOG="/tmp/optimus-anvil.log"

echo "Starting Anvil..."
nohup anvil --block-time 1 > "$ANVIL_LOG" 2>&1 &
ANVIL_PID=$!
trap 'echo "Killing anvil ($ANVIL_PID)"; kill $ANVIL_PID || true' EXIT
sleep 1

echo "Extracting private key from Anvil log..."
PK=$(grep -A 1 "Private Keys" "$ANVIL_LOG" | tail -n1 | awk '{print $1}')
if [ -z "$PK" ]; then
  echo "Could not read Anvil private key — is anvil running and printing keys?"
  tail -n 200 "$ANVIL_LOG" || true
  exit 1
fi

echo "Using PK=$PK"

echo "Running deploy script against Anvil"
forge script script/Deploy.s.sol:Deploy --rpc-url "$RPC_URL" --private-key "$PK" --broadcast | tee /tmp/optimus-deploy.log

DEPLOY_LOG=/tmp/optimus-deploy.log
DAO_ADDR=$(grep "DAOManager:" "$DEPLOY_LOG" | awk '{print $2}')
BNPL_ADDR=$(grep "BNPLManager:" "$DEPLOY_LOG" | awk '{print $2}')
VAULT_ADDR=$(grep "TokenVault:" "$DEPLOY_LOG" | awk '{print $2}')

echo "Verifying post-deploy invariants..."

echo " - dao.tokenVault() == vault"
TV=$(cast call $DAO_ADDR "tokenVault()(address)" --rpc-url $RPC_URL)
if [ "$TV" != "$VAULT_ADDR" ]; then
  echo "Mismatch: dao.tokenVault()=$TV expected=$VAULT_ADDR"
  exit 1
fi

echo " - bnpl.daoManager() == dao"
DM=$(cast call $BNPL_ADDR "daoManager()(address)" --rpc-url $RPC_URL)
if [ "$DM" != "$DAO_ADDR" ]; then
  echo "Mismatch: bnpl.daoManager()=$DM expected=$DAO_ADDR"
  exit 1
fi

TREASURY_HASH=$(cast keccak "TREASURY_FUNDER_ROLE")
HAS_TREASURY=$(cast call $DAO_ADDR "hasRole(bytes32,address)(bool)" $TREASURY_HASH $BNPL_ADDR --rpc-url $RPC_URL)
if ! echo "$HAS_TREASURY" | grep -q true; then
  echo "DAO does not have TREASURY_FUNDER_ROLE granted to BNPLManager"
  exit 1
fi

VAULT_HASH=$(cast keccak "VAULT_MANAGER_ROLE")
HAS_VAULT=$(cast call $VAULT_ADDR "hasRole(bytes32,address)(bool)" $VAULT_HASH $DAO_ADDR --rpc-url $RPC_URL)
if ! echo "$HAS_VAULT" | grep -q true; then
  echo "TokenVault does not have VAULT_MANAGER_ROLE granted to DAOManager"
  exit 1
fi

echo "Integration checks passed."
