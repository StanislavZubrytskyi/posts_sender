package handlers

import (
	"encoding/json"
	"net/http"
	"posts_sender/internal/service"
	"posts_sender/pkg/httputil"
)

// PostHandler обробляє HTTP запити пов'язані з постами
type PostHandler struct {
	service service.PostService
}

// NewPostHandler створює новий екземпляр обробника постів
func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

// GetPosts обробляє GET запит для отримання всіх постів
func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetPosts()
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// SyncPosts обробляє POST запит для синхронізації постів
func (h *PostHandler) SyncPosts(w http.ResponseWriter, r *http.Request) {
	if err := h.service.SyncPosts(); err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "Failed to sync posts")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Posts synchronized successfully"})
}
