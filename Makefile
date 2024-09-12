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
	sudo cp ./$(EXE) /usr/bin/$(EXE)
	sudo cp ./$(EXE)-cli /usr/bin/$(EXE)-cli

uninstall:
	sudo rm /usr/bin/$(EXE)
	sudo rm /usr/bin/$(EXE)-cli

.PHONY: all download_pkgs tui cli clean install uninstall
