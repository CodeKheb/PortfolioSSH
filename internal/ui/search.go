package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ContentLine struct
// Style contains a set of rules that comprise a style as a whole
// gets called as an array and called by {currentScreen}Line() functions
type ContentLine struct {
	Text  string
	Style lipgloss.Style
}

// render the Lines
// gets called by {currentScreen}Content()
func (model Model) renderLines(lines []ContentLine) string {
	var builder strings.Builder
	// query gets the input value
	query := model.search.Value()

	// search every index int ContentLine array
	for i, line := range lines {
		// highlights the line that is searched
		text := highlightSearch(line.Text, query)
		builder.WriteString(line.Style.Render(text))

		// add a new line in between lines
		if i < len(lines)-1 {
			builder.WriteString("\n")
		}

	}
	return builder.String()
}

// search highlighting gets called in renderLines()
func highlightSearch(text string, query string) string {
	// remove trailing whitespace of the search
	query = strings.TrimSpace(query)

	// if no search, return
	if query == "" {
		return text
	}

	// lowerCase the texts and the search
	lowerText := strings.ToLower(text)
	lowerQuery := strings.ToLower(query)

	var builder strings.Builder
	start := 0

	// find the highlighted text
	for {
		// compares the text starting at start: and the search
		index := strings.Index(lowerText[start:], lowerQuery)

		// if search found nothing
		if index == -1 {
			builder.WriteString(text[start:])
			break
		}

		// adjust the index to start
		index += start

		// no highlight before the match
		builder.WriteString(text[start:index])
		// highlight the match
		builder.WriteString(
			searchHighlightStyle.Render(
				// from the search index to the search length
				text[index : index+len(query)],
			),
		)

		// move the start to the next match
		start = index + len(query)
	}

	return builder.String()
}

// gets called in searchContent()
// handles the different screens
func (model Model) searchableLines() []ContentLine {
	switch model.screen {
	case AboutScreen:
		return model.aboutLines()
	case ProjectScreen:
		return model.projectLines()

	// TODO: More Screen cases

	default:
		// return model.menuLines()
		return nil
	}
}

// search content, scrolls the viewport to the searched
func (model *Model) searchContent() {
	// remove trailing whitespace of the search
	query := strings.TrimSpace(model.search.Value())

	// if search empty, go to top
	if query == "" {
		model.searchMatches = nil
		model.searchIndex = 0
		model.viewport.GotoTop()
		return
	}

	// lower case both the lines in the currentScreen
	// and the user search
	query = strings.ToLower(query)
	lines := model.searchableLines()

	// initialize as null
	model.searchMatches = nil

	// loop to find a match
	for i, line := range lines {
		// if the searchableLines contains the search
		// append the index of the match
		if strings.Contains(strings.ToLower(line.Text), query) {
			model.searchMatches = append(model.searchMatches, i)
		}
	}

	// if no match, return
	if len(model.searchMatches) == 0 {
		return
	}

	// jump to first result
	model.searchIndex = 0
	model.viewport.SetYOffset(model.searchMatches[model.searchIndex])
}

// jump to next search
// called inside input.go with keypress "n"
func (model *Model) nextSearchMatch() {
	// no match return
	if len(model.searchMatches) == 0 {
		return
	}

	// jump next match
	model.searchIndex++

	// if already last, go back
	if model.searchIndex >= len(model.searchMatches) {
		model.searchIndex = 0
	}

	match := model.searchMatches[model.searchIndex]
	query := strings.TrimSpace(model.search.Value())

	// projects screen
	if model.screen == ProjectScreen {
		for i, project := range projectItems {

			searchable := strings.ToLower(
				project.Title + " " +
					project.Technology + " " +
					project.Description,
			)

			if strings.Contains(searchable, query) {
				model.searchMatches = append(
					model.searchMatches,
					i,
				)
			}
		}
		model.selected = match
		model.updateContent()
		return
	}

	model.viewport.SetYOffset(match)
}

// jump to previous search
// called inside input.go with keypress "N"
func (model *Model) previousSearchMatch() {
	if len(model.searchMatches) == 0 {
		return
	}

	model.searchIndex--

	if model.searchIndex < 0 {
		model.searchIndex = len(model.searchMatches) - 1
	}

	match := model.searchMatches[model.searchIndex]
	query := strings.TrimSpace(model.search.Value())

	// projects screen
	if model.screen == ProjectScreen {
		for i, project := range projectItems {

			searchable := strings.ToLower(
				project.Title + " " +
					project.Technology + " " +
					project.Description,
			)

			if strings.Contains(searchable, query) {
				model.searchMatches = append(
					model.searchMatches,
					i,
				)
			}
		}
		model.selected = match
		model.updateContent()
		return
	}

	model.viewport.SetYOffset(match)
}

// search project_view
func (model *Model) searchProjects() {
	query := strings.TrimSpace(model.search.Value())

	if query == "" {
		model.searchMatches = nil
		model.searchIndex = 0
		model.updateContent()
		return
	}

	query = strings.ToLower(query)

	model.searchMatches = nil

	for i, project := range projectItems {
		searchable := strings.ToLower(
			project.Title + " " +
				project.Technology + " " +
				project.Description,
		)

		if strings.Contains(searchable, query) {
			model.searchMatches = append(
				model.searchMatches,
				i,
			)
		}
	}

	if len(model.searchMatches) == 0 {
		return
	}

	model.searchIndex = 0
	model.selected = model.searchMatches[0]

	model.updateContent()
}

func (model *Model) updateContent() {
	switch model.screen {
	case AboutScreen:
		model.viewport.SetContent(
			model.renderLines(model.aboutLines()),
		)

	case ProjectScreen:
		model.viewport.SetContent(
			model.renderLines(model.projectLines()),
		)

		// TODO: Add more screens
	}
}
