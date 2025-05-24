package controllers

import (
	"backend/config"
	"backend/hasura"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	// Create user via Hasura
	cfg := config.LoadConfig()
	client := hasura.NewClient(cfg)

	query := `
		mutation CreateUser($id: uuid!, $username: String!, $email: String!, $password: String!) {
			insert_Users_one(object: {
				id: $id,
				username: $username,
				email: $email,
				password: $password
			}) {
				id
				username
				email
			}
		}
	`

	variables := map[string]interface{}{
		"id":       uuid.New().String(),
		"username": input.Username,
		"email":    input.Email,
		"password": string(hashedPassword),
	}

	var response struct {
		InsertUsersOne struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"insert_Users_one"`
	}

	if err := client.Execute(r.Context(), query, variables, &response); err != nil {
		log.Printf("Error creating user: %v", err)
		log.Printf("Query: %s", query)
		log.Printf("Variables: %+v", variables)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Log successful registration
	log.Printf("User registered successfully: %+v", response.InsertUsersOne)

	json.NewEncoder(w).Encode(hasura.NewActionResponse("success", "User registered successfully", response.InsertUsersOne))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Get user from Hasura
	cfg := config.LoadConfig()
	client := hasura.NewClient(cfg)

	query := `
		query GetUser($email: String!) {
			Users(where: {email: {_eq: $email}}) {
				id
				username
				email
				password
			}
		}
	`

	variables := map[string]interface{}{
		"email": input.Email,
	}

	var response struct {
		Users []User `json:"users"`
	}

	if err := client.Execute(r.Context(), query, variables, &response); err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	if len(response.Users) == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	user := response.Users[0]

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hasura.NewActionResponse("success", "Login successful", map[string]string{
		"token": tokenString,
	}))
}
