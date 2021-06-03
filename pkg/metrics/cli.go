package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	Registry   *prometheus.Registry
	CliCounter *prometheus.CounterVec
}

// NewMetrics gives you access to metrics
func NewMetrics() *Metrics {

	var (
		reg     = prometheus.NewRegistry()
		factory = promauto.With(reg)
	)

	return &Metrics{
		Registry: reg,
		CliCounter: factory.NewCounterVec(prometheus.CounterOpts{
			Name: "cli_command_executed",
			Help: "Counter to count when a client command was executed",
		}, []string{"command"}),
	}
}
