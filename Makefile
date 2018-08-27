NAME := $(shell basename `pwd`)
VERSION := $(shell git describe --match 'v[0-9]*' --dirty --always | sed 's/^v//')

PREFIX := /usr/local
DESTDIR := $(PREFIX)/opt/$(NAME)

GOARCH := $(shell go env GOARCH)
GOFLAGS := -ldflags "-X main.Name=$(NAME) -X main.Version=$(VERSION)"
GOOS := $(shell go env GOOS)

all: test build

bin/$(NAME):
	@echo "==> Building $@…"
	@mkdir -p $(@D)
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build $(GOFLAGS) -o "$@"

build: clean bin/$(NAME)

clean:
	@echo "==> Cleaning generated files…"
	@rm -f bin/*

deps:
	@echo "==> Installing dependencies…"
	@go get
	@go mod tidy

install: build
	@echo "==> Installing to $(PREFIX)…"
	@install -d $(DESTDIR)/bin $(DESTDIR)/completions $(DESTDIR)/libexec
	@install -m 0755 -pv bin/$(NAME) $(DESTDIR)/bin
	@install -m 0755 -pv completions/* $(DESTDIR)/completions
	@install -m 0755 -pv libexec/* $(DESTDIR)/libexec
	@ln -fsv $(DESTDIR)/bin/$(NAME) $(PREFIX)/bin/$(NAME)
	@ln -fsv $(DESTDIR)/completions/$(NAME).zsh $(PREFIX)/share/zsh/site-functions/_$(NAME)

lint:
	@echo "==> Running static analysis…"
	@shellcheck -f gcc libexec/*

test: lint

uninstall:
	@echo "==> Uninstalling from $(PREFIX)…"
	@rm -frv $(DESTDIR)
	@rm -fv $(PREFIX)/bin/$(NAME)
	@rm -fv $(PREFIX)/share/zsh/site-functions/_$(NAME)

.PHONY: all build clean deps install lint test uninstall
