package middleware

import (
	"net/http"
	"time"
)

type Metrics interface {
	RequestsTotal(method string, status int, path string)
	Latency(method string, status int, path string, d time.Duration)
}

func MetricsMiddleware(metrics Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			rw := NewCommonResponseWriter(w)
			next.ServeHTTP(rw, r)

			duration := time.Since(startTime)

			metrics.RequestsTotal(r.Method, rw.statusCode, r.URL.Path)
			metrics.Latency(r.Method, rw.statusCode, r.URL.Path, duration)

		})
	}
}
