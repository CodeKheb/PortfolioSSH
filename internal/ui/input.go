package ui

import tea "github.com/charmbracelet/bubbletea"

func (model Model) keyHandler(message tea.KeyMsg) (Model, tea.Cmd) {
	if model.screen == AboutScreen {
		switch message.String() {
		case "esc":
			model.screen = MenuScreen

		case "q":
			return model, tea.Quit

		case "j", "down":
			model.viewport.ScrollDown(1)

		case "k", "up":
			model.viewport.ScrollUp(1)
		}

		return model, nil
	}

	switch message.String() {
	case "q":
		return model, tea.Quit

	case "k", "up":
		{
			if model.selected > 0 {
				model.selected--
			}
		}
	case "j", "down":
		{
			if model.selected < len(menuItems)-1 {
				model.selected++
			}
		}
	case "enter":
		switch model.selected {
		case 0:
			model.screen = AboutScreen
			model.viewport.SetContent(aboutContent())
			model.viewport.GotoTop()
		}
	}
	return model, nil
}
