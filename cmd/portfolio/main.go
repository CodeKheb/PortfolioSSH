package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CodeKheb/PortfolioSSH/internal/database"
	"github.com/CodeKheb/PortfolioSSH/internal/metrics"
	"github.com/CodeKheb/PortfolioSSH/internal/notifier"
	"github.com/CodeKheb/PortfolioSSH/internal/ratelimiter"
	sshserver "github.com/CodeKheb/PortfolioSSH/internal/ssh"
	"github.com/CodeKheb/PortfolioSSH/internal/ui"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"github.com/muesli/termenv"
)

// Main program
// starts the lipgloss UI
func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env")
	}

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

	limiter := ratelimiter.New(5 * time.Minute)

	telegramToken := strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN"))
	telegramChatID := strings.TrimSpace(os.Getenv("TELEGRAM_CHAT_ID"))

	telegram := notifier.NewTelegram(
		telegramToken,
		telegramChatID,
	)

	server, err := sshserver.MainServer(db, limiter, telegram)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("CONNECTED!!")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
