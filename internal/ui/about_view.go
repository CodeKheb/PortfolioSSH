package ui

const aboutContent = `

	Kherbin Clloyde Buenaventura

	20 years old
	Located at Bataan, Philippines

	Computer Science Student @ Bataan Peninsula State University
	Vice President Externals - Student Society of Information Technology Education

	This is an SSH portfolio written in the golang programming language and deployed in a Google VM.

	A question you might ask is "Why SSH?"

	Instead of making a website portfolio and deploying it on a platform that does everythhing for me
	I wanted to learn how deployment works internally.
	Learning things like:
		- Terraform
		- Google Cloud 
		- Networking
		- SSH

	It is also a good insight into my terminal centric workflow.

	I am an aspiring DevOps engineer, and this is also why I made a terminal portfolio.


	Education:

	Bachelor of Science in Computer Science 
	Bataan Peninsula State University | 2025 - Present
		- President's Lister 
		- GWA (1.3x)
	
	Technical Vocational and Livelihood - Information and Communications Technology
	Bataan National High School | 2023 - 2025
		- With High Honors
		- Awarded Outstanding in Computer Systems Servicing (Grade: 100%)


	Experiences:

	Regional Assembly on Information Technology Education (RAITE)
	Quiz Bee Representative | 2025
	Nueva Ecija University of Science and Technology
		- Competed as the Quiz Bee Representative for Bataan Peninsula State University

	Student Society of Information Technology Education
	Vice Chairperson for External Affairs | 2025
		- Helped with the first external sponsorship of SSITE
	
	National Service Training Program (NSTP)
	Project Lead | 2026
	Leadership Awardee | 2026
		- Lead the Literacy Training Service cluster to complete a Campus Based Project
		  under significant budget and time constraints.

	
	Skillset:

	DevOps and Infrastructre:
		- Docker
		- Terraform
		- GitHub actions
		- Linux

	Backend and Database:
		- Golang
		- Node.js
		- Java
		- PostgreSQL
		- SQLite
	
	Systems Programming and Embedded:
		- C
		- Rust
		- ESP32
	
	Scripting and Automation:
		- Bash
		- Lua
		- Python
		- Maven
		- Gradle

	Frontend and Frameworks:
		- Tauri v2 (Rust)
		- Express.js
		- React.js
		- Tailwind.js
		- Alpine.js
	
	Tools for Development:
		- Arch Linux
		- Neovim
		- Blender
		- Godot
		- Git
	


	Certifications:
	IC3 Digital Literacy Certification - Level 1 | October 2025
	ITS: Java Certification | April 2026
`


func (model Model) AboutView() string {
	header := model.headerView("ABOUT KHERBIN BUENAVENTURA")
	return model.layoutWithHeader(
		header,
		model.viewport.View(),
	)
}
