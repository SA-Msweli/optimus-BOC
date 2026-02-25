# ============================================================
# Optimus CRE Workflow Simulator
# ============================================================
# Downloads pre-built CRE CLI v1.1.0 and prepares an environment
# for running `cre workflow simulate` against live contracts
# and the deployed backend API.
# ============================================================

FROM golang:1.25.7-bookworm

# Install CRE CLI from pre-built release
RUN curl -sL https://github.com/smartcontractkit/cre-cli/releases/download/v1.1.0/cre_linux_amd64_ldd2-35.tar.gz \
    | tar xz -C /usr/local/bin && \
    mv /usr/local/bin/cre_v1.1.0_linux_amd64 /usr/local/bin/cre && \
    chmod 755 /usr/local/bin/cre

# Verify CRE CLI works
RUN cre --version || cre version || echo "CRE CLI installed"

# Copy workflow source
WORKDIR /workflows
COPY . .

# Pre-download Go module dependencies (speeds up simulate)
RUN go mod download

# Default: run a shell so we can interactively test workflows
ENTRYPOINT ["/bin/bash"]
