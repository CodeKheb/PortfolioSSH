package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Sessions = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "portfolio_ssh_sessions",
		Help: "Total number of SSH sessions.",
	},
)

func init() {
	prometheus.MustRegister(Sessions)
}

func Start(address string) error {
	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(address, nil)
}
