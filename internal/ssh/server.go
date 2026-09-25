package ssh

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	gossh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"

	"github.com/CodeKheb/PortfolioSSH/internal/metrics"
	"github.com/CodeKheb/PortfolioSSH/internal/ui"
)

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func MainServer() (*gossh.Server, error) {
	server, err := wish.NewServer(
		wish.WithAddress(env("SSH_ADDRESS", ":42069")),
		wish.WithHostKeyPath(env("SSH_HOST_KEY", "./.docker-data/host_key")),

		wish.WithMiddleware(
			wishbubbletea.Middleware(
				func(session gossh.Session) (tea.Model, []tea.ProgramOption) {
					metrics.Sessions.Inc()

					return ui.ViewportModel(), []tea.ProgramOption{
						tea.WithAltScreen(),
					}
				},
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create server: %w", err)
	}
	return server, nil
}

// test ci
