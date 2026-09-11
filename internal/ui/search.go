package ui

import "strings"

func (model Model) aboutSearchText() string {
	lines := model.aboutLines()

	var builder strings.Builder

	for i, line := range lines {
		builder.WriteString(line.Text)

		if i < len(lines)-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}


func highlightSearch(text string, query string) string {
	query = strings.TrimSpace(query)

	if query == "" {
		return text
	}

	lowerText := strings.ToLower(text)
	lowerQuery := strings.ToLower(query)

	var builder strings.Builder
	start := 0

	for {
		index := strings.Index(lowerText[start:], lowerQuery)

		if index == -1 {
			builder.WriteString(text[start:])
			break
		}

		index += start

		builder.WriteString(text[start:index])
		builder.WriteString(
			searchHighlightStyle.Render(
				text[index : index+len(query)],
			),
		)

		start = index + len(query)
	}

	return builder.String()
}

func (model *Model) searchContent() {
	query := strings.TrimSpace(model.search.Value())

	if query == "" {
		model.viewport.GotoTop()
		return
	}

	query = strings.ToLower(query)

	lines := model.aboutLines()

	for i, line := range lines {
		if strings.Contains(strings.ToLower(line.Text), query) {
			model.viewport.SetYOffset(i)
			return
		}
	}
}
