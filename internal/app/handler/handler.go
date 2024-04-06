package handler

import (
	"net/http"

	"github.com/bagashiz/portfolio/internal/app/model"
	"github.com/bagashiz/portfolio/web"
	"github.com/bagashiz/portfolio/web/template"
)

// The StaticFiles function serves the static files such as CSS, JavaScript, and images.
func StaticFiles(w http.ResponseWriter, r *http.Request) {
	http.FileServerFS(web.StaticFiles).ServeHTTP(w, r)
}

// The Home function is the handler for the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	templ := template.Home(
		model.Socials,
		model.Educations,
		model.Works,
		model.Volunteers,
		model.Awards,
		model.Certifications,
		model.SkillsFaIcons,
		model.Workflows,
	)

	_ = template.Layout("Bagashiz | Resume", templ).Render(r.Context(), w)
}
