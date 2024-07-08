package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/bagashiz/portfolio/internal/cache"
)

// The Start function creates a new http.Server type, configures the routes, adds middleware, and starts the server gracefully.
func Start(ctx context.Context, config map[string]string, cache cache.Cache) {
	mux := http.NewServeMux()
	addRoutes(mux, cache, config)

	var handler http.Handler = mux
	handler = logger(handler)

	addr := net.JoinHostPort(config["APP_HOST"], config["APP_PORT"])

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error listening and serving: %s\n", err)
		}
	}()

	log.Printf("listening on %s\n", server.Addr)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down server: %s\n", err)
			return
		}

		log.Println("server shut down gracefully")
	}()

	wg.Wait()
}

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux, cache cache.Cache, config map[string]string) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", index())
	mux.Handle("GET /assets/", staticFiles())
	mux.Handle("GET /blogs/", blogs(cache, config["DEV_USERNAME"]))
	mux.Handle("GET /projects/", projects(cache, config["GITHUB_USERNAME"], config["GITHUB_ACCESS_TOKEN"]))
}
