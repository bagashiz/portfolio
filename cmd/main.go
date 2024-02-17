package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	mux.HandleFunc("GET /hello/{name}/", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte("Hello, " + name + "!"))
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
