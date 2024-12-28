package httputil

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse структура для відповіді з помилкою
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteError записує помилку у відповідь HTTP
func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
