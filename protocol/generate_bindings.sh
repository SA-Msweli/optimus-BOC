#!/usr/bin/env bash
set -euo pipefail

# Generates Go bindings for Optimus smart contracts.
# Run from the repository root (e.g. `cd protocol && ./generate_bindings.sh`).

# Locate abigen, it may be in $PATH or in $(go env GOPATH)/bin
ABIGEN_BIN="$(command -v abigen 2>/dev/null || true)"
if [ -z "$ABIGEN_BIN" ]; then
    GOPATH=$(go env GOPATH)
    if [ -n "$GOPATH" ] && [ -x "$GOPATH/bin/abigen" ]; then
        ABIGEN_BIN="$GOPATH/bin/abigen"
    fi
fi
if [ -z "$ABIGEN_BIN" ]; then
    echo "error: abigen binary not found."
    echo "Install it with: go install github.com/ethereum/go-ethereum/cmd/abigen@latest" >&2
    exit 1
fi

OUT="bindings"
mkdir -p "$OUT"

# list of contracts we care about
for NAME in BNPLManager DAOManager LoanManager TokenVault DIDRegistry; do
    JSON="../contracts/out/${NAME}.sol/${NAME}.json"
    if [[ -f "$JSON" ]]; then
        echo "generating $NAME"
        # abigen expects raw ABI array; extract using a short Go snippet
        ABI_TMP=$(mktemp)
        BIN_TMP=$(mktemp)
        # extract ABI and bytecode via temporary Go program
        EXTRACT_SRC=$(mktemp --suffix=.go)
        cat > "$EXTRACT_SRC" <<'GOPROG'
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "usage: extract <file>")
        os.Exit(1)
    }
    data, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        panic(err)
    }
    var obj map[string]interface{}
    if err := json.Unmarshal(data, &obj); err != nil {
        panic(err)
    }
    abi, err := json.Marshal(obj["abi"])
    if err != nil {
        panic(err)
    }
    fmt.Println(string(abi))
    if rawBin, ok := obj["bytecode"]; ok {
        fmt.Fprintln(os.Stdout, "---BIN---")
        switch tb := rawBin.(type) {
        case string:
            fmt.Println(tb)
        case map[string]interface{}:
            if inner, ok := tb["object"].(string); ok {
                fmt.Println(inner)
            }
        default:
            // ignore
        }
    }
}
GOPROG
        # run and split output
        go run "$EXTRACT_SRC" "$JSON" | awk '/^---BIN---/ {flag=1; next} !flag {print > "'$ABI_TMP'"} flag {print > "'$BIN_TMP'"}'
        rm -f "$EXTRACT_SRC"
        # call abigen with both files if bin exists
        CMD=("$ABIGEN_BIN" --abi "$ABI_TMP" --pkg bindings --type "$NAME" --out "$OUT/${NAME}.go")
        if [ -s "$BIN_TMP" ]; then
            CMD+=(--bin "$BIN_TMP")
        fi
        "${CMD[@]}"
        rm -f "$ABI_TMP" "$BIN_TMP"
    else
        echo "warning: $JSON not found, skipping $NAME" >&2
    fi
done

# additional helpers could be generated here (events, enums, etc.)

echo "bindings generated in $OUT"