NAME := handles
VERSION := $(shell git describe --match 'v[0-9]*' --dirty --always | sed 's/^v//')

LDFLAGS := -X main.Name=$(NAME) -X main.Version=$(VERSION)

all: clean test build

include vendor/github.com/jamieconnolly/mondas/make/*.mk

check:
	@echo "==> Running static analysis…"
	@shellcheck -f gcc libexec/*

test: check

.PHONY: all check test
