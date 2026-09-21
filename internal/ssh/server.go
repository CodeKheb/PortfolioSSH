package ssh

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	gossh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"

	"github.com/CodeKheb/PortfolioSSH/internal/ui"
)

func MainServer() (*gossh.Server, error) {
	server, err := wish.NewServer(
		wish.WithAddress("localhost:42069"),
		wish.WithHostKeyPath("host_key"),

		wish.WithMiddleware(
			wishbubbletea.Middleware(
				func(session gossh.Session) (tea.Model, []tea.ProgramOption) {
					return ui.ViewportModel(), []tea.ProgramOption{
						tea.WithAltScreen(),
					}
				},
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("Err: %w", err)
	}
	return server, nil
}
