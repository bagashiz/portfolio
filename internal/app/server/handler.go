package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bagashiz/portfolio/internal/app/model"
	"github.com/bagashiz/portfolio/web"
	"github.com/bagashiz/portfolio/web/template"
)

// The staticFiles function serves the static files such as CSS, JavaScript, and images.
func staticFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(web.StaticFiles).ServeHTTP(w, r)
	})
}

// The home function is the handler for the home page.
func homePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		_ = template.Base("Bagashiz | Resume", templ).Render(r.Context(), w)
	})
}

// The blog function is the handler for the blog page.
func blogPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ := template.Blog()

		_ = template.Base("Bagashiz | Blog", templ).Render(r.Context(), w)
	})
}

// The blogs function is the handler for fetching the blog posts and rendering them.
func blogs(getEnv func(string) string) http.Handler {
	blogUrl := getEnv("DEV_API_URL")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(blogUrl)
		if err != nil {
			slog.Error(err.Error())
			_ = template.ErrorFetchCard().Render(r.Context(), w)
			return
		}
		defer resp.Body.Close()

		var blogs []model.Blog
		err = json.NewDecoder(resp.Body).Decode(&blogs)
		if err != nil {
			slog.Error(err.Error())
			_ = template.ErrorFetchCard().Render(r.Context(), w)
			return
		}

		_ = template.BlogCard(blogs).Render(r.Context(), w)
	})
}
