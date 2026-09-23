package main

import (
	"fmt"
	"log"
	"time"

	sshserver "github.com/CodeKheb/PortfolioSSH/internal/ssh"
	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Main program
// starts the lipgloss UI
func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)
	go ui.PollREADME(5 * time.Minute)

	server, err := sshserver.MainServer()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("CONNECTED!!")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
