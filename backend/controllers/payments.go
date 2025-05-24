// controllers/payments.go
package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"backend/config"
	"backend/hasura"
	"backend/middleware"

	"github.com/google/uuid"
)

const (
	chapaURL     = "https://api.chapa.co/v1/transaction/initialize"
	minAmountETB = 5.0 // Minimum amount in ETB
	minAmountUSD = 0.5 // Minimum amount in USD
	minAmountEUR = 0.5 // Minimum amount in EUR
)

var supportedCurrencies = map[string]bool{
	"ETB": true,
	"USD": true,
	"EUR": true,
}

type PaymentRequest struct {
	RecipeID string  `json:"recipeId"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type ChapaResponse struct {
	Status string `json:"status"`
	Data   struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
	Message interface{} `json:"message"`
}

func validatePaymentRequest(req PaymentRequest) error {
	log.Printf("Validating payment request: %+v", req)
	// Check if required fields are present
	if req.RecipeID == "" {
		return fmt.Errorf("recipe ID is required")
	}

	if req.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	currency := strings.ToUpper(req.Currency)
	if !supportedCurrencies[currency] {
		return fmt.Errorf("unsupported currency: %s. Supported currencies are: ETB, USD, EUR", currency)
	}

	return validateAmount(req.Amount, currency)
}

func validateAmount(amount float64, currency string) error {
	switch currency {
	case "ETB":
		if amount < minAmountETB {
			return fmt.Errorf("minimum amount for ETB is %.2f", minAmountETB)
		}
	case "USD":
		if amount < minAmountUSD {
			return fmt.Errorf("minimum amount for USD is %.2f", minAmountUSD)
		}
	case "EUR":
		if amount < minAmountEUR {
			return fmt.Errorf("minimum amount for EUR is %.2f", minAmountEUR)
		}
	default:
		return fmt.Errorf("unsupported currency: %s", currency)
	}
	return nil
}

func PaymentInitHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT token
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Read and parse request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	log.Printf("Received payment request body: %s", string(body))

	var paymentReq PaymentRequest
	if err := json.Unmarshal(body, &paymentReq); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received payment request: %+v", paymentReq)

	// Validate the entire request
	if err := validatePaymentRequest(paymentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Normalize currency to uppercase
	currency := strings.ToUpper(paymentReq.Currency)

	// Generate transaction reference
	txRef := uuid.New().String()

	// Prepare Chapa request
	chapaPayload := map[string]interface{}{
		"amount":       paymentReq.Amount,
		"currency":     currency,
		"tx_ref":       txRef,
		"callback_url": os.Getenv("CHAPA_CALLBACK_URL"),
		"return_url":   os.Getenv("CHAPA_RETURN_URL"),
		"customization": map[string]interface{}{
			"title":       "Recipe Purchase",
			"description": "Payment for recipe purchase",
		},
	}

	chapaPayloadBytes, err := json.Marshal(chapaPayload)
	if err != nil {
		http.Error(w, "Error preparing Chapa request", http.StatusInternalServerError)
		return
	}

	// Get Chapa API key from config
	cfg := config.LoadConfig()
	chapaReq, err := http.NewRequest("POST", chapaURL, bytes.NewBuffer(chapaPayloadBytes))
	if err != nil {
		http.Error(w, "Error creating Chapa request", http.StatusInternalServerError)
		return
	}

	chapaReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.ChapaSecretKey))
	chapaReq.Header.Set("Content-Type", "application/json")

	// Log the request for debugging
	log.Printf("Making Chapa request with payload: %s", string(chapaPayloadBytes))

	chapaClient := &http.Client{}
	chapaResp, err := chapaClient.Do(chapaReq)
	if err != nil {
		http.Error(w, "Error making payment request to Chapa", http.StatusInternalServerError)
		return
	}
	defer chapaResp.Body.Close()

	// Log the response status and body for debugging
	bodyBytes, _ := io.ReadAll(chapaResp.Body)
	log.Printf("Chapa response status: %d", chapaResp.StatusCode)
	log.Printf("Chapa response body: %s", string(bodyBytes))

	if chapaResp.StatusCode != http.StatusOK {
		var errorResp ChapaResponse
		if err := json.Unmarshal(bodyBytes, &errorResp); err == nil {
			if errorMessage, ok := errorResp.Message.(map[string]interface{}); ok {
				// Format validation errors nicely
				var formattedErrors []string
				for field, errors := range errorMessage {
					if errorList, ok := errors.([]interface{}); ok {
						for _, err := range errorList {
							formattedErrors = append(formattedErrors, fmt.Sprintf("%s: %v", field, err))
						}
					}
				}
				if len(formattedErrors) > 0 {
					http.Error(w, fmt.Sprintf("Chapa validation errors: %s", strings.Join(formattedErrors, ", ")), http.StatusBadRequest)
					return
				}
			}
		}
		http.Error(w, fmt.Sprintf("Received non-200 response from Chapa: %s", string(bodyBytes)), http.StatusInternalServerError)
		return
	}

	// Create a new reader for the response body since we've already read it
	var chapaResponse ChapaResponse
	if err := json.Unmarshal(bodyBytes, &chapaResponse); err != nil {
		http.Error(w, "Error decoding Chapa response", http.StatusInternalServerError)
		return
	}

	// Create purchase record in Hasura
	client := hasura.NewClient(cfg)

	query := `
		mutation CreatePurchase($object: Purchases_insert_input!) {
			insert_Purchases_one(object: $object) {
				id
				user_id
				recipe_id
				chapa_tx_id
				amount
				created_at
			}
		}
	`

	variables := map[string]interface{}{
		"object": map[string]interface{}{
			"id":          uuid.New().String(),
			"user_id":     userID,
			"recipe_id":   paymentReq.RecipeID,
			"chapa_tx_id": txRef,
			"amount":      int(paymentReq.Amount),
			"created_at":  "now()",
		},
	}

	var response struct {
		InsertPurchasesOne struct {
			ID        string `json:"id"`
			UserID    string `json:"user_id"`
			RecipeID  string `json:"recipe_id"`
			ChapaTxID string `json:"chapa_tx_id"`
			Amount    int    `json:"amount"`
			CreatedAt string `json:"created_at"`
		} `json:"insert_Purchases_one"`
	}

	if err := client.Execute(r.Context(), query, variables, &response); err != nil {
		log.Printf("Error creating purchase record: %v", err)
		http.Error(w, "Error creating purchase record", http.StatusInternalServerError)
		return
	}

	// Return response in Hasura Action format
	json.NewEncoder(w).Encode(hasura.NewActionResponse("success", "Payment initiated", map[string]interface{}{
		"checkout_url": chapaResponse.Data.CheckoutURL,
		"tx_ref":       txRef,
	}))
}

func PaymentWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Verify event type
	event, _ := payload["event"].(string)
	if event != "charge.completed" {
		http.Error(w, "Invalid event", http.StatusBadRequest)
		return
	}

	// Extract transaction data
	data := payload["data"].(map[string]interface{})
	status := data["status"].(string)
	txRef := data["tx_ref"].(string)

	if status == "success" {
		// Update purchase status in Hasura
		cfg := config.LoadConfig()
		client := hasura.NewClient(cfg)

		query := `
			mutation UpdatePurchase($tx_ref: String!) {
				update_purchases(where: {chapa_tx_id: {_eq: $tx_ref}}, _set: {status: "completed"}) {
					affected_rows
				}
			}
		`

		variables := map[string]interface{}{
			"tx_ref": txRef,
		}

		var response struct {
			UpdatePurchases struct {
				AffectedRows int `json:"affected_rows"`
			} `json:"update_purchases"`
		}

		if err := client.Execute(r.Context(), query, variables, &response); err != nil {
			log.Printf("Error updating purchase status: %v", err)
			http.Error(w, "Error updating purchase status", http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(hasura.NewActionResponse("success", "Webhook processed", nil))
}
