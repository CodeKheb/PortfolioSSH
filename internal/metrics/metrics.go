package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Sessions = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "portfolio_ssh_sessions",
			Help: "Total number of SSH sessions.",
		},
	)

	MessagesReceived = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "portfolio_messages_received",
			Help: "Total number of messages received.",
		},
	)

	MessagesFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "portfolio_messages_failed",
			Help: "Total number of messages failed.",
		},
	)
)

func init() {
	prometheus.MustRegister(Sessions)
	prometheus.MustRegister(MessagesReceived)
	prometheus.MustRegister(MessagesFailed)
}

func Start(address string) error {
	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(address, nil)
}
