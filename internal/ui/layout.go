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

func (model Model) layoutWithHeader(header, content string) string {
	footer := model.footerView()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		content,
		footer,
	)
}

// footer
func (model Model) footerView() string {
    return instructionsStyle.Render(
        "↑/↓ Navigate • Enter Select • q Quit\n\n" +
            "Supports vim navigation  j/k • b/Esc Back",
    )
}

func (model Model) headerView(header string) string {
	return headerStyle.
		Width(model.width).
		Render(header)
}
