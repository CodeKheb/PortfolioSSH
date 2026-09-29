package ssh

import (
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	gossh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"

	"github.com/CodeKheb/PortfolioSSH/internal/database"
	"github.com/CodeKheb/PortfolioSSH/internal/metrics"
	"github.com/CodeKheb/PortfolioSSH/internal/ratelimiter"
	"github.com/CodeKheb/PortfolioSSH/internal/ui"
)

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func MainServer(database *database.Database, limiter *ratelimiter.Limiter) (*gossh.Server, error) {
	server, err := wish.NewServer(
		wish.WithAddress(env("SSH_ADDRESS", ":42069")),
		wish.WithHostKeyPath(env("SSH_HOST_KEY", "./.docker-data/host_key")),

		wish.WithMiddleware(
			wishbubbletea.Middleware(
				func(session gossh.Session) (tea.Model, []tea.ProgramOption) {
					metrics.Sessions.Inc()

					host, _, err := net.SplitHostPort(session.RemoteAddr().String())
					if err != nil {
						host = session.RemoteAddr().String()
					}
					return ui.ViewportModel(database, limiter, host), []tea.ProgramOption{
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
