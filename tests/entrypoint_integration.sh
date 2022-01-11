#!/usr/bin/env bash

set -euo pipefail

# Wait for the cache and db to be created
sleep 10

go test -v -p 1 -tags=integration -coverprofile=cover.out ./...
go tool cover -func=cover.out
