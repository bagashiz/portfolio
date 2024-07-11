package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/bagashiz/portfolio/internal/web/components"
)

// handlerFunc is a function that handles an HTTP request and returns an error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// handlerError is a custom error for HTTP handlers that includes a status code.
type handlerError struct {
	statusCode int
	message    string
}

// Error returns the error message for the handlerError type.
func (h handlerError) Error() string {
	return h.message
}

// responseWriter extends the http.ResponseWriter type to store the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the WriteHeader method to store the status code.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// handle wraps a handlerFunc type as an http.Handler, switch to custom ResponseWriter,
// handles errors, and logs request information.
func handle(h handlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		if err := h(writer, r); err != nil {
			statusCode := http.StatusInternalServerError
			message := http.StatusText(statusCode)

			switch err {
			case err.(handlerError):
				handlerErr := err.(handlerError)
				statusCode = handlerErr.statusCode
				message = handlerErr.Error()
				_ = components.ErrorFetchCard().Render(r.Context(), w)
			default:
				http.Error(writer, message, statusCode)
			}

			slog.Error("request failed",
				slog.Int("status", statusCode),
				slog.String("error", message),
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
			)
			return
		}

		slog.Info("request served",
			slog.Int("status", writer.statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", time.Since(start)),
		)
	})
}
