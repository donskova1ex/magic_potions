package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	requestsTotal *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Number of incoming requets to API",
		}, []string{"method", "status", "path"}),
	}
}

func (m *Metrics) RequestsTotal(method string, status int, path string) {
	m.requestsTotal.With(prometheus.Labels{
		"method": method,
		"status": strconv.Itoa(status),
		"path":   path,
	}).Inc()
}
