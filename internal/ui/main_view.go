package ui

import "github.com/charmbracelet/lipgloss"


func (model Model) MainView() string {
	main := MainStyle.
		Width(model.width - 4).
		Height(model.height - 4).
		Align(lipgloss.Center, lipgloss.Center).
		Render("Hi lol")
	return main
}
