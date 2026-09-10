package main

import (
	"fmt"

	"github.com/CodeKheb/PortfolioSSH/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width int
	height int
}

func main() {
	program := tea.NewProgram(Model{})
	if _, err := program.Run()

	err != nil {
		fmt.Print(err)
	}

	fmt.Print("Hi lol, this my portfolio")
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	main := ui.MainStyle.
		Width(model.width - 4).
		Height(model.height - 4).
		Align(lipgloss.Center, lipgloss.Center).
		Render("Hi lol")
	return main
}



