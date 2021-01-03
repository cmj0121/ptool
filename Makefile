SUBDIR := rssh
BIN := forensic

.PHONY: all clean help install $(SUBDIR)

all:		# build all

clean:		# clean-up the environment

help:		# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

build push:	$(SUBDIR)

$(SUBDIR):
	$(MAKE) -C $@ $(MAKECMDGOALS)

install:
	install -m755 $(BIN) /usr/local/bin/
