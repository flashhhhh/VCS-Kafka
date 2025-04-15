package api

import (
	"gateway/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterProtectedRoutes(r *mux.Router, orderHandler *handler.OrderHandler) {
	r.HandleFunc("/create", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/getOrder", orderHandler.GetOrder).Methods("GET")
}