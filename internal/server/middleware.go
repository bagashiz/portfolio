package server

import (
	"log/slog"
	"net/http"
	"time"
)

// ResponseWriter is a wrapper around http.ResponseWriter that stores the status code.
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the WriteHeader method to store the status code.
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// The logger middleware logs the request method and URL.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &ResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		slog.Info("request", slog.Int("status", writer.statusCode), slog.String("method", r.Method), slog.String("url", r.URL.Path), slog.Duration("duration", time.Since(start)))
	})
}
