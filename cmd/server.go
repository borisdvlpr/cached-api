package cmd

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK."))
	})

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			log.Printf("Error fetching API: %v", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found."))
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading response body."))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(body))
	})

	http.ListenAndServe(":3000", r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
