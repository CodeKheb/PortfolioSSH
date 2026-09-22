package ui

import (
	"strings"
)

type Project struct {
	Title       string
	Technology  string
	Description string
	RepoURL     string
	README      string
	// TODO: Add more larp here
}

var projectItems = []Project{
	{
		Title:       "ProjectSSH",
		Technology:  "Go, Terraform, GCP, SSH",
		Description: "Portfolio larp",
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
	},
	{
		Title:      "ESPresso",
		Technology: "ESP32, Tauri v2 (Rust), TypeScript, React",
		Description: `
		Tauri 2 + React + TypeScript desktop app for sharing coffee profiles over any
		local network. **The WiFi is the DNS** — every device with ESPresso open hosts
		a pot, and pots on the same network find each other automatically via mDNS.`,
	},
	{
		Title:      "typetest_TUI",
		Technology: "Go, Bubbletea, Lipgloss",
		Description: `
		A terminal-based typing speed test written in Go, using
		[Bubbletea](https://github.com/charmbracelet/bubbletea) and
		[Lipgloss](https://github.com/charmbracelet/lipgloss) for the UI.
		`,
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
	},
	{
		Title:      "dotfiles",
		Technology: "Neovim, Linux, Lua, Bash",
		Description: `
		Personal Linux configuration files, 
		built around an Arch-based setup with Neovim as the primary editor, 
		a Sway (Wayland) desktop, and a set of rofi-powered launcher scripts.
		`,
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
