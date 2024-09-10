all: tui cli

tui:
	go build ./cmd/converter

cli:
	go build ./cmd/converter-cli

.PHONY: all tui cli
