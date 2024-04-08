package server

import (
	"net"
	"net/http"
)

// The NewServer function creates a new http.Server type, configures the routes, and adds middleware.
func NewServer(getEnv func(string) string) *http.Server {
	mux := http.NewServeMux()
	addRoutes(mux, getEnv)

	var handler http.Handler = mux
	handler = logger(handler)

	addr := net.JoinHostPort(getEnv("APP_HOST"), getEnv("APP_PORT"))

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return server
}

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux, getEnv func(string) string) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", homePage())
	mux.Handle("GET /assets/", staticFiles())
	mux.Handle("GET /blogs/", blogPage())
	mux.Handle("GET /api/blogs/", blogs(getEnv))
}
