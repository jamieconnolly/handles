DESTDIR ?= /usr/local
PREFIX ?= $(DESTDIR)/opt/$(NAME)

.PHONY: uninstall
uninstall:
	@echo "==> Uninstalling from $(DESTDIR)…"
	@rm -frv $(PREFIX)
	@rm -fv $(DESTDIR)/bin/$(NAME)
