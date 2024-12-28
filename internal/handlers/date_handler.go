package handlers

import (
	"encoding/json"
	"net/http"
	"posts_sender/internal/models"
	"posts_sender/internal/service"
	"posts_sender/pkg/httputil"
)

// DateHandler обробляє HTTP запити пов'язані з датами
type DateHandler struct {
	service service.DateService
}

// NewDateHandler створює новий екземпляр обробника дат
func NewDateHandler(service service.DateService) *DateHandler {
	return &DateHandler{
		service: service,
	}
}

// CalculateTimeRemaining обробляє POST запит для обчислення залишку часу
func (h *DateHandler) CalculateTimeRemaining(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	var request models.DateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	timeRemaining := h.service.CalculateTimeRemaining(request.TargetDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeRemaining)
}
