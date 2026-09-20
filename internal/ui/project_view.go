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
