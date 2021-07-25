.DEFAULT_GOAL	:= install
.PHONY: build install

build:
	go build ./...

install:
	go install ./...