package ui

/*
	Model, this is what gets rendered in cmd/main.go
	Here also lives the Init(), Update(), and View() for bubbletea
*/

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	scroll   int

	viewport viewport.Model

	showFooter bool
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
		headerHeight := lipgloss.Height(model.headerView(" "))
		footerHeight := lipgloss.Height(model.footerView())

		model.viewport.Width = msg.Width
		model.viewport.Height = msg.Height - headerHeight - footerHeight

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

func ViewportModel() Model {
	vp := viewport.New(0, 0)

	return Model{
		viewport:   vp,
		showFooter: true,
	}
}

