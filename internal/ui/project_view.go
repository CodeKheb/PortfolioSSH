package ui

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
		Title:       "CodeShare",
		Technology:  "Node.js, OAuth2, Express.js",
		Description: "DevKada Hackathon",
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
		A 2D restaurant management game built with FXGL (JavaFX) that visualizes CPU scheduling algorithms through a restaurant.
		Customers arrive as processes with Arrival Time (AT) and Burst Time (BT), 
		simulating a First Come First Serve (FCFS) scheduling algorithm.
		The game also includes a rhythm minigame where players click to serve customers at the right time.
		`,
	},

	// TODO: ADD MORE PROJECTS
}

func (model Model) projectLines() []ContentLine {
	lines := []ContentLine{
		{
			Text:  "Projects",
			Style: titleStyle,
		},
	}

	for i, project := range projectItems {
		if i == model.selected {
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
					Text:  "Description:",
					Style: descriptionStyle,
				},
				ContentLine{
					Text:  project.Description,
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
	return model.layout(
		model.renderLines(model.projectLines()),
	)
}
