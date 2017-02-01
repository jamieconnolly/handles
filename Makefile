NAME := handles
VERSION := $(shell git describe --match 'v[0-9]*' --dirty --always | sed 's/^v//')

.PHONY: all
all: clean test build

include vendor/github.com/jamieconnolly/mondas/make/*.mk

.PHONY: check
check:
	@echo "==> Running static analysis…"
	@shellcheck -f gcc libexec/*

.PHONY: test
test: check
