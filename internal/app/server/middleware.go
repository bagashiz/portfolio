package server

import (
	"log/slog"
	"net/http"
	"time"
)

// The logger middleware logs the request method and URL.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("Request", slog.String("method", r.Method), slog.String("url", r.URL.Path), slog.Duration("duration", time.Since(start)))
	})
}

func htmxHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
