package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var menuItems = []string {
	"About",
	"Projects",
	"Skills",
	"Contact",
}

func (model Model) MainView() string {
	var view string

	for i, items := range menuItems {
		cursor := " "

		if i == model.selected {
			cursor = ">"
		}

		view += cursor + items + "\n"
	}

	main := MainStyle.
		Width(model.width - 2).
		Height(model.height - 3).
		Align(lipgloss.Center, lipgloss.Center).
		Render(fmt.Sprint(
			"Hi lol\n\n",
		) + view)
	return main
}
