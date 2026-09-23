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

type Project struct {
	Title       string
	Technology  string
	Description string
	RepoURL     string
	README      string
}

var markdownParser = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
)

var projectItems = []Project{
	{
		Title:       "Portfolio",
		Technology:  "Go, Terraform, GCP, SSH",
		Description: "Portfolio larp",
		RepoURL:     "https://github.com/CodeKheb/PortfolioSSH",
	},
	{
		Title:      "CodeShare",
		Technology: "Node.js, OAuth2, Express.js",
		Description: `
		DevKada Hackathon
		CodeShare unifies real-time team chat with live GitHub integration. 
		Attach a repository to a group chat, and every push, pull request, 
		and release instantly surfaces as a system message right where your team is already talking. 
		No notifications to miss. No context to lose. One conversation. Everything you need.
		`,
		RepoURL: "https://github.com/CodeKheb/HelloWorld-CodeShare",
	},
	{
		Title:      "ESPresso",
		Technology: "ESP32, Tauri v2 (Rust), TypeScript, React",
		Description: `
		Tauri 2 + React + TypeScript desktop app for sharing coffee profiles over any
		local network. **The WiFi is the DNS** — every device with ESPresso open hosts
		a pot, and pots on the same network find each other automatically via mDNS.`,
		RepoURL: "https://github.com/CodeKheb/ESPresso",
	},
	{
		Title:      "typetest_TUI",
		Technology: "Go, Bubbletea, Lipgloss",
		Description: `
		A terminal-based typing speed test written in Go, using
		[Bubbletea](https://github.com/charmbracelet/bubbletea) and
		[Lipgloss](https://github.com/charmbracelet/lipgloss) for the UI.
		`,
		RepoURL: "https://github.com/CodeKheb/typetest_TUI",
	},
	{
		Title:      "OrderUp",
		Technology: "JavaFX, FXGL",
		Description: `
		A 2D restaurant management game built with FXGL (JavaFX) 
		that visualizes CPU scheduling algorithms through a restaurant.
		Customers arrive as processes with Arrival Time (AT) and Burst Time (BT), 
		simulating a First Come First Serve (FCFS) scheduling algorithm.
		The game also includes a rhythm minigame where players click to serve customers at the right time.
		`,
		RepoURL: "https://github.com/CodeKheb/OrderUp",
	},
	{
		Title:      "dotfiles",
		Technology: "Neovim, Linux, Lua, Bash",
		Description: `
		Personal Linux configuration files, 
		built around an Arch-based setup with Neovim as the primary editor, 
		a Sway (Wayland) desktop, and a set of rofi-powered launcher scripts.
		`,
		RepoURL: "https://github.com/CodeKheb/dotfiles",
	},

	// TODO: ADD MORE PROJECTS
}

func (model Model) projectLines() []ContentLine {
	lines := []ContentLine{}

	lines = append(lines, ContentLine{
		Text: "\n",
	})

	for i, project := range projectItems {
		if i == model.selected {
			description := parseLine(project.Description)

			lines = append(

				lines,
				ContentLine{
					Text:  "> " + project.Title,
					Style: selectedItemStyle,
				},
				ContentLine{
					Text:  project.Technology,
					Style: descriptionStyle,
				},
				ContentLine{
					Text:  "",
					Style: descriptionStyle,
				},
				ContentLine{
					Text:  description,
					Style: descriptionStyle,
				},
			)
		} else {
			lines = append(
				lines,
				ContentLine{
					Text:  "  " + project.Title,
					Style: itemStyle,
				},
				ContentLine{
					Text: "\n",
				},
			)
		}

		lines = append(lines, ContentLine{})
	}

	return lines
}

func (model Model) projectContent() string {
	return model.renderLines(model.projectLines())
}

func (model Model) ProjectView() string {
	header := model.headerView("PROJECTS")

	return model.layout(
		header,
		model.renderLines(model.projectLines()),
	)
}

func parseLine(line string) string {
	lines := strings.Split(strings.TrimSpace(line), "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return strings.Join(lines, "\n")
}

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
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		switch row := child.(type) {
		case *extast.TableHeader:
			parseTableRow(source, row, lines, true)

		case *extast.TableRow:
			parseTableRow(source, row, lines, false)
		}
	}
}

func parseTableRow(source []byte, row ast.Node, lines *[]ContentLine, header bool) {
	var cells []string

	for cell := row.FirstChild(); cell != nil; cell = cell.NextSibling() {
		if cell.Kind() != extast.KindTableCell {
			continue
		}

		cells = append(cells, strings.TrimSpace(
			nodeText(source, cell),
		))
	}

	style := bodyStyle
	if header {
		style = subheadingStyle
	}

	*lines = append(*lines, ContentLine{
		Text:  strings.Join(cells, " │ "),
		Style: style,
	})
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
