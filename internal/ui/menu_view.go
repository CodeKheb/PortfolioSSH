package ui

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

func (model Model) menuLines() []ContentLine {
	lines := []ContentLine{
		{
		},
	}

	lines = append(lines, ContentLine{
		Text: "\n",
	})

	for i, menu := range menuItems {
		if i == model.selected {
			lines = append(
				lines,
				ContentLine{
					Text:  "> " + menu.Title,
					Style: selectedItemStyle,
				},
				ContentLine{
					Text:  menu.Description,
					Style: descriptionStyle,
				},
			)
		} else {
			lines = append(
				lines,
				ContentLine{
					Text:  "  " + menu.Title,
					Style: itemStyle,
				},
			)
		}

		lines = append(lines, ContentLine{
			Text: "\n",
		})

		lines = append(lines, ContentLine{})
	}

	return lines
}

// The first ui the user sees, contains the navigation
func (model Model) MenuView() string {
	header := model.headerView("Kherbin's Portfolio")

	return model.layout(
		header,
		model.renderLines(model.menuLines()),
	)
}
