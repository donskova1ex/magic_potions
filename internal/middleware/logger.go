package middleware

import (
	"log/slog"
	"net/http"
)

// TODO: off default logger
func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			requestID, _ := r.Context().Value(RequestIDCtxKey).(RequestID)
			logger.Info(
				"http request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				//slog.Int64("status", int64(r.Response.StatusCode)),
				slog.String("request_id", string(requestID)),
			)
		})
	}
}
