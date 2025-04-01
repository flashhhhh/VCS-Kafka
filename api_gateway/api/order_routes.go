package api

import (
	"gateway/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterProtectedRoutes(r *mux.Router, orderHandler *handler.OrderHandler) {
	r.HandleFunc("/create", orderHandler.CreateOrder).Methods("POST")
}