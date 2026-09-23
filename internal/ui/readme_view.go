package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

var markdownParser = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
)

func parseREADME(markdown string) []ContentLine {
	source := []byte(markdown)

	reader := text.NewReader(source)
	docs := markdownParser.Parser().Parse(reader)

	var lines []ContentLine

	for node := docs.FirstChild(); node != nil; node = node.NextSibling() {
		switch node.Kind() {

		case ast.KindHeading:
			heading := node.(*ast.Heading)

			lines = append(lines, ContentLine{
				Text:  string(nodeText(source, heading)),
				Style: sectionStyle,
			})

		case ast.KindParagraph:
			lines = append(lines, ContentLine{
				Text:  string(nodeText(source, node)),
				Style: bodyStyle,
			})

		case ast.KindList:
			parseList(source, node, &lines)

		case ast.KindFencedCodeBlock:
			parseCodeBlock(source, node, &lines)

		case extast.KindTable:
			parseTable(source, node.(*extast.Table), &lines)

		case ast.KindThematicBreak:
			lines = append(lines, ContentLine{
				Text:  "────────────────────────────────",
				Style: mutedStyle,
			})
		case ast.KindHTMLBlock:
			continue

		case ast.KindRawHTML:
			continue
		}
		lines = append(lines, ContentLine{
			Text:  "",
			Style: bodyStyle,
		})
	}
	return lines
}

func nodeText(source []byte, node ast.Node) string {
	var builder strings.Builder

	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if textNode, ok := n.(*ast.Text); ok {
			builder.Write(textNode.Segment.Value(source))
		}

		return ast.WalkContinue, nil
	})
	return builder.String()
}

func parseList(source []byte, node ast.Node, lines *[]ContentLine) {
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		if child.Kind() != ast.KindListItem {
			continue
		}

		*lines = append(*lines, ContentLine{
			Text:  "• " + nodeText(source, child),
			Style: bodyStyle,
		})
	}
}

func parseCodeBlock(source []byte, node ast.Node, lines *[]ContentLine) {
	codeBlock := node.(*ast.FencedCodeBlock)

	for i := 0; i < codeBlock.Lines().Len(); i++ {
		line := codeBlock.Lines().At(i)

		*lines = append(*lines, ContentLine{
			Text:  string(line.Value(source)),
			Style: codeStyle,
		})
	}
}

func parseTable(source []byte, node *extast.Table, lines *[]ContentLine) {
	var rows [][]string

	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		var cells []string

		for cell := child.FirstChild(); cell != nil; cell = cell.NextSibling() {
			cells = append(cells, strings.TrimSpace(
				nodeText(source, cell),
			))
		}

		if len(cells) > 0 {
			rows = append(rows, cells)
		}
	}

	parseTableRow(rows, lines)
}

func parseTableRow(rows [][]string, lines *[]ContentLine) {
	if len(rows) == 0 {
		return
	}

	columnCount := 0
	for _, row := range rows {
		if len(row) > columnCount {
			columnCount = len(row)
		}
	}

	widths := make([]int, columnCount)

	for _, row := range rows {
		for i, cell := range row {
			width := lipgloss.Width(cell)

			if width > widths[i] {
				widths[i] = width
			}
		}
	}

	for rowIndex, row := range rows {
		var cells []string

		for i := 0; i < columnCount; i++ {
			cell := ""

			if i < len(row) {
				cell = row[i]
			}

			padding := widths[i] - lipgloss.Width(cell)

			cells = append(
				cells,
				cell+strings.Repeat(" ", padding),
			)
		}

		style := bodyStyle

		if rowIndex == 0 {
			style = subheadingStyle
		}

		*lines = append(*lines, ContentLine{
			Text:  strings.Join(cells, " │ "),
			Style: style,
		})
	}
}

func (model Model) READMEView() string {
	project := projectItems[model.selected]

	header := model.headerView("PROJECT: " + project.Title)

	content := lipgloss.NewStyle().
		Render(model.viewport.View())

	return model.layout(
		header,
		content,
	)
}

func (model Model) readmeLines() []ContentLine {
	project := projectItems[model.selected]

	if project.README == "" {
		return []ContentLine{
			{
				Text:  "I'm lazy and did not write a README for this yet",
				Style: mutedStyle,
			},
		}
	}

	return parseREADME(project.README)
}

func (model Model) readmeContent() string {
	return model.renderLines(model.readmeLines())
}
