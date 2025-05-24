// utils/cloudinary.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type CloudinaryResponse struct {
	SecureURL string `json:"secure_url"`
	PublicID  string `json:"public_id"`
	Error     struct {
		Message string `json:"message"`
	} `json:"error"`
}

func UploadToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return "", fmt.Errorf("missing Cloudinary credentials")
	}

	url := "https://api.cloudinary.com/v1_1/" + cloudName + "/image/upload"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file part
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.SetBasicAuth(apiKey, apiSecret)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body for debugging
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	log.Printf("Cloudinary response: %s", string(respBody))

	var cloudResp CloudinaryResponse
	if err := json.Unmarshal(respBody, &cloudResp); err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	if cloudResp.Error.Message != "" {
		return "", fmt.Errorf("cloudinary error: %s", cloudResp.Error.Message)
	}

	if cloudResp.SecureURL == "" {
		return "", fmt.Errorf("no secure URL in response")
	}

	return cloudResp.SecureURL, nil
}
