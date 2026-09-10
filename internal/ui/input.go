package ui

import tea "github.com/charmbracelet/bubbletea"

func (model Model) keyHandler(message tea.KeyMsg) (Model, tea.Cmd) {
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
			if model.selected < len(menuItems) - 1 {
				model.selected++
			}
		}
	}
	return model, nil
}
