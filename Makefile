NAME = handles
VERSION = $(shell git describe --tags 2>/dev/null | cut -d - -f 1,2 | sed 's/^v//' | tr - .)

PREFIX ?= /usr/local/opt/$(NAME)
BINDIR ?= $(PREFIX)/bin
LIBEXECDIR ?= $(PREFIX)/libexec

all: build

LDFLAGS = -ldflags "-X=main.Version=$(VERSION)"

bin/$(NAME):
	@mkdir -p $(dir $@)
	go build -o $@ $(LDFLAGS) main.go

build: clean bin/$(NAME)

clean:
	@rm -f bin/$(NAME)

install: build
	@install -d $(BINDIR) $(LIBEXECDIR)
	@install -m 0755 -pv bin/$(NAME) $(BINDIR)
	@install -m 0755 -pv libexec/* $(LIBEXECDIR)
	@rm -f /usr/local/bin/$(NAME)
	@ln -sv $(BINDIR)/$(NAME) /usr/local/bin/$(NAME)

uninstall:
	@rm -frv $(PREFIX)
	@rm -fv /usr/local/bin/$(NAME)

.PHONY: all build clean install uninstall
