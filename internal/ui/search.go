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
		model.searchMatches = nil
		model.searchIndex = 0
		model.viewport.GotoTop()
		return
	}

	query = strings.ToLower(query)

	lines := model.aboutLines()

	model.searchMatches = nil

	for i, line := range lines {
		if strings.Contains(strings.ToLower(line.Text), query) {
			model.searchMatches = append(model.searchMatches, i)
		}
	}

	if len(model.searchMatches) == 0 {
		return
	}

	model.searchIndex = 0
	model.viewport.SetYOffset(model.searchMatches[model.searchIndex])
}

func (model *Model) nextSearchMatch() {
	if len(model.searchMatches) == 0 {
		return
	}

	model.searchIndex++

	if model.searchIndex >= len(model.searchMatches) {
		model.searchIndex = 0
	}

	model.viewport.SetYOffset(
		model.searchMatches[model.searchIndex],
	)
}

func (model *Model) previousSearchMatch() {
	if len(model.searchMatches) == 0 {
		return
	}

	model.searchIndex--

	if model.searchIndex < 0 {
		model.searchIndex = len(model.searchMatches) - 1
	}

	model.viewport.SetYOffset(
		model.searchMatches[model.searchIndex],
	)
}
