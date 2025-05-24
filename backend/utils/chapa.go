// utils/chapa.go
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type ChapaInitRequest struct {
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	TxRef       string `json:"tx_ref"`
	ReturnURL   string `json:"return_url"`
	CallbackURL string `json:"callback_url"`
}

func InitiateChapaPayment(data ChapaInitRequest) (map[string]interface{}, error) {
	apiKey := os.Getenv("CHAPA_API_KEY")

	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "https://api.chapa.co/v1/transaction/initialize", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var chapaResp map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&chapaResp)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to initiate payment")
	}

	return chapaResp, nil
}
