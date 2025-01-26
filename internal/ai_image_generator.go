package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api dbAPI) generateProfile(ctx *gin.Context) {
	// Define a struct to bind the incoming JSON request body
	type GenerateProfileRequest struct {
		Name       string `json:"name" binding:"required"`
		HairColor  string `json:"hairColor" binding:"required"`
		Profession string `json:"profession" binding:"required"`
		Build      string `json:"build" binding:"required"`
		Gender     string `json:"gender" binding:"required"`
	}

	// Parse and validate the request body
	var req GenerateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Construct the prompt for the AI service
	prompt := fmt.Sprintf(
		"Create a character portrait for a character in a grimdark setting. Keywords: grim, realism, gritty, hopeful, heroism, underground, futuristic. Features: Name: %s, Hair Color: %s, Profession: %s, Build: %s, Gender: %s.",
		req.Name, req.HairColor, req.Profession, req.Build, req.Gender,
	)

	// Construct the AI API request
	requestBody := map[string]interface{}{
		"instances": []map[string]interface{}{
			{"prompt": prompt},
		},
		"parameters": map[string]interface{}{
			"sampleCount": 1,
			"temperature": 0.8,
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize request", "details": err.Error()})
		return
	}

	// Call the GCP AI endpoint
	apiURL := "https://us-central1-aiplatform.googleapis.com/v1/projects/august-journey-434715-u0/locations/us-central1/publishers/google/models/imagen-3.0-generate-001:predict"
	reqGCP, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request", "details": err.Error()})
		return
	}

	reqGCP.Header.Set("Authorization", "Bearer "+getAccessToken())
	reqGCP.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(reqGCP)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call AI service", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "AI service responded with an error", "details": string(body)})
		return
	}

	var aiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&aiResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse AI service response", "details": err.Error()})
		return
	}

	// Extract the Base64-encoded image from the response
	predictions, ok := aiResponse["predictions"].([]interface{})
	if !ok || len(predictions) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No predictions in AI service response"})
		return
	}

	imageBase64, ok := predictions[0].(map[string]interface{})["bytesBase64Encoded"].(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract image data"})
		return
	}

	// Return the Base64-encoded image
	ctx.JSON(http.StatusOK, gin.H{
		"imageBase64": imageBase64,
	})
}

func getAccessToken() string {
	cmd := exec.Command("gcloud", "auth", "print-access-token")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}
	return strings.TrimSpace(string(out))
}
