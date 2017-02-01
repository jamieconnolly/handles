GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)

LDFLAGS ?= -X main.Name=$(NAME) -X main.Version=$(VERSION)
GOFLAGS ?= -ldflags "$(LDFLAGS)"

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p bin
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GOFLAGS) -o "$@"

.PHONY: build
build: clean bin/$(NAME)
