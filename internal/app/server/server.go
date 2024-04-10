package server

import (
	"net"
	"net/http"

	"github.com/bagashiz/portfolio/internal/app/store"
)

// The NewServer function creates a new http.Server type, configures the routes, and adds middleware.
func NewServer(config map[string]string, cache store.Cache) *http.Server {
	mux := http.NewServeMux()
	addRoutes(mux, cache, config)

	var handler http.Handler = mux
	handler = logger(handler)

	addr := net.JoinHostPort(config["APP_HOST"], config["APP_PORT"])

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return server
}

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux, cache store.Cache, config map[string]string) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", index())
	mux.Handle("GET /assets/", staticFiles())

	mux.Handle("GET /resume/", htmxHandler(resumePage()))
	mux.Handle("GET /blogs/", htmxHandler(blogPage()))
	mux.Handle("GET /projects/", htmxHandler(projectPage()))

	mux.Handle("GET /api/blogs/", blogs(cache, config["DEV_USERNAME"]))
	mux.Handle("GET /api/projects/", projects(cache, config["GITHUB_USERNAME"], config["GITHUB_ACCESS_TOKEN"]))
}
