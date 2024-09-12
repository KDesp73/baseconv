package main

import (
	"baseconv/cmd/baseconv/tui"
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	debug := false
	flag.BoolVar(&debug, "debug", debug, "Enable debug logging")
	flag.Parse()

	if debug {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatalf("ERRO: %v", err)
		}
		defer f.Close()
	}


	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
