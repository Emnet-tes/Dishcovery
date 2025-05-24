package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"backend/hasura"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type UploadImagesRequest struct {
	Input struct {
		UserInput struct {
			Files []string `json:"files"`
		} `json:"userInput"`
	} `json:"input"`
}

type UploadImagesResponse struct {
	ImageURLs []string `json:"imageUrls"`
}

func UploadImagesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadImagesHandler hit")

	// Check for Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, `{"error": "Authorization header required"}`, http.StatusUnauthorized)
		return
	}

	// Validate Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, `{"error": "Invalid authorization header format"}`, http.StatusUnauthorized)
		return
	}

	// Parse and validate JWT token
	tokenString := parts[1]
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
		return
	}

	// Get user claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, `{"error": "Invalid token claims"}`, http.StatusUnauthorized)
		return
	}

	// Get user ID and role from claims
	userId, ok := claims["sub"].(string)
	if !ok {
		http.Error(w, `{"error": "Invalid user ID in token"}`, http.StatusUnauthorized)
		return
	}

	// Log user info
	log.Printf("Upload request from user: %s", userId)

	err = godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env: %v", err)
		http.Error(w, `{"error": "Failed to load .env"}`, http.StatusInternalServerError)
		return
	}

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Printf("Failed to initialize Cloudinary: %v", err)
		http.Error(w, `{"error": "Failed to initialize Cloudinary"}`, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Only POST allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req UploadImagesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON payload"}`, http.StatusBadRequest)
		return
	}

	files := req.Input.UserInput.Files
	if len(files) == 0 {
		http.Error(w, `{"error": "No image data provided"}`, http.StatusBadRequest)
		return
	}

	var urls []string

	for i, file := range files {
		// Clean base64 if it includes a prefix like "data:image/png;base64,..."
		if commaIdx := bytes.IndexByte([]byte(file), ','); commaIdx != -1 {
			file = file[commaIdx+1:]
		}

		imageData, err := base64.StdEncoding.DecodeString(file)
		if err != nil {
			log.Printf("Failed to decode base64 at index %d: %v", i, err)
			http.Error(w, `{"error": "Invalid base64 image"}`, http.StatusBadRequest)
			return
		}

		// Create a folder structure that includes user ID
		folderPath := fmt.Sprintf("RecipeImages/%s", userId)

		uploadResp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(imageData), uploader.UploadParams{
			Folder: folderPath,
		})
		if err != nil {
			log.Printf("Upload failed at index %d: %v", i, err)
			http.Error(w, `{"error": "Failed to upload image"}`, http.StatusInternalServerError)
			return
		}
		log.Printf("Upload response: %+v", uploadResp.Error)
		log.Printf("Uploaded image %d to: %s", i, uploadResp.SecureURL)
		urls = append(urls, uploadResp.SecureURL)
	}

	// Return response in Hasura Action format
	json.NewEncoder(w).Encode(hasura.NewActionResponse("success", "Images uploaded successfully", map[string]interface{}{
		"urls": urls,
	}))
}
