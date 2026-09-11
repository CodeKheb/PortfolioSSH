package ui

/*
	Model, this is what gets rendered in cmd/main.go
	Here also lives the Init(), Update(), and View() for bubbletea
*/

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Screen type, iota MenuScreen so it's default at 0
type Screen int

const (
	MenuScreen Screen = iota
	AboutScreen
)

// The Model struct, here is what the bubbletea func render
type Model struct {
	width    int
	height   int
	selected int
	screen   Screen
}

// Initialize tea
func (model Model) Init() tea.Cmd {
	return nil
}

// Update
/*
	Here we get the size of the users terminal 
	and adjust the sizes accordingly
	keyHandler() gets called
 */
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

// View(), switch if user is in another screen
func (model Model) View() string {
	switch model.screen {
	case AboutScreen:
		return model.AboutView()
	default:
		return model.MenuView()
	}
}
