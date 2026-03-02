#!/usr/bin/env bash
# ===========================================================================
# One-time EC2 setup for Optimus backend deployment
#
# Installs Go, creates the optimus system user, sets up PostgreSQL,
# installs the systemd service, and provisions the env file template.
#
# Usage:
#   ./scripts/setup-ec2.sh <user@host> [key.pem]
#
# Example:
#   ./scripts/setup-ec2.sh ubuntu@13.60.166.148 optimus.pem
# ===========================================================================
set -euo pipefail

REMOTE=${1:?Usage: $0 <user@host> [key.pem]}
KEY=${2:-optimus.pem}
SSH="ssh -o StrictHostKeyChecking=no -i $KEY"
SCP="scp -o StrictHostKeyChecking=no -i $KEY"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "============================================"
echo "  Optimus EC2 Setup"
echo "  Target: $REMOTE"
echo "============================================"

# ── 1. Install / upgrade Go ──────────────────────────────────────────────
echo ""
echo "=== 1/6  Installing Go 1.25.7 ==="
$SSH "$REMOTE" "sudo rm -rf /usr/local/go && curl -fsSL https://go.dev/dl/go1.25.7.linux-amd64.tar.gz | sudo tar -C /usr/local -xz && echo 'export PATH=/usr/local/go/bin:\$PATH' | sudo tee /etc/profile.d/go.sh > /dev/null && /usr/local/go/bin/go version"

# ── 2. Create optimus system user ────────────────────────────────────────
echo ""
echo "=== 2/6  Creating optimus system user ==="
$SSH "$REMOTE" "id optimus 2>/dev/null || sudo useradd -m -s /bin/bash optimus"

# ── 3. Create directories ────────────────────────────────────────────────
echo ""
echo "=== 3/6  Creating directories ==="
$SSH "$REMOTE" "sudo mkdir -p /home/optimus/optimus-BOC/protocol /etc/optimus /var/log/optimus"
$SSH "$REMOTE" "sudo chown -R optimus:optimus /home/optimus/optimus-BOC /var/log/optimus"

# ── 4. Set up PostgreSQL ─────────────────────────────────────────────────
echo ""
echo "=== 4/6  Setting up PostgreSQL ==="
$SSH "$REMOTE" "sudo -u postgres psql -tc \"SELECT 1 FROM pg_database WHERE datname='optimus'\" | grep -q 1 || sudo -u postgres createdb optimus"
$SSH "$REMOTE" "sudo -u postgres psql -tc \"SELECT 1 FROM pg_roles WHERE rolname='optimus'\" | grep -q 1 || sudo -u postgres psql -c \"CREATE ROLE optimus WITH LOGIN PASSWORD 'MASTER_Optimust#BOC';\""
$SSH "$REMOTE" "sudo -u postgres psql -c \"GRANT ALL PRIVILEGES ON DATABASE optimus TO optimus;\""

# ── 5. Install systemd service ───────────────────────────────────────────
echo ""
echo "=== 5/6  Installing systemd service ==="
$SCP "$SCRIPT_DIR/optimus.service" "$REMOTE:/tmp/optimus.service"
$SSH "$REMOTE" "sudo cp /tmp/optimus.service /etc/systemd/system/optimus.service && sudo systemctl daemon-reload && sudo systemctl enable optimus"

# ── 6. Provision env file template ───────────────────────────────────────
echo ""
echo "=== 6/6  Provisioning environment file ==="
$SCP "$SCRIPT_DIR/backend.env.example" "$REMOTE:/tmp/backend.env.example"
$SSH "$REMOTE" "if [ ! -f /etc/optimus/backend.env ]; then sudo cp /tmp/backend.env.example /etc/optimus/backend.env && sudo chmod 600 /etc/optimus/backend.env && echo 'Created /etc/optimus/backend.env from template — EDIT IT NOW'; else echo '/etc/optimus/backend.env already exists — skipping'; fi"

echo ""
echo "============================================"
echo "  EC2 setup complete!"
echo "============================================"
echo ""
echo "NEXT STEPS:"
echo "  1. SSH to EC2:  ssh -i $KEY $REMOTE"
echo "  2. Edit env:    sudo nano /etc/optimus/backend.env"
echo "     Fill in: CHAIN_RPC_URL, PRIVATE_KEY, PRIVY_APP_ID, PRIVY_APP_SECRET"
echo "  3. Deploy:      ./scripts/deploy-ec2.sh $REMOTE $KEY"
