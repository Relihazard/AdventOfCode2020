GO		:= go
BUILD	:= build

SRC		:= \
	$(shell find . -type f -name '*.go' -print)
BIN		:= \
	aoc_2020

GOBIN	= $(shell go env GOBIN)/
ifeq ($(GOBIN),)
GOBIN	= $(shell go env GOPATH)/bin
endif

QUIET	:= @
ECHO	:= echo
RM		:= rm
RMFLAGS	:= -f

.PHONY: all clean
all: $(GOBIN)$(BIN)

$(GOBIN)$(BIN): $(SRC)
	$(QUIET)$(ECHO) "[$(BIN)] go build"
	$(QUIET)$(GO) $(BUILD) -o $(GOBIN)$(BIN) $(SRC)

clean:
	$(RM) $(RMFLAGS) $(GOBIN)$(BIN)