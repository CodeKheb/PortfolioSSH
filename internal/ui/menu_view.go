package ui

import (
	"strings"
)

type MenuItems struct {
	Title       string
	Description string
}

var menuItems = []MenuItems{
	{
		Title:       "About",
		Description: "Hi lol",
	},
	{
		Title:       "Projects",
		Description: "What I larp about",
	},
	{
		Title:       "Skills",
		Description: "A bit of this a bit of that",
	},
	{
		Title:       "Contact",
		Description: "CodeKheb",
	},
}

// TODO: Make menuLines
// func (model Model) menuLines() []ContentLine

// The first ui the user sees, contains the navigation
func (model Model) MenuView() string {
	var builder strings.Builder

	builder.WriteString(titleStyle.Render("Kherbin's Portfolio"))
	builder.WriteString("\n\n\n")

	for i, items := range menuItems {

		if i == model.selected {
			builder.WriteString(selectedItemStyle.Render("> " + items.Title))
			builder.WriteString(
				descriptionStyle.Render("\n" + menuItems[i].Description),
			)
		} else {
			builder.WriteString(itemStyle.Render("  " + items.Title))
		}

		builder.WriteString("\n\n\n")
	}

	builder.WriteString("\n\n")

	return model.layout(builder.String())
}
