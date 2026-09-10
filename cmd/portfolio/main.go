package main

import (
	"fmt"

	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width int
	height int
}

func main() {
	program := tea.NewProgram(ui.Model{})
	if _, err := program.Run()

	err != nil {
		fmt.Print(err)
	}
}





