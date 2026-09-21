package ui

import "github.com/charmbracelet/lipgloss"

// overall layout, gets called in different view
func (model Model) layout(content string) string {
	footer := model.footerView()

	footerHeight := lipgloss.Height(footer)

	contentHeight := model.height - footerHeight

	body := lipgloss.NewStyle().
		Width(model.width).
		Height(contentHeight).
		Padding(1, 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		body,
		footer,
	)
}

// layout with header
func (model Model) layoutWithHeader(header, content string) string {
	footer := model.footerView()

	footerHeight := lipgloss.Height(footer)
	headerHeight := lipgloss.Height(header)

	contentHeight := model.height - headerHeight - footerHeight

	content = lipgloss.NewStyle().
		Width(model.width).
		Height(contentHeight).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		content,
		footer,
	)
}

// footer
func (model Model) footerView() string {
	helper := ""

	if model.searching {
		return instructionsStyle.Render(
			keyStyle.Render("> ") + model.search.View(),
		)

	}

	if model.showFooter {
		helper =
			labelStyle.Render("Navigate") + " " +
				keyStyle.Render("↑/↓ j/k") + "    " +
				labelStyle.Render("Scroll") + " " +
				keyStyle.Render("Ctrl+D/U") + "    " +
				labelStyle.Render("Select") + " " +
				keyStyle.Render("Enter") + "    " +
				labelStyle.Render("Back") + " " +
				keyStyle.Render("b/Esc") + "    " +
				labelStyle.Render("Quit") + " " +
				keyStyle.Render("q") +  "    " +
				labelStyle.Render("Search") + " " +
				keyStyle.Render("/ ") + "	"
			}

	return instructionsStyle.Render(
		helper +
			labelStyle.Render("? Help"),
	)
}

// headerView that takes in a string 
// adjusts to the width of the users terminal
// returns with headerStyle
func (model Model) headerView(header string) string {
	return headerStyle.
		Width(model.width).
		Render(header)
}
