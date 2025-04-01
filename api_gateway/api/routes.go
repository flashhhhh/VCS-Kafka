package api

import (
	"gateway/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, gatewayHandler *handler.GatewayHandler) {
	r.HandleFunc("/user/create", gatewayHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/login", gatewayHandler.LoginHandler).Methods("POST")
	r.HandleFunc("/user/getUserById", gatewayHandler.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/user/getAllUsers", gatewayHandler.GetAllUsersHandler).Methods("GET")
}

func RegisterProtectedRoutes(r *mux.Router, gatewayHandler *handler.GatewayHandler) {
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) { // Removed "/"
		userId, ok := r.Context().Value("userId").(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to retrieve user ID from context"))
			return
		}
		w.Write([]byte("Welcome to the Protected API! User ID: " + userId))
	}).Methods("GET")
}