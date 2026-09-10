package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	MainStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#666666")).
		Padding(10, 10).
		Bold(true)
)
