.PHONY: all clean help install restart

GO  = ~/go/bin/go1.16rc1
BIN =$(subst .go,,$(wildcard cmd/*.go))

all: $(BIN)

clean:	# cleanup environment
	rm -f $(BIN)

help:  # show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

PREFIX := /usr/local/bin
install: $(BIN)	# install to the system
	install -m755 $(BIN) $(PREFIX)/shell-saas
	install -m644 init/shell-saas.systemd /etc/systemd/system/shell-saas.service
	install -m644 init/shell-saas.nginx   /etc/nginx/site-enabled/shell-sass.nginx

restart:	# restart the SaaS service
	systemctl daemon-reload
	systemctl restart shell-saas.service
	nginx -t
	systemctl restart nginx

%: %.go
	$(GO) build -o $@ $<
