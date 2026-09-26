package ui

/*
	TODO:
	Initial Plan, users can send a message, store in db, send to my email?
	also add clickable copyable/escape links
*/

type Contact struct {
	Name     string
	Email    string
	LinkedIn string
	GitHub   string
}

var contactItems = []Contact {
	{
		Name: "Kherbin Clloyde Buenaventura",
		Email: "kherbinbuenaventura@gmail.com",
		LinkedIn: "https://www.linkedin.com/in/kherbin-clloyde-buenaventura-875414401/",
		GitHub: "https://github.com/CodeKheb",
	},
}

func (model Model) contactLines() []ContentLine {
	lines := []ContentLine{}

	lines = append(lines, ContentLine{
		Text: contactItems[0].Name,
		Style: selectedItemStyle,
	})
	lines = append(lines, ContentLine{
		Text: contactItems[0].Email,
		Style: selectedItemStyle,
	})
	lines = append(lines, ContentLine{
		Text: contactItems[0].LinkedIn,
		Style: selectedItemStyle,
	})
	lines = append(lines, ContentLine{
		Text: contactItems[0].GitHub,
		Style: selectedItemStyle,
	})

	return lines

}

func (model Model) contactContent() string {
	return  model.renderLines(model.contactLines())
}

func (model Model) ContactView() string {
	header := model.headerView("CONTACT")

	return model.layout(
		header,
		model.viewport.View(),
	)
}
