package main

import (
	"fmt"

	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(ui.Model{}, tea.WithAltScreen())
	if _, err := program.Run()

	err != nil {
		fmt.Print(err)
	}
}
