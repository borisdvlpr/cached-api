package handler

import (
	"io"
	"log"
	"net/http"
)

type ApiHandler struct {
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

func (h *ApiHandler) Get(w http.ResponseWriter, r *http.Request) {
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
}
