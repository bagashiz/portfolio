package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/bagashiz/portfolio/internal/app/model"
	"github.com/bagashiz/portfolio/internal/app/store"
	"github.com/bagashiz/portfolio/web"
	"github.com/bagashiz/portfolio/web/template"
)

// The staticFiles function serves the static files such as CSS, JavaScript, and images.
func staticFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(web.StaticFiles).ServeHTTP(w, r)
	})
}

// The index function is the handler for the index page.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resumeTempl := template.Resume(
			model.Socials,
			model.Educations,
			model.Works,
			model.Volunteers,
			model.Awards,
			model.Certifications,
			model.SkillsFaIcons,
			model.Workflows,
		)

		_ = template.Index(resumeTempl).Render(r.Context(), w)
	})
}

// The resumePage function is the handler for the resume page.
func resumePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resumeTempl := template.Resume(
			model.Socials,
			model.Educations,
			model.Works,
			model.Volunteers,
			model.Awards,
			model.Certifications,
			model.SkillsFaIcons,
			model.Workflows,
		)

		_ = resumeTempl.Render(r.Context(), w)
	})
}

// The blogPage function is the handler for the blog page.
func blogPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = template.Blogs().Render(r.Context(), w)
	})
}

// The projectPage function is the handler for the project page.
func projectPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = template.Projects().Render(r.Context(), w)
	})
}

// The blogs function is the handler for fetching the blog posts and rendering them.
func blogs(cache store.Cache, blogUsername string) http.Handler {
	blogUrl := fmt.Sprintf("https://dev.to/api/articles?username=%s", blogUsername)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cachedBlogs, err := cache.GetCache(r.Context(), "blogs")
		if err == nil {
			slog.Info("cached blogs hit")

			blogs, err := decode[[]model.Blog](bytes.NewReader(cachedBlogs))
			if err != nil {
				errorFetch(w, r, err)
				return
			}

			_ = template.BlogCard(blogs).Render(r.Context(), w)
			return
		}

		slog.Info("cached blogs miss")

		req, err := http.NewRequest(http.MethodGet, blogUrl, nil)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if err != nil {
			errorFetch(w, r, err)
			return
		}
		defer resp.Body.Close()

		blogs, err := decode[[]model.Blog](resp.Body)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		err = cache.SetCache(r.Context(), "blogs", blogs, 1*time.Hour)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		_ = template.BlogCard(blogs).Render(r.Context(), w)
	})
}

// The projects function is the handler for fetching the pinned GitHub projects and rendering them.
func projects(cache store.Cache, githubUsername, githubAccessToken string) http.Handler {
	projectUrl := "https://api.github.com/graphql"
	query := fmt.Sprintf(`{
		user(login: "%s") {
			pinnedItems(first: 6, types: REPOSITORY) {
				nodes {
					... on Repository {
						name
						description
						url
						primaryLanguage {
							name
							color
						}
						stargazers {
							totalCount
						}
						forks {
							totalCount
						}
					}
				}
			}
		}
	}`, githubUsername)

	jsonData := map[string]string{
		"query": query,
	}
	reqBody, err := json.Marshal(jsonData)
	if err != nil {
		slog.Error(err.Error())
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cachedProjects, err := cache.GetCache(r.Context(), "projects")
		if err == nil {
			slog.Info("cached projects hit")
			projects, err := decode[model.Project](bytes.NewReader(cachedProjects))
			if err != nil {
				errorFetch(w, r, err)
				return
			}

			_ = template.ProjectCard(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
			return
		}

		slog.Info("cached projects miss")

		req, err := http.NewRequest(http.MethodPost, projectUrl, bytes.NewBuffer(reqBody))
		if err != nil {
			errorFetch(w, r, err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubAccessToken))

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if err != nil {
			errorFetch(w, r, err)
			return
		}
		defer resp.Body.Close()

		projects, err := decode[model.Project](resp.Body)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		err = cache.SetCache(r.Context(), "projects", projects, 1*time.Hour)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		_ = template.ProjectCard(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
	})
}

// The decode function decodes the JSON request/response body into the provided type.
func decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

// The errorFetch function logs the error and renders the error page when failed to fetch data.
func errorFetch(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error(err.Error())
	_ = template.ErrorFetchCard().Render(r.Context(), w)
}
