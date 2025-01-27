package aiservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"lore-keeper-be/internal/types"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
)

// TODO what to name this?
type AIService struct {
	mode   string
	apiURL string
}

func New(projectID, location, model, mode string) *AIService {
	return &AIService{
		mode: mode,
		apiURL: fmt.Sprintf(
			"https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:predict",
			location, projectID, location, model),
	}
}

// TODO where should we keep this?
func (service AIService) GenerateProfile(ctx *gin.Context) {
	var req types.GenerateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Construct the prompt for the AI service
	prompt := fmt.Sprintf(
		"Create a character portrait for a character in a grimdark setting. Keywords: grim, realism, gritty, hopeful, heroism, underground, futuristic. Features: Name: %s, Hair Color: %s, Profession: %s, Build: %s, Gender: %s.",
		req.Name, req.HairColor, req.Profession, req.Build, req.Gender,
	)

	imageBase64, err := callVertexAIWithREST(context.Background(), service.apiURL, prompt, service.mode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate profile", "details": err.Error()})
		return
	}

	// Return the Base64-encoded image
	ctx.JSON(http.StatusOK, gin.H{
		"imageBase64": imageBase64,
	})
}

func callVertexAIWithREST(ctx context.Context, apiURL, prompt, mode string) (string, error) {

	// Construct the request payload
	requestBody := map[string]interface{}{
		"instances": []map[string]interface{}{
			{"prompt": prompt},
		},
		"parameters": map[string]interface{}{
			"temperature": 0.8,
			"sampleCount": 1,
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to serialize request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add Authorization header with ADC token
	token := getAccessToken(mode)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Vertex AI: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Vertex AI API returned error: %s", string(body))
	}

	// Parse the response
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to parse Vertex AI response: %w", err)
	}

	// Extract the Base64-encoded image
	predictions, ok := response["predictions"].([]interface{})
	if !ok || len(predictions) == 0 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("no predictions found in Vertex AI response. Response: %s", string(body))
	}

	imageBase64, ok := predictions[0].(map[string]interface{})["bytesBase64Encoded"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract image data")
	}

	return imageBase64, nil
}

// Helper function to get the access token
func getAccessToken(mode string) string {

	if mode == "dev" {
		log.Println("Using hardcoded access token for local development")
		token := os.Getenv("ACCESS_TOKEN")
		if token == "" {
			log.Printf("ACCESS_TOKEN environment variable not set")
		}
		return token
	} else {
		ctx := context.Background()

		// Create the default token source (automatically tied to the service account in Cloud Run)
		tokenSource, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/cloud-platform")
		if err != nil {
			log.Printf("Failed to create token source: %v", err)
		}

		// Retrieve the token
		token, err := tokenSource.Token()
		if err != nil {
			log.Printf("Failed to fetch access token: %v", err)
		}

		return token.AccessToken
	}

}
