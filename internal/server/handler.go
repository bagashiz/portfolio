package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/bagashiz/portfolio/internal/cache"
	"github.com/bagashiz/portfolio/internal/model"
	"github.com/bagashiz/portfolio/internal/web"
	"github.com/bagashiz/portfolio/internal/web/components"
	"github.com/bagashiz/portfolio/internal/web/pages"
)

// The staticFiles function serves the static files such as CSS, JavaScript, and images.
func staticFiles() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		http.FileServerFS(web.StaticFiles).ServeHTTP(w, r)
		return nil
	}
}

// The index function is the handler for the index page.
func index() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if isHTMXRequest(r) {
			return components.Resume(
				model.Socials,
				model.Educations,
				model.Works,
				model.Volunteers,
				model.Awards,
				model.Certifications,
				model.SkillsFaIcons,
				model.Workflows,
			).Render(r.Context(), w)
		}

		return pages.Resume(
			model.Socials,
			model.Educations,
			model.Works,
			model.Volunteers,
			model.Awards,
			model.Certifications,
			model.SkillsFaIcons,
			model.Workflows,
		).Render(r.Context(), w)
	}
}

// The blogs function is the handler for fetching the blog posts and rendering them.
func blogs(cache cache.Cache, blogUsername string) handlerFunc {
	blogUrl := fmt.Sprintf("https://dev.to/api/articles?username=%s", blogUsername)

	return func(w http.ResponseWriter, r *http.Request) error {
		cachedBlogs, err := cache.GetCache(r.Context(), "blogs")
		if err == nil {
			slog.Info("cached blogs hit")

			blogs, err := decode[[]model.Blog](bytes.NewReader(cachedBlogs))
			if err != nil {
				return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
			}

			if isHTMXRequest(r) {
				return components.BlogCard(blogs).Render(r.Context(), w)
			}

			return pages.Blogs(blogs).Render(r.Context(), w)
		}

		slog.Info("cached blogs miss")

		req, err := http.NewRequest(http.MethodGet, blogUrl, nil)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			err := fmt.Errorf("failed to fetch blogs: %s", resp.Status)
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		blogs, err := decode[[]model.Blog](resp.Body)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		blogsJSON, err := json.Marshal(blogs)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		err = cache.SetCache(r.Context(), "blogs", blogsJSON)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		if isHTMXRequest(r) {
			return components.BlogCard(blogs).Render(r.Context(), w)
		}

		return pages.Blogs(blogs).Render(r.Context(), w)
	}
}

// The projects function is the handler for fetching the pinned GitHub projects and rendering them.
func projects(cache cache.Cache, githubUsername, githubAccessToken string) handlerFunc {
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

	return func(w http.ResponseWriter, r *http.Request) error {
		cachedProjects, err := cache.GetCache(r.Context(), "projects")
		if err == nil {
			slog.Info("cached projects hit")
			projects, err := decode[model.Project](bytes.NewReader(cachedProjects))
			if err != nil {
				return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
			}

			if isHTMXRequest(r) {
				return components.ProjectCard(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
			}

			return pages.Projects(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
		}

		slog.Info("cached projects miss")

		req, err := http.NewRequest(http.MethodPost, projectUrl, bytes.NewBuffer(reqBody))
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubAccessToken))

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			err := fmt.Errorf("failed to fetch projects: %s", resp.Status)
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		projects, err := decode[model.Project](resp.Body)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		projectsJSON, err := json.Marshal(projects)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		err = cache.SetCache(r.Context(), "projects", projectsJSON)
		if err != nil {
			return handlerError{statusCode: http.StatusInternalServerError, message: err.Error()}
		}

		if isHTMXRequest(r) {
			return components.ProjectCard(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
		}

		return pages.Projects(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
	}
}

// The decode function decodes the JSON request/response body into the provided type.
func decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

// isHTMXRequest checks request headers to determine if the request is an htmx request.
func isHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
