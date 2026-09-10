package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width int
	height int
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch m := message.(type) {
		case tea.KeyMsg:
			if m.String() == "q" {
				return model, tea.Quit
			}
		}
	return model, nil
}

func (model Model) View() string {
	return model.MainView()
}

