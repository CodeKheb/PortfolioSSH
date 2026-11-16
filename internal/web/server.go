package web

import (
	"log"
	"net/http"
	"time"

	"github.com/CodeKheb/PortfolioSSH/internal/database"
	"github.com/CodeKheb/PortfolioSSH/internal/notifier"
	"github.com/CodeKheb/PortfolioSSH/internal/ratelimiter"
)

type Server struct {
	db       *database.Database
	limiter  *ratelimiter.Limiter
	telegram *notifier.Telegram
}

func NewServer(db *database.Database, limiter *ratelimiter.Limiter, telegram *notifier.Telegram) *Server {
	return &Server{
		db:       db,
		limiter:  limiter,
		telegram: telegram,
	}
}

func (s *Server) Start(address string) error {
	server := &http.Server{
		Addr: address,
		Handler: s.routes(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Printf("Starting web at %s", address)

	return server.ListenAndServe()
}
