package app

import (
	"net/http"

	"github.com/bagashiz/portfolio/internal/app/middleware"
)

// The NewServer function creates a new http.Server type, configures the routes, and adds middleware.
func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	handler = middleware.Logger(handler)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return server
}
