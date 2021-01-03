PREFIX := /usr/local/bin

SRC := $(wildcard *.go)
BIN := $(subst .go,,$(SRC))

.PHONY: all clean help install
all: $(BIN)	# build all

clean:		# clean-up the environment
	rm -f $(BIN)

help:		# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

install: $(BIN)
	install -m755 $(BIN) $(PREFIX)/$(BIN)

$(BIN): linter

GO      := go
GOFMT   := $(GO)fmt -w -s
GOFLAG  := -ldflags="-s -w"
GOTEST  := $(GO) test -cover -failfast -timeout 2s
GOBENCH := $(GO) test -bench=. -cover -failfast -benchmem

linter: .benchmark
	$(GOFMT) $(shell find . -name '*.go')
	$(GOTEST) ./...

.benchmark:
	$(GOBENCH)
	touch $@

%: %.go
	$(GO) build $(GOFLAG) -o $@ $<
