package ui

import tea "github.com/charmbracelet/bubbletea"

func (model Model) keyHandler(message tea.KeyMsg) (Model, tea.Cmd) {
	switch message.String() {
		case "q":
			return model, tea.Quit
		}
		return model, nil
}
