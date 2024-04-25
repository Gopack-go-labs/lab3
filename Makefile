LANG=en_US.UTF-8
SHELL=/bin/bash
.SHELLFLAGS=--norc --noprofile -e -u -o pipefail -c

run:
	go run ./cmd/painter

test:
	go test ./...
