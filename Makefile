NAME = handles

DESTDIR ?= /usr/local/opt/$(NAME)
BINDIR ?= $(DESTDIR)/bin
LIBEXECDIR ?= $(DESTDIR)/libexec

all: build

bin/$(NAME):
	@mkdir -p $(dir $@)
	go build -o $@ main.go

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
	@rm -frv $(DESTDIR)
	@rm -fv /usr/local/bin/$(NAME)

.PHONY: all build clean install uninstall
