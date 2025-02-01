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

type Config struct {
	Host              string
	Port              string
	DevUsername       string
	GithubUsername    string
	GithubAccessToken string
}

// The New function creates a new http.Server type, configures the routes, adds middleware, and starts the server gracefully.
func New(ctx context.Context, cache cache.Cache, cfg Config) *Server {
	mux := http.NewServeMux()
	addRoutes(mux, cache, cfg)

	var handler http.Handler = mux

	addr := net.JoinHostPort(cfg.Host, cfg.Port)

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
func addRoutes(mux *http.ServeMux, cache cache.Cache, cfg Config) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", handle(index()))
	mux.Handle("GET /assets/", handle(staticFiles()))
	mux.Handle("GET /blogs/", handle(blogs(cache, cfg.DevUsername)))
	mux.Handle("GET /projects/", handle(projects(cache, cfg.GithubUsername, cfg.GithubAccessToken)))
}
