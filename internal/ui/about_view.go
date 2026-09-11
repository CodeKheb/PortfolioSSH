package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func aboutContent() string {
	var builder strings.Builder

	// Identity
	builder.WriteString("\n\n")
	builder.WriteString(nameStyle.Render("Kherbin Clloyde Buenaventura"))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render("20 years old"))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render("Based in Bataan, Philippines"))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render(
		"Computer Science Student @ Bataan Peninsula State University",
	))
	builder.WriteString("\n")
	builder.WriteString(bodyStyle.Render(
		"Vice President for Externals - Student Society of Information Technology Education",
	))
	builder.WriteString("\n\n")

	// Introduction
	builder.WriteString(bodyStyle.Render(
		"This is an SSH portfolio written in the Go programming language and deployed in a Google VM.",
	))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render(
		`A question you might ask is "Why SSH?"`,
	))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render(
		"Instead of making a website portfolio and deploying it on a platform that does everything for me,",
	))
	builder.WriteString("\n")
	builder.WriteString(bodyStyle.Render(
		"I wanted to learn how deployment works internally.",
	))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render("Learning things like:"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Terraform"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Google Cloud"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Networking"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("SSH"))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render(
		"It is also a good insight into my terminal-centric workflow.",
	))
	builder.WriteString("\n\n")

	builder.WriteString(bodyStyle.Render(
		"I am an aspiring DevOps engineer, and this is also why I made a terminal portfolio.",
	))
	builder.WriteString("\n\n")

	// Education
	builder.WriteString(sectionStyle.Render("EDUCATION"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"Bachelor of Science in Computer Science",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Bataan Peninsula State University | 2025 - Present",
	))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("President's Lister"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("GWA (1.3x)"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"Technical Vocational and Livelihood - Information and Communications Technology",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Bataan National High School | 2023 - 2025",
	))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("With High Honors"))
	builder.WriteString("\n")
	builder.WriteString(
		bulletStyle.Render("    • ") +
			bodyStyle.Render("Awarded Outstanding in Computer Systems Servicing (Grade: 100%)"),
	)
	builder.WriteString("\n\n")

	// Experience
	builder.WriteString(sectionStyle.Render("EXPERIENCES"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"Regional Assembly on Information Technology Education (RAITE)",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Quiz Bee Representative | 2025",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Nueva Ecija University of Science and Technology",
	))
	builder.WriteString("\n")
	builder.WriteString(
		bulletStyle.Render("    • ") +
			bodyStyle.Render("Competed as the Quiz Bee Representative for Bataan Peninsula State University"),
	)
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"Student Society of Information Technology Education",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Vice Chairperson for External Affairs | 2025",
	))
	builder.WriteString("\n")
	builder.WriteString(
		bulletStyle.Render("    • ") +
			bodyStyle.Render("Helped with the first external sponsorship of SSITE"),
	)
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"National Service Training Program (NSTP)",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Project Lead | 2026",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render(
		"Leadership Awardee | 2026",
	))
	builder.WriteString("\n")
	builder.WriteString(
		bulletStyle.Render("    • ") +
			bodyStyle.Render("Led the Literacy Training Service cluster to complete a Campus Based Project"),
	)
	builder.WriteString("\n")
	builder.WriteString(
		bulletStyle.Render("      ") +
			bodyStyle.Render("under significant budget and time constraints."),
	)
	builder.WriteString("\n\n")

	// Skills
	builder.WriteString(sectionStyle.Render("SKILLSET"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("DevOps and Infrastructure"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Docker"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Terraform"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("GitHub Actions"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Linux"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("Backend and Database"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Go"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Node.js"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Java"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("PostgreSQL"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("SQLite"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("Systems Programming and Embedded"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("C"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Rust"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("ESP32"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("Scripting and Automation"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Bash"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Lua"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Python"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Maven"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Gradle"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("Frontend and Frameworks"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Tauri v2 (Rust)"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Express.js"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("React.js"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Tailwind CSS"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Alpine.js"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render("Tools for Development"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Arch Linux"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Neovim"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Blender"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Godot"))
	builder.WriteString("\n")
	builder.WriteString(bulletStyle.Render("    • ") + bodyStyle.Render("Git"))
	builder.WriteString("\n\n")

	// Certifications
	builder.WriteString(sectionStyle.Render("CERTIFICATIONS"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"IC3 Digital Literacy Certification - Level 1",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render("October 2025"))
	builder.WriteString("\n\n")

	builder.WriteString(subheadingStyle.Render(
		"ITS: Java Certification",
	))
	builder.WriteString("\n")
	builder.WriteString(mutedStyle.Render("April 2026"))

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
