package app

import (
	"net/http"

	"github.com/bagashiz/portfolio/internal/app/handler"
)

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/static/", handler.StaticFiles)
}
