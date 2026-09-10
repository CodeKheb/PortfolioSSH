package ui

import (
	"strings"
)

type MenuItems struct {
	Title string
	Description string
}

var menuItems = []MenuItems {
	{
		Title: "About",
		Description: "Hi lol",
	},
	{
		Title: "Projects",
		Description: "What I larp about",
	},
	{
		Title: "Skills",
		Description: "A bit of this a bit of that",
	},
	{
		Title: "Contact",
		Description: "CodeKheb",
	},
}

func (model Model) MainView() string {
	var builder strings.Builder

	builder.WriteString(titleStyle.Render("Kherbin's Portfolio"))
	builder.WriteString("\n\n")

	for i, items := range menuItems {

		if i == model.selected {
			builder.WriteString(selectedItemStyle.Render(">" + items.Title))
		} else {
			builder.WriteString(selectedItemStyle.Render(" " + items.Title))
		}

		builder.WriteString("\n")

		builder.WriteString(
			descriptionStyle.Render(items.Description),
		)

		builder.WriteString("\n\n")
	}

	return appStyle.Render(builder.String())
}
