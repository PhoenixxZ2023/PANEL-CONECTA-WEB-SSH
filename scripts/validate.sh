#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
command -v go >/dev/null 2>&1 || { echo "ERROR: Go is required"; exit 1; }
go test ./...
go build ./cmd/panel
echo "Validation OK"
