package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// The Middleware type is an adapter to allow the use of ordinary functions as HTTP middleware.
type Middleware func(http.Handler) http.Handler

// The NewMiddleware function creates a new Middleware type.
func NewMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

// The Logger middleware logs the request method and URL.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("Request", slog.String("method", r.Method), slog.String("url", r.URL.Path), slog.Duration("duration", time.Since(start)))

	})
}
