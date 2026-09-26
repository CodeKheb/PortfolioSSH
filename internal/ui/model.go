package ui

/*
	Model, this is what gets rendered in cmd/main.go
	Here also lives the Init(), Update(), and View() for bubbletea
*/

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Screen type, iota MenuScreen so it's default at 0
type Screen int

const (
	MenuScreen Screen = iota
	AboutScreen
	ProjectScreen
	ContactScreen
	READMEScreen
)

// The Model struct, here is what the bubbletea func render
type Model struct {
	width    int
	height   int
	selected int
	screen   Screen
	scroll   int

	viewport viewport.Model
	search   textinput.Model

	searchMatches []int
	searchIndex   int
	searching     bool

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

		var header string

		switch model.screen {
		case MenuScreen:
			header = model.headerView("Kherbin's Portfolio")
		case AboutScreen:
			header = model.headerView("ABOUT  KHERBIN BUENAVENTURA")
		case ProjectScreen:
			header = model.headerView("PROJECTS")
		case READMEScreen:
			project := projectItems[model.selected]
			header = model.headerView("PROJECT: " + project.Title)
		}

		headerHeight := lipgloss.Height(header)
		footerHeight := lipgloss.Height(model.footerView())

		model.viewport.Width = msg.Width
		model.viewport.Height = max(
			0,
			msg.Height-headerHeight-footerHeight,
		)

		model.updateContent()

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
	case ProjectScreen:
		return model.ProjectView()
	case READMEScreen:
		return model.READMEView()
	case ContactScreen:
		return model.ContactView()
	default:
		return model.MenuView()
	}
}

// ViewportModel gets called in main.go
// Initializes everything
func ViewportModel() Model {
	vp := viewport.New(0, 0)

	input := textinput.New()
	input.Placeholder = "Search...  •  n/N to navigate"
	input.CharLimit = 100
	input.Prompt = ""

	return Model{
		viewport:   vp,
		search:     input,
		showFooter: true,
	}
}
