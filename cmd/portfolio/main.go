package main

import (
	"fmt"

	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// Main program
// starts the lipgloss UI
func main() {
	program := tea.NewProgram(
		ui.ViewportModel(),
		tea.WithAltScreen(),
	)
	if _, err := program.Run()

	err != nil {
		fmt.Print(err)
	}
}
