#!/usr/bin/env bash
set -euo pipefail

# Enforce comment policy for Solidity sources/tests in the Foundry project.
# - Disallow single-line `//` comments inside `contracts/src` and `contracts/test` (except SPDX)
# - Disallow `///` docstrings in test files (tests must be self-describing)
# - Disallow trivial `/// @notice` tags in tests

# ensure script runs from the contracts/ project root
cd "$(dirname "$0")/.."
ROOT_DIR="$PWD"
FAIL=0

echo "Checking for disallowed inline '//' comments (excluding SPDX and third-party)..."
# Find `//` that are not SPDX and not `///` (i.e. normal inline comments)
INLINE=$(grep -R --line-number -n --exclude-dir=lib -E "^[[:space:]]*//[^/]" src test || true)
# Filter out SPDX lines
INLINE=$(echo "$INLINE" | grep -v "// SPDX" || true)
if [ -n "$INLINE" ]; then
  echo "Found inline '//' comments (not allowed):"
  echo "$INLINE"
  FAIL=1
fi

# Disallow triple-slash docstrings in tests
TSLASH=$(grep -R --line-number -n -E "^[[:space:]]*///" test || true)
if [ -n "$TSLASH" ]; then
  echo "Found triple-slash (///) comments in test files — remove them or convert to clearer test names:" 
  echo "$TSLASH"
  FAIL=1
fi

# Disallow trivial @notice entries in tests
NOTICE=$(grep -R --line-number -n -E "^[[:space:]]*///[[:space:]]*@notice" test || true)
if [ -n "$NOTICE" ]; then
  echo "Found @notice in tests (trivial NatSpec) — remove or justify:" 
  echo "$NOTICE"
  FAIL=1
fi

if [ "$FAIL" -ne 0 ]; then
  echo "\nComment policy checks failed. Remove inline comments and trivial NatSpec from tests.\n"
  exit 1
fi

echo "Comment policy checks passed."