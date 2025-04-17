package api

import (
	rest_handler "user_service/internal/handler/rest"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, gatewayHandler *rest_handler.UserHandler) {
	r.HandleFunc("/user/create", gatewayHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/login", gatewayHandler.Login).Methods("POST")
	r.HandleFunc("/user/getUserById", gatewayHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/user/getAllUsers", gatewayHandler.GetAllUsers).Methods("GET")
}