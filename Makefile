all: tui cli

tui:
	go build ./cmd/baseconv

cli:
	go build ./cmd/baseconv-cli

.PHONY: all tui cli
