GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)

LDFLAGS ?= -X main.Name=$(NAME) -X main.Version=$(VERSION)
GOFLAGS ?= -ldflags "$(LDFLAGS)"

SELF_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p $(@D)
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GOFLAGS) -o "$@"

.PHONY: build
build: clean bin/$(NAME) completions/$(NAME).zsh

completions/$(NAME).%:
	@echo "==> Generating $@…"
	@mkdir -p $(@D)
	@cp $(SELF_DIR)../completions/mondas.$* "$@"
	@sed -e "s/mondas/$(NAME)/g" "$@" | tee "$@" >/dev/null
