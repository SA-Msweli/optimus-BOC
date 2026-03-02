#!/usr/bin/env bash
# ===========================================================================
# Deploy Optimus protocol backend to EC2
#
# Prerequisites:
#   - EC2 provisioned with ./scripts/setup-ec2.sh
#   - /etc/optimus/backend.env populated on EC2
#
# Usage:
#   ./scripts/deploy-ec2.sh <user@host> [key.pem]
#
# Example:
#   ./scripts/deploy-ec2.sh ubuntu@13.60.166.148 optimus.pem
# ===========================================================================
set -euo pipefail

REMOTE=${1:?Usage: $0 <user@host> [key.pem]}
KEY=${2:-optimus.pem}
SSH="ssh -o StrictHostKeyChecking=no -i $KEY"
RSYNC="rsync -avz --exclude=.git -e 'ssh -i $KEY'"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT="$(dirname "$SCRIPT_DIR")"

echo "============================================"
echo "  Optimus Backend Deployment"
echo "  Target: $REMOTE"
echo "============================================"

# ── 1. Sync protocol source ──────────────────────────────────────────────
echo ""
echo "=== 1/5  Syncing protocol source ==="
eval $RSYNC "$ROOT/protocol/" "$REMOTE:/tmp/protocol-sync/"
$SSH "$REMOTE" "sudo rsync -a /tmp/protocol-sync/ /home/optimus/optimus-BOC/protocol/ && sudo chown -R optimus:optimus /home/optimus/optimus-BOC/protocol/"

# ── 2. Run DB migrations ─────────────────────────────────────────────────
echo ""
echo "=== 2/5  Running DB migrations ==="
$SSH "$REMOTE" "sudo -u postgres psql -d optimus < /home/optimus/optimus-BOC/protocol/db/migrations/001_init.sql 2>&1 || echo 'Migration already applied or tables exist'"

# ── 3. Build backend binary ──────────────────────────────────────────────
echo ""
echo "=== 3/5  Building protocol backend ==="
$SSH "$REMOTE" "sudo -u optimus bash -c 'export PATH=/usr/local/go/bin:\$PATH HOME=/home/optimus GOPATH=/home/optimus/go && cd /home/optimus/optimus-BOC/protocol && go build -o /home/optimus/optimus-backend-new . 2>&1'"

# ── 4. Install systemd service (idempotent) ──────────────────────────────
echo ""
echo "=== 4/5  Installing systemd service ==="
eval $RSYNC "$SCRIPT_DIR/optimus.service" "$REMOTE:/tmp/optimus.service"
$SSH "$REMOTE" "sudo cp /tmp/optimus.service /etc/systemd/system/optimus.service && sudo systemctl daemon-reload && sudo systemctl enable optimus"

# ── 5. Deploy backend ────────────────────────────────────────────────────
echo ""
echo "=== 5/5  Deploying backend binary ==="
$SSH "$REMOTE" "sudo systemctl stop optimus 2>/dev/null || true"
$SSH "$REMOTE" "sudo -u optimus cp /home/optimus/optimus-backend-new /home/optimus/optimus-backend"
$SSH "$REMOTE" "sudo systemctl start optimus && sleep 2"
$SSH "$REMOTE" "curl -sf http://localhost:8000/health && echo ' ✓ Backend healthy' || echo ' ✗ Health check failed'"

echo ""
echo "============================================"
echo "  Deployment complete!"
echo "============================================"
echo ""
echo "Backend: http://$( echo "$REMOTE" | cut -d@ -f2):8000"
echo "Health:  http://$( echo "$REMOTE" | cut -d@ -f2):8000/health"
echo "Logs:    sudo journalctl -u optimus -f"
echo "         or: tail -f /var/log/optimus/backend.log"
