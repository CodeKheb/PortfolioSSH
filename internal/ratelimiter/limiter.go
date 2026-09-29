package ratelimiter

import (
	"sync"
	"time"
)

type Limiter struct {
	mutex      sync.Mutex
	clients map[string]time.Time
	window  time.Duration
}

func New(window time.Duration) *Limiter {
	return &Limiter{
		clients: make(map[string]time.Time),
		window:  window,
	}
}

func (limit *Limiter) Allow(ip string) bool {
	limit.mutex.Lock()
	defer limit.mutex.Unlock()

	last, exists := limit.clients[ip]

	if exists && time.Since(last) < limit.window {
		return false
	}

	limit.clients[ip] = time.Now()
	return true
}


