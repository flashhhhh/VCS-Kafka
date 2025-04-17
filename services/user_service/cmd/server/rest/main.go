package main

import (
	"log"
	"net/http"
	"os"
	"user_service/api"
	database "user_service/infrastructure/postgres"
	"user_service/infrastructure/redis"
	rest_handler "user_service/internal/handler/rest"
	"user_service/internal/repository"
	"user_service/internal/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/flashhhhh/pkg/env"
)

func checkCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Vary", "Origin")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var environment string

	args := os.Args
	if len(args) > 1 {
		environment = args[1]
		println("Running in environment: ", environment)
	} else {
		println("No environment specified, defaulting to local")
		environment = "local"
	}

	// Load environment variables from the specified file
	envFile := "config/user." + environment + ".env"
	env.LoadEnv(envFile)
	log.Println("Environment variables loaded from file:", envFile)

	// Connect to the database
	log.Println("Connecting to the database...")
	dsn := "host=" + env.GetEnv("USER_DB_HOST", "localhost") + " user=" + env.GetEnv("USER_DB_USER", "root") + " password=" + env.GetEnv("USER_DB_PASSWORD", "") + " dbname=" + env.GetEnv("USER_DB_NAME", "user_service") + " port=" + env.GetEnv("USER_DB_PORT", "5432") + " sslmode=disable"
	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Migrate the database
	log.Println("Migrating the database...")
	err = database.MigrateDB(db)
	if err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
	
	// Connect to Redis
	log.Println("Connecting to Redis...")
	redisAddr := env.GetEnv("REDIS_HOST", "localhost") + ":" + env.GetEnv("REDIS_PORT", "6379")
	redisClient := redis.NewRedisClient(redisAddr)

	// Initialize the repository, service, and handler
	userRepo := repository.NewUserRepository(db, redisClient)
	userService := service.NewUserService(userRepo)
	userHandler := rest_handler.NewUserHandler(userService)

	r := mux.NewRouter()
	api.RegisterRoutes(r, userHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins, change this for security
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	serverPort := env.GetEnv("REST_USER_SERVER_PORT", "8001")

	// Open http server
	log.Println("Starting HTTP server at port", env.GetEnv("REST_USER_SERVER_PORT", "8001"))
	log.Fatal(http.ListenAndServe(":"+serverPort, checkCORS(corsHandler)))
}