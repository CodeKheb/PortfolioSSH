package main

import (
	"log"

	sshserver "github.com/CodeKheb/PortfolioSSH/internal/ssh"
)

// Main program
// starts the lipgloss UI
func main() {
	server, err := sshserver.MainServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

//	program := tea.NewProgram(
//		ui.ViewportModel(),
//		tea.WithAltScreen(),
//	)
//	if _, err := program.Run()
//
//	err != nil {
//		fmt.Print(err)
//	}
}
