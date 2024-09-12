EXE = baseconv

all: tui cli

download_pkgs:
	go mod download

tui: download_pkgs
	go build ./cmd/$(EXE)

cli: download_pkgs
	go build ./cmd/$(EXE)-cli

clean:
	rm $(EXE)
	rm $(EXE)-cli

install: all
	sudo mv ./$(EXE) /usr/bin/$(EXE)
	sudo mv ./$(EXE)-cli /usr/bin/$(EXE)-cli

uninstall:
	sudo rm ./$(EXE) /usr/bin/$(EXE)
	sudo rm ./$(EXE)-cli /usr/bin/$(EXE)-cli

.PHONY: all download_pkgs tui cli clean install uninstall
