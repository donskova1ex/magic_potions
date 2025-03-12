package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TODO: где-то накосячили. проверить
type Metrics struct {
	requestsTotal *prometheus.CounterVec
	latency       *prometheus.HistogramVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Number of incoming requests to API",
		}, []string{"method", "status", "path"}),
		latency: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "request_duration",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
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

func (m *Metrics) Latency(method string, status int, path string, d time.Duration) {
	m.latency.With(prometheus.Labels{
		"method": method,
		"status": strconv.Itoa(status),
		"path":   path,
	}).Observe(d.Seconds())
}
