package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

func (h *ApiHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%s", id))
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
}
