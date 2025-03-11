package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TODO: где-то накосячили. проверить
type Metrics struct {
	requestsTotal *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Number of incoming requests to API",
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

//TODO: latecy
