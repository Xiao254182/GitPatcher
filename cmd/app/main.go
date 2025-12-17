package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"GitPatcher/internal/ui"
)

func main() {
	p := tea.NewProgram(
		ui.NewModel(),
		tea.WithAltScreen(),
	)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
