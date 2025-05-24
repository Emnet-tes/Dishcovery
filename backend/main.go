// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"backend/controllers"
	"backend/middleware"
)

func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// Public routes (no auth required)
	r.HandleFunc("/auth/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", controllers.LoginHandler).Methods("POST")

	// Protected routes (auth required)
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Upload
	protected.HandleFunc("/upload/recipe-images", controllers.UploadImagesHandler).Methods("POST")

	// Payments
	protected.HandleFunc("/payments/initiate", controllers.PaymentInitHandler).Methods("POST")
	protected.HandleFunc("/payments/webhook", controllers.PaymentWebhookHandler).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
