package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// lipgloss style variables
var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#1A1B26")).
			Background(lipgloss.Color("#7AA2F7")).
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 4)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Bold(true).
				Foreground(lipgloss.Color("#7D56F4"))

	descriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#777777")).
				PaddingLeft(4)

	instructionsStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#565F89")).
				Padding(2, 4)
)
