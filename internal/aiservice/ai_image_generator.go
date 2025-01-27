package aiservice

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
)

// TODO what to name this?
type AIService struct {
	projectID string
	location  string
	model     string
}

func New(projectID, location, model string) *AIService {
	return &AIService{
		projectID: projectID,
		location:  location,
		model:     model,
	}
}

// TODO where should we keep this?
func (service AIService) GenerateProfile(ctx *gin.Context) {

	// Parse and validate the request body
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

	imageBase64, err := callVertexAIWithSDK(ctx, "august-journey-434715-u0", "us-central1", "imagen-3.0-generate-001", prompt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate profile", "details": err.Error()})
		return
	}

	// Return the Base64-encoded image
	ctx.JSON(http.StatusOK, gin.H{
		"imageBase64": imageBase64,
	})
}

func callVertexAIWithSDK(ctx context.Context, projectID, location, modelID, prompt string) (string, error) {
	// Create the Vertex AI Prediction client
	client, err := aiplatform.NewPredictionClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create Vertex AI client: %w", err)
	}
	defer client.Close()

	// Define the model endpoint
	modelEndpoint := fmt.Sprintf("projects/%s/locations/%s/publishers/google/models/%s", projectID, location, modelID)

	// Build the request payload
	req := &aiplatformpb.PredictRequest{
		Endpoint: modelEndpoint,
		Instances: []*structpb.Value{
			structpb.NewStructValue(&structpb.Struct{
				Fields: map[string]*structpb.Value{
					"prompt": structpb.NewStringValue(prompt),
				},
			}),
		},
		Parameters: structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"temperature": structpb.NewNumberValue(0.8),
				"sampleCount": structpb.NewNumberValue(1),
			},
		}),
	}

	// Call the Vertex AI API
	resp, err := client.Predict(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to call Vertex AI: %w", err)
	}

	// Extract predictions from the response
	if len(resp.Predictions) == 0 {
		return "", fmt.Errorf("no predictions found in response")
	}

	// Extract the Base64-encoded image from the predictions
	prediction := resp.Predictions[0].GetStructValue().Fields
	imageBase64 := prediction["bytesBase64Encoded"].GetStringValue()

	if imageBase64 == "" {
		return "", fmt.Errorf("failed to extract image data from response")
	}

	return imageBase64, nil
}
