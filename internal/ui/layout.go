package ui

import "github.com/charmbracelet/lipgloss"

// overall layout, gets called in different view
func (model Model) layout(content string) string {
    footer := model.footerView()

    contentHeight := model.height - lipgloss.Height(footer)

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

// footer
func (model Model) footerView() string {
    return instructionsStyle.Render(
        "↑/↓ Navigate • Enter Select • q Quit\n\n" +
            "Supports vim navigation  j/k",
    )
}
