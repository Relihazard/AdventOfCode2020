2020_DIR		:= 2020/
2020_BIN		:= aoc_2020
2020_SRC		:= \
	$(shell find . -type f -name '*.go' -print)

GOBIN	= $(shell go env GOBIN)/
ifeq ($(GOBIN),)
GOBIN	= $(shell go env GOPATH)/bin
endif

QUIET	:= @
MAKE	:= make

.PHONY: all
all: $(GOBIN)$(2020_BIN)

$(GOBIN)$(2020_BIN): $(2020_SRC)
	$(QUIET)$(MAKE) -C $(2020_DIR)

