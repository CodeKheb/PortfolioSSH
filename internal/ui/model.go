package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width  int
	height int
	selected int
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height
		return model, nil

	case tea.KeyMsg:
		return model.keyHandler(msg)
	}

	return model, nil
}

func (model Model) View() string {
	return model.MainView()
}
