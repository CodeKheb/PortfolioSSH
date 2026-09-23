package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// aboutLines, returns the ContentLine struct as an array
func (model Model) aboutLines() []ContentLine {
	return []ContentLine{
		// Identity
		{Text: "ITS: Java Certification", Style: subheadingStyle},
		{Text: "April 2026", Style: mutedStyle},
		{Text: "", Style: bodyStyle},
		{Text: "", Style: bodyStyle},
		{Text: "", Style: bodyStyle},
		{Text: "", Style: bodyStyle},
		{Text: "Kherbin Clloyde Buenaventura", Style: nameStyle},
		{Text: "Age: \t\t20 years old", Style: mutedStyle},
		{Text: "Location:    Based in Bataan, Philippines", Style: mutedStyle},
		{Text: "", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Computer Science Student @ Bataan Peninsula State University", Style: bodyStyle},
		{Text: "Vice President for Externals - Student Society of Information Technology Education", Style: bodyStyle},
		{Text: "", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		// Introduction
		{Text: "What this project is:", Style: sectionStyle},
		{
			Text:  "This is an SSH portfolio written in the Go programming language and deployed in a Google VM.",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		{Text: `A question you might ask is "Why SSH?"`, Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{
			Text:  "Instead of making a website portfolio and deploying it on a platform that does everything for me,",
			Style: bodyStyle,
		},
		{
			Text:  "I wanted to learn how deployment works internally.",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		{Text: "Learning things like:", Style: bodyStyle},
		{Text: "    • Terraform", Style: bodyStyle},
		{Text: "    • Google Cloud", Style: bodyStyle},
		{Text: "    • Networking", Style: bodyStyle},
		{Text: "    • SSH", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "It is also a good insight into my terminal-centric workflow.", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{
			Text:  "I am an aspiring DevOps engineer, and this is also why I made a terminal portfolio.",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		// Education
		{Text: "EDUCATION", Style: sectionStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Bachelor of Science in Computer Science", Style: subheadingStyle},
		{Text: "Bataan Peninsula State University | 2025 - Present", Style: mutedStyle},
		{Text: "    • President's Lister", Style: bodyStyle},
		{Text: "    • GWA (1.3x)", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{
			Text:  "Technical Vocational and Livelihood - Information and Communications Technology",
			Style: subheadingStyle,
		},
		{Text: "Bataan National High School | 2023 - 2025", Style: mutedStyle},
		{Text: "    • With High Honors", Style: bodyStyle},
		{
			Text:  "    • Awarded Outstanding in Computer Systems Servicing (Grade: 100%)",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		// Experience
		{Text: "EXPERIENCES", Style: sectionStyle},
		{Text: "", Style: bodyStyle},

		{
			Text:  "Regional Assembly on Information Technology Education (RAITE)",
			Style: subheadingStyle,
		},
		{Text: "Quiz Bee Representative | 2025", Style: mutedStyle},
		{Text: "Nueva Ecija University of Science and Technology", Style: mutedStyle},
		{
			Text:  "    • Competed as the Quiz Bee Representative for Bataan Peninsula State University",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		{
			Text:  "Student Society of Information Technology Education",
			Style: subheadingStyle,
		},
		{Text: "Vice Chairperson for External Affairs | 2025", Style: mutedStyle},
		{
			Text:  "    • Helped with the first external sponsorship of SSITE",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		{
			Text:  "National Service Training Program (NSTP)",
			Style: subheadingStyle,
		},
		{Text: "Project Lead | 2026", Style: mutedStyle},
		{Text: "Leadership Awardee | 2026", Style: mutedStyle},
		{
			Text:  "    • Led the Literacy Training Service cluster to complete a Campus Based Project",
			Style: bodyStyle,
		},
		{
			Text:  "      under significant budget and time constraints.",
			Style: bodyStyle,
		},
		{Text: "", Style: bodyStyle},

		// Skills
		{Text: "SKILLSET & LANGUAGES", Style: sectionStyle},
		{Text: "", Style: bodyStyle},

		{Text: "DevOps and Infrastructure", Style: subheadingStyle},
		{Text: "    • Docker", Style: bodyStyle},
		{Text: "    • Terraform", Style: bodyStyle},
		{Text: "    • GitHub Actions", Style: bodyStyle},
		{Text: "    • Linux", Style: bodyStyle},
		{Text: "    • Google Cloud Platform (GCP)", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Backend and Database", Style: subheadingStyle},
		{Text: "    • Go", Style: bodyStyle},
		{Text: "    • Node.js", Style: bodyStyle},
		{Text: "    • Java", Style: bodyStyle},
		{Text: "    • C#", Style: bodyStyle},
		{Text: "    • PostgreSQL", Style: bodyStyle},
		{Text: "    • SQLite", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Systems Programming and Embedded", Style: subheadingStyle},
		{Text: "    • C", Style: bodyStyle},
		{Text: "    • Rust", Style: bodyStyle},
		{Text: "    • Arduino", Style: bodyStyle},
		{Text: "    • ESP32", Style: bodyStyle},
		{Text: "    • ESP8266", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Scripting and Automation", Style: subheadingStyle},
		{Text: "    • Bash", Style: bodyStyle},
		{Text: "    • Lua", Style: bodyStyle},
		{Text: "    • Python", Style: bodyStyle},
		{Text: "    • Maven", Style: bodyStyle},
		{Text: "    • Gradle", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Frontend and Frameworks", Style: subheadingStyle},
		{Text: "    • Tauri v2 (Rust)", Style: bodyStyle},
		{Text: "    • Express.js", Style: bodyStyle},
		{Text: "    • React.js", Style: bodyStyle},
		{Text: "    • Tailwind CSS", Style: bodyStyle},
		{Text: "    • Alpine.js", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		{Text: "Tools for Development", Style: subheadingStyle},
		{Text: "    • Arch Linux", Style: bodyStyle},
		{Text: "    • Neovim", Style: bodyStyle},
		{Text: "    • Blender", Style: bodyStyle},
		{Text: "    • Godot", Style: bodyStyle},
		{Text: "    • Git", Style: bodyStyle},
		{Text: "", Style: bodyStyle},

		// Certifications
		{Text: "CERTIFICATIONS", Style: sectionStyle},
		{Text: "", Style: bodyStyle},

		{Text: "IC3 Digital Literacy Certification - Level 1", Style: subheadingStyle},
		{Text: "October 2025", Style: mutedStyle},
		{Text: "", Style: bodyStyle},
	}
}


// aboutContent calls aboutLines and pass it to renderLines() inside search.go
func (model Model) aboutContent() string {
	return model.renderLines(model.aboutLines())
}

// about view, gets switched inside View() in model.go
func (model Model) AboutView() string {
	header := model.headerView("ABOUT 	KHERBIN BUENAVENTURA")

	leftMargin := max(0, (model.width-80)/2)

	content := lipgloss.NewStyle().
		MarginLeft(leftMargin).
		Render(model.viewport.View())

	return model.layout(
		header,
		content,
	)
}
