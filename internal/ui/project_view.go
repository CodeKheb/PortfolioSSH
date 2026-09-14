package ui

import "strings"

type ProjectItems struct {
	Title       string
	Technology  string
	Description string
	// TODO: Add more larp here
}

var projectItems = []ProjectItems{
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

func (model Model) ProjectView() string {
	var builder strings.Builder

	builder.WriteString(titleStyle.Render("Projects"))
	builder.WriteString("\n\n")

	for i, items := range projectItems {

		if i == model.selected {
			builder.WriteString(selectedItemStyle.Render(">" + projectItems[i].Title))
			builder.WriteString("\n")
			builder.WriteString(descriptionStyle.Render(
				projectItems[i].Technology,
			))
			builder.WriteString("\n\n")
			builder.WriteString(descriptionStyle.Render(
				"Description:\n" +
				projectItems[i].Description,
			))
		} else {
			builder.WriteString(itemStyle.Render(" " + items.Title))
		}
		builder.WriteString("\n\n\n")
	}

	return model.layout(builder.String())
}
