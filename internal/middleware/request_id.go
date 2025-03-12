package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestID string

const RequestIDCtxKey RequestID = "request_id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := RequestID(uuid.NewString())
		ctx := context.WithValue(r.Context(), RequestIDCtxKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
