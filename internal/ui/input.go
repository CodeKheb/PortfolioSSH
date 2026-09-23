package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (model Model) keyHandler(message tea.KeyMsg) (Model, tea.Cmd) {
	if model.screen == AboutScreen {
		var cmd tea.Cmd

		if model.searching {
			switch message.String() {
			case "esc", "enter":
				model.searching = false
				model.search.Blur()
				model.search.Reset()

			default:
				model.search, cmd = model.search.Update(message)
				model.viewport.SetContent(model.aboutContent())

				model.searchContent()
				return model, cmd
			}

			return model, nil
		}

		switch message.String() {
		case "esc", "b":
			model.screen = MenuScreen

		case "q":
			return model, tea.Quit
		case "j", "down":
			model.viewport.ScrollDown(1)
		case "ctrl+d":
			model.viewport.HalfPageDown()

		case "k", "up":
			model.viewport.ScrollUp(1)
		case "ctrl+u":
			model.viewport.HalfPageUp()
		case "?":
			model.showFooter = !model.showFooter
		case "/":
			model.searching = true
			model.search.Focus()
			return model, textinput.Blink
		case "n":
			model.nextSearchMatch()
			return model, nil
		case "N":
			model.previousSearchMatch()
			return model, nil
		}

		return model, nil
	}

	if model.screen == READMEScreen {
		var cmd tea.Cmd

		if model.searching {
			switch message.String() {
			case "esc", "enter":
				model.searching = false
				model.search.Blur()

			default:
				model.search, cmd = model.search.Update(message)

				model.viewport.SetContent(model.readmeContent())
				model.searchContent()
				return model, cmd
			}

			return model, nil
		}
		switch message.String() {
		case "esc", "b":
			model.screen = ProjectScreen

		case "q":
			return model, tea.Quit
		case "j", "down":
			model.viewport.ScrollDown(1)
		case "ctrl+d":
			model.viewport.HalfPageDown()

		case "k", "up":
			model.viewport.ScrollUp(1)
		case "ctrl+u":
			model.viewport.HalfPageUp()
		case "?":
			model.showFooter = !model.showFooter
		case "/":
			model.searching = true
			model.search.Reset()
			model.search.Focus()
			return model, textinput.Blink
		case "n":
			model.nextSearchMatch()
			return model, nil
		case "N":
			model.previousSearchMatch()
			return model, nil
		}

		return model, nil
	}

	if model.screen == ProjectScreen {
		var cmd tea.Cmd

		if model.searching {
			switch message.String() {
			case "esc", "enter":
				model.searching = false
				model.search.Blur()

			default:
				model.search, cmd = model.search.Update(message)
				model.viewport.SetContent(model.ProjectView())

				model.searchProjects()
				return model, cmd
			}

			return model, nil
		}

		switch message.String() {
		case "k", "up":
			if model.selected > 0 {
				model.selected--
				model.updateContent()
			}
		case "j", "down":
			if model.selected < len(projectItems)-1 {
				model.selected++
				model.updateContent()
			}
		case "esc", "b":
			model.screen = MenuScreen
			model.selected = 0
			return model, nil
		case "q":
			return model, tea.Quit

		case "/":
			model.searching = true
			model.search.Reset()
			model.search.Focus()
			return model, textinput.Blink
		case "?":
			model.showFooter = !model.showFooter
		case "n":
			model.nextSearchMatch()
			return model, nil
		case "N":
			model.previousSearchMatch()
			return model, nil
		case "enter":
			model.screen = READMEScreen

			model.searching = false
			model.search.Reset()
			model.search.Blur()

			model.viewport.SetContent(model.readmeContent())
			model.viewport.GotoTop()
			return model, nil
		}

		return model, nil
	}

	switch message.String() {
	case "q":
		return model, tea.Quit

	case "k", "up":
		if model.selected > 0 {
			model.selected--
		}
	case "j", "down":
		if model.selected < len(menuItems)-1 {
			model.selected++
		}
	case "enter":
		switch model.selected {
		case 0:
			model.selected = 0
			model.screen = AboutScreen
			model.viewport.SetContent(model.aboutContent())
			model.viewport.GotoTop()
		case 1:
			model.selected = 0
			model.screen = ProjectScreen
			model.updateContent()
			model.viewport.GotoTop()
		}
	case "?":
		model.showFooter = !model.showFooter
	}
	return model, nil
}
