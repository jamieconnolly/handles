GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)

LDFLAGS ?= -X main.Name=$(NAME) -X main.Version=$(VERSION)
GOFLAGS ?= -ldflags "$(LDFLAGS)"

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p $(@D)
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GOFLAGS) -o "$@"

.PHONY: build
build: clean bin/$(NAME) completions/$(NAME).zsh

completions/$(NAME).%:
	@echo "==> Generating $@…"
	@mkdir -p $(@D)
	@cp $(dir $(realpath $(lastword $(MAKEFILE_LIST))))../completions/mondas.$* $@
	@sed -i "" "s/mondas/$(NAME)/g" $@
