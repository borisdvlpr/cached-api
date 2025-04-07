package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"cachedapi/internal/service"
)

type ApiHandler struct {
	svc *service.ApiService
}

func NewApiHandler(svc *service.ApiService) *ApiHandler {
	return &ApiHandler{
		svc: svc,
	}
}

func (h *ApiHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	cacheKey := fmt.Sprintf("todo:%s", id)

	cachedData, err := h.svc.GetCache(ctx, cacheKey)
	if err == nil {
		log.Printf("Cache hit for todo ID: %s", id)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		w.Write(cachedData)
		return
	}

	log.Printf("Todo ID %s: %v", id, err)

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%s", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
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

	if err := h.svc.SetCache(ctx, cacheKey, body); err != nil {
		log.Printf("Error caching response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Cache", "MISS")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
