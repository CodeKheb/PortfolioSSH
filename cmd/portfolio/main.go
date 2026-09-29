package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CodeKheb/PortfolioSSH/internal/database"
	"github.com/CodeKheb/PortfolioSSH/internal/metrics"
	sshserver "github.com/CodeKheb/PortfolioSSH/internal/ssh"
	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Main program
// starts the lipgloss UI
func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)

	databasePath := os.Getenv("DATABASE_PATH")

	if databasePath == "" {
		databasePath = "./messages.db"
	}

	db, err := database.Open(databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go ui.PollREADME(5 * time.Minute)

	go func() {
		log.Println("Starting metrics server on :8080")
		if err := metrics.Start(":8080"); err != nil {
			log.Fatal("metrics server %w", err)
		}
	}()

	server, err := sshserver.MainServer(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("CONNECTED!!")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
