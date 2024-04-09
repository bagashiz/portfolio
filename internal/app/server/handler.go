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
	"github.com/bagashiz/portfolio/web"
	"github.com/bagashiz/portfolio/web/template"
)

// The staticFiles function serves the static files such as CSS, JavaScript, and images.
func staticFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(web.StaticFiles).ServeHTTP(w, r)
	})
}

// The homePage function is the handler for the home page.
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

// The blogPage function is the handler for the blog page.
func blogPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ := template.Blogs()

		_ = template.Base("Bagashiz | Blog", templ).Render(r.Context(), w)
	})
}

// The blogs function is the handler for fetching the blog posts and rendering them.
func blogs(blogUsername string) http.Handler {
	blogUrl := fmt.Sprintf("https://dev.to/api/articles?username=%s", blogUsername)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(http.MethodGet, blogUrl, nil)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if resp.StatusCode != http.StatusOK && err != nil {
			errorFetch(w, r, err)
			return
		}
		defer resp.Body.Close()

		blogs, err := decode[[]model.Blog](resp.Body)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		_ = template.BlogCard(blogs).Render(r.Context(), w)
	})
}

// The projectPage function is the handler for the project page.
func projectPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ := template.Projects()

		_ = template.Base("Bagashiz | Projects", templ).Render(r.Context(), w)
	})
}

// The projects function is the handler for fetching the pinned GitHub projects and rendering them.
func projects(githubUsername, githubAccessToken string) http.Handler {
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
		req, err := http.NewRequest(http.MethodPost, projectUrl, bytes.NewBuffer(reqBody))
		if err != nil {
			errorFetch(w, r, err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubAccessToken))

		client := &http.Client{Timeout: 5 * time.Second}

		resp, err := client.Do(req)
		if resp.StatusCode != http.StatusOK && err != nil {
			errorFetch(w, r, err)
			return
		}
		defer resp.Body.Close()

		projects, err := decode[model.Project](resp.Body)
		if err != nil {
			errorFetch(w, r, err)
			return
		}

		_ = template.ProjectCard(projects.Data.User.PinnedItems.Nodes).Render(r.Context(), w)
	})
}

// The errorFetch function logs the error and renders the error page when failed to fetch data.
func errorFetch(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error(err.Error())
	_ = template.ErrorFetchCard().Render(r.Context(), w)
}

// The decode function decodes the JSON request/response body into the provided type.
func decode[T any](b io.ReadCloser) (T, error) {
	var v T
	if err := json.NewDecoder(b).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}
