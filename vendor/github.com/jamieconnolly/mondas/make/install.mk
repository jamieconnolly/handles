DESTDIR ?= /usr/local
PREFIX ?= $(DESTDIR)/opt/$(NAME)
BINDIR ?= $(PREFIX)/bin
LIBEXECDIR ?= $(PREFIX)/libexec

.PHONY: install
install: build
	@echo "==> Installing to $(DESTDIR)…"
	@install -d $(BINDIR) $(LIBEXECDIR)
	@install -m 0755 -pv bin/$(NAME) $(BINDIR)
	@install -m 0755 -pv libexec/* $(LIBEXECDIR)
	@ln -fsv $(BINDIR)/$(NAME) $(DESTDIR)/bin/$(NAME)
