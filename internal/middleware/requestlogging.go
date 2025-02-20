package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func RequestIDMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return newRequestWithID(next, logger)
	}
}
func newRequestWithID(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.NewString()

		logger.Info(
			"New request with: ",
			slog.String("requestID", requestID),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
