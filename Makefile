SHELL := /bin/bash
BASEDIR = $(shell pwd)

all: fmt mod lint test
fmt:
	gofmt -w .
mod:
	go mod tidy
lint:
	golangci-lint run
.PHONY: test
test: mod
	sh scripts/test.sh
update:
	# https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
	go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null	
redis:
	sh scripts/redis.sh
help:
	@echo "fmt - format the source code"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "update - list updateable packages"
	@echo "redis - launch a docker redis"