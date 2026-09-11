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
			Background(lipgloss.Color("#7D56F4")).
			Padding(1, 2)

	// About page
	nameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#2D7DFF")).
			MarginBottom(1)

	sectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#BB9AF7")).
			MarginTop(1).
			MarginBottom(1)

	subheadingStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#73DACA"))

	bulletStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F7768E"))

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9B1D6"))

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	// Menu Style
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

	// Search
	searchHighlightStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#2D7DFF")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)

	// Footer Style
	keyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#2D7DFF"))

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9B1D6"))
)
