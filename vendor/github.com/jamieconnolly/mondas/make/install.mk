PREFIX ?= /usr/local
DESTDIR ?= $(PREFIX)/opt/$(NAME)

.PHONY: install
install: build
	@echo "==> Installing to $(PREFIX)…"
	@install -d $(DESTDIR)/bin $(DESTDIR)/completions $(DESTDIR)/libexec
	@install -m 0755 -pv bin/$(NAME) $(DESTDIR)/bin
	@install -m 0755 -pv completions/* $(DESTDIR)/completions
	@install -m 0755 -pv libexec/* $(DESTDIR)/libexec
	@ln -fsv $(DESTDIR)/bin/$(NAME) $(PREFIX)/bin/$(NAME)
	@ln -fsv $(DESTDIR)/completions/$(NAME).zsh $(PREFIX)/share/zsh/site-functions/_$(NAME)
