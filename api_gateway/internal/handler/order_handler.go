package handler

import (
	"encoding/json"
	"gateway/internal/service"
	"log"
	"net/http"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Get user ID from the request context
	userID := r.Context().Value("userId").(string)
	log.Printf("User ID: %s\n", userID)

	// Parse the order from the request body
	var order map[string]int
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid order format", http.StatusBadRequest)
		return
	}

	// Send the order to Kafka
	err := h.service.SendOrder(order)
	if err != nil {
		http.Error(w, "Failed to send order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}