NAME := handles
VERSION := $(shell git describe --match 'v[0-9]*' --dirty --always | sed 's/^v//')

DESTDIR ?= /usr/local
PREFIX ?= $(DESTDIR)/opt/$(NAME)
BINDIR ?= $(PREFIX)/bin
LIBEXECDIR ?= $(PREFIX)/libexec

GOARCH ?= $(shell go env GOARCH)
GOOS ?= $(shell go env GOOS)

LDFLAGS := -X github.com/jamieconnolly/mondas.Version=$(VERSION)
GOFLAGS := -ldflags "$(LDFLAGS)"

all: clean test build

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p bin
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GOFLAGS) -o "$@"

build: clean bin/$(NAME)

check: lint

clean:
	@echo "==> Cleaning generated files…"
	@rm -rf bin/*

install: build
	@echo "==> Installing to $(DESTDIR)…"
	@install -d $(BINDIR) $(LIBEXECDIR)
	@install -m 0755 -pv bin/$(NAME) $(BINDIR)
	@install -m 0755 -pv libexec/* $(LIBEXECDIR)
	@rm -f $(DESTDIR)/bin/$(NAME)
	@ln -sv $(BINDIR)/handles $(DESTDIR)/bin/$(NAME)

lint:
	@echo "==> Running static analysis…"
	@shellcheck -f gcc libexec/*

test: check

uninstall:
	@echo "==> Uninstalling from $(DESTDIR)…"
	@rm -frv $(PREFIX)
	@rm -fv $(DESTDIR)/bin/$(NAME)

.PHONY: all build check clean install lint test uninstall
