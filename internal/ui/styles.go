package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// colors
var (
	BLUE     = lipgloss.Color("#2D7DFF")
	PURPLE   = lipgloss.Color("#7D56F4")
	GREEN    = lipgloss.Color("#73DACA")
	RED      = lipgloss.Color("#F7768E")
	WHITE    = lipgloss.Color("#FFFFFF")
	CREAM    = lipgloss.Color("#FFFDF5")
	TEXT     = lipgloss.Color("#C0CAF5")
	MUTED    = lipgloss.Color("#A9B1D6")
	DIM      = lipgloss.Color("#565F89")
	DARK     = lipgloss.Color("#1A1B26")
	LAVENDER = lipgloss.Color("#BB9AF7")
)

// lipgloss style variables
var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(DARK).
			Background(PURPLE).
			Padding(1, 2)

	// About page
	nameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BLUE).
			MarginBottom(1)

	sectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(LAVENDER).
			MarginTop(1).
			MarginBottom(1)

	subheadingStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(GREEN)

	bulletStyle = lipgloss.NewStyle().
			Foreground(RED)

	codeStyle = lipgloss.NewStyle().
			Foreground(MUTED).
			PaddingLeft(2)

	mutedStyle = lipgloss.NewStyle().
			Foreground(MUTED)

	bodyStyle = lipgloss.NewStyle().
			Foreground(WHITE)

	// Menu Style
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(CREAM).
			Background(PURPLE).
			Padding(0, 4)

	itemStyle = lipgloss.NewStyle().
			Foreground(WHITE).
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Bold(true).
				Foreground(PURPLE)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(TEXT).
				PaddingLeft(4)

	instructionsStyle = lipgloss.NewStyle().
				Foreground(DIM).
				Padding(2, 4)

	// Search
	searchHighlightStyle = lipgloss.NewStyle().
				Background(BLUE).
				Foreground(WHITE).
				Bold(true)

	// Footer Style
	keyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BLUE)

	labelStyle = lipgloss.NewStyle().
			Foreground(MUTED)
)
