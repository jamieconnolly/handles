PREFIX ?= /usr/local
DESTDIR ?= $(PREFIX)/opt/$(NAME)

.PHONY: uninstall
uninstall:
	@echo "==> Uninstalling from $(PREFIX)…"
	@rm -frv $(DESTDIR)
	@rm -fv $(PREFIX)/bin/$(NAME)
	@rm -fv $(PREFIX)/share/zsh/site-functions/_$(NAME)
