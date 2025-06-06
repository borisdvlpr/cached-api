package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"cachedapi/internal/handler"
	"cachedapi/internal/service"
	"cachedapi/pkg/cache"
	"cachedapi/pkg/config"
)

func Run(cfg *config.Config) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cacheClient, err := cache.NewClient(cfg)
	if err != nil {
		log.Printf("Valkey error: %v", err)
	}
	defer cacheClient.Close()

	apiService := service.NewApiService(cacheClient)
	apiHandler := handler.NewApiHandler(apiService)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK."))
	})

	r.Get("/todo/{id}", apiHandler.GetTodo)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
