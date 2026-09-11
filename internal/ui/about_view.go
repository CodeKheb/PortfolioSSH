package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type AboutLine struct {
	Text  string
	Style lipgloss.Style
}

func (model Model) aboutLines() []AboutLine {
	return []AboutLine{
		// Identity
		{"", bodyStyle},
		{"", bodyStyle},
		{"Kherbin Clloyde Buenaventura", nameStyle},
		{"20 years old", mutedStyle},
		{"Based in Bataan, Philippines", mutedStyle},
		{"", bodyStyle},
		{"", bodyStyle},

		{"Computer Science Student @ Bataan Peninsula State University", bodyStyle},
		{"Vice President for Externals - Student Society of Information Technology Education", bodyStyle},
		{"", bodyStyle},
		{"", bodyStyle},

		// Introduction
		{"This is an SSH portfolio written in the Go programming language and deployed in a Google VM.", bodyStyle},
		{"", bodyStyle},

		{`A question you might ask is "Why SSH?"`, bodyStyle},
		{"", bodyStyle},

		{"Instead of making a website portfolio and deploying it on a platform that does everything for me,", bodyStyle},
		{"I wanted to learn how deployment works internally.", bodyStyle},
		{"", bodyStyle},

		{"Learning things like:", bodyStyle},
		{"    • Terraform", bodyStyle},
		{"    • Google Cloud", bodyStyle},
		{"    • Networking", bodyStyle},
		{"    • SSH", bodyStyle},
		{"", bodyStyle},

		{"It is also a good insight into my terminal-centric workflow.", bodyStyle},
		{"", bodyStyle},

		{"I am an aspiring DevOps engineer, and this is also why I made a terminal portfolio.", bodyStyle},
		{"", bodyStyle},

		// Education
		{"EDUCATION", sectionStyle},
		{"", bodyStyle},

		{"Bachelor of Science in Computer Science", subheadingStyle},
		{"Bataan Peninsula State University | 2025 - Present", mutedStyle},
		{"    • President's Lister", bodyStyle},
		{"    • GWA (1.3x)", bodyStyle},
		{"", bodyStyle},

		{
			"Technical Vocational and Livelihood - Information and Communications Technology",
			subheadingStyle,
		},
		{"Bataan National High School | 2023 - 2025", mutedStyle},
		{"    • With High Honors", bodyStyle},
		{
			"    • Awarded Outstanding in Computer Systems Servicing (Grade: 100%)",
			bodyStyle,
		},
		{"", bodyStyle},

		// Experience
		{"EXPERIENCES", sectionStyle},
		{"", bodyStyle},

		{
			"Regional Assembly on Information Technology Education (RAITE)",
			subheadingStyle,
		},
		{"Quiz Bee Representative | 2025", mutedStyle},
		{"Nueva Ecija University of Science and Technology", mutedStyle},
		{
			"    • Competed as the Quiz Bee Representative for Bataan Peninsula State University",
			bodyStyle,
		},
		{"", bodyStyle},

		{
			"Student Society of Information Technology Education",
			subheadingStyle,
		},
		{"Vice Chairperson for External Affairs | 2025", mutedStyle},
		{
			"    • Helped with the first external sponsorship of SSITE",
			bodyStyle,
		},
		{"", bodyStyle},

		{"National Service Training Program (NSTP)", subheadingStyle},
		{"Project Lead | 2026", mutedStyle},
		{"Leadership Awardee | 2026", mutedStyle},
		{
			"    • Led the Literacy Training Service cluster to complete a Campus Based Project",
			bodyStyle,
		},
		{
			"      under significant budget and time constraints.",
			bodyStyle,
		},
		{"", bodyStyle},

		// Skills
		{"SKILLSET", sectionStyle},
		{"", bodyStyle},

		{"DevOps and Infrastructure", subheadingStyle},
		{"    • Docker", bodyStyle},
		{"    • Terraform", bodyStyle},
		{"    • GitHub Actions", bodyStyle},
		{"    • Linux", bodyStyle},
		{"", bodyStyle},

		{"Backend and Database", subheadingStyle},
		{"    • Go", bodyStyle},
		{"    • Node.js", bodyStyle},
		{"    • Java", bodyStyle},
		{"    • PostgreSQL", bodyStyle},
		{"    • SQLite", bodyStyle},
		{"", bodyStyle},

		{"Systems Programming and Embedded", subheadingStyle},
		{"    • C", bodyStyle},
		{"    • Rust", bodyStyle},
		{"    • ESP32", bodyStyle},
		{"", bodyStyle},

		{"Scripting and Automation", subheadingStyle},
		{"    • Bash", bodyStyle},
		{"    • Lua", bodyStyle},
		{"    • Python", bodyStyle},
		{"    • Maven", bodyStyle},
		{"    • Gradle", bodyStyle},
		{"", bodyStyle},

		{"Frontend and Frameworks", subheadingStyle},
		{"    • Tauri v2 (Rust)", bodyStyle},
		{"    • Express.js", bodyStyle},
		{"    • React.js", bodyStyle},
		{"    • Tailwind CSS", bodyStyle},
		{"    • Alpine.js", bodyStyle},
		{"", bodyStyle},

		{"Tools for Development", subheadingStyle},
		{"    • Arch Linux", bodyStyle},
		{"    • Neovim", bodyStyle},
		{"    • Blender", bodyStyle},
		{"    • Godot", bodyStyle},
		{"    • Git", bodyStyle},
		{"", bodyStyle},

		// Certifications
		{"CERTIFICATIONS", sectionStyle},
		{"", bodyStyle},

		{"IC3 Digital Literacy Certification - Level 1", subheadingStyle},
		{"October 2025", mutedStyle},
		{"", bodyStyle},

		{"ITS: Java Certification", subheadingStyle},
		{"April 2026", mutedStyle},
	}
}

func (model Model) aboutContent() string {
	var builder strings.Builder
	query := model.search.Value()
	lines := model.aboutLines()

	for i, line := range lines {
		text := highlightSearch(line.Text, query)
		builder.WriteString(line.Style.Render(text))

		if i < len(lines) -1 {
			builder.WriteString("\n")
		}

	}
	return builder.String()
}

func (model Model) AboutView() string {
	header := model.headerView("ABOUT 	KHERBIN BUENAVENTURA")

	leftMargin := max(0, (model.width-80)/2)

	content := lipgloss.NewStyle().
		MarginLeft(leftMargin).
		Render(model.viewport.View())

	return model.layoutWithHeader(
		header,
		content,
	)
}
