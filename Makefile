all: tui cli

tui:
	go build ./cmd/baseconv

cli:
	go build ./cmd/baseconv-cli

clean:
	rm baseconv
	rm baseconv-cli

.PHONY: all tui cli clean
