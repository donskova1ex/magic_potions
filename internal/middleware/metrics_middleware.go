package middleware

import "net/http"

type Metrics interface {
	RequestsTotal(method string, status int, path string)
}

func MetricsMiddleware(metrics Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := NewCommonResponseWriter(w)
			next.ServeHTTP(rw, r)
			metrics.RequestsTotal(r.Method, rw.statusCode, r.URL.Path)
		})
	}
}
