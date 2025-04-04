package handler

import (
	"encoding/json"
	"fmt"
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

	orderInfo := map[string]any{
		"userID": userID,
		"order":  order,
	}

	// Send the order to Kafka
	id, err := h.service.SendOrder(orderInfo)
	if err != nil {
		http.Error(w, "Failed to send order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]any{
		"orderID": id,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOrder called")

	// Get user ID from the request context
	userID := r.Context().Value("userId").(string)
	log.Printf("User ID: %s\n", userID)

	// Parse the order ID from the request URL
	orderID := r.URL.Query().Get("orderID")
	if orderID == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	// Get the order information
	orderInfo := h.service.GetOrder(orderID)
	if orderInfo == nil {
		// Announce that the order has not been updated yet
		http.Error(w, "Order not found or not updated yet", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderInfo)
}