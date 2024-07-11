package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/bagashiz/portfolio/internal/cache"
	"golang.org/x/sync/errgroup"
)

// Server wraps the http.Server type for extending functionality.
type Server struct {
	*http.Server
}

// The New function creates a new http.Server type, configures the routes, adds middleware, and starts the server gracefully.
func New(ctx context.Context, config map[string]string, cache cache.Cache) *Server {
	mux := http.NewServeMux()
	addRoutes(mux, cache, config)

	var handler http.Handler = mux

	addr := net.JoinHostPort(config["APP_HOST"], config["APP_PORT"])

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{server}
}

// Start starts the HTTP server in a separate goroutine and listens for
// the context cancellation signal to shut down the server gracefully.
func (s *Server) Start(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	errs.Go(func() error {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil
	})

	return errs.Wait()
}

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux, cache cache.Cache, config map[string]string) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", handle(index()))
	mux.Handle("GET /assets/", handle(staticFiles()))
	mux.Handle("GET /blogs/", handle(blogs(cache, config["DEV_USERNAME"])))
	mux.Handle("GET /projects/", handle(projects(cache, config["GITHUB_USERNAME"], config["GITHUB_ACCESS_TOKEN"])))
}
