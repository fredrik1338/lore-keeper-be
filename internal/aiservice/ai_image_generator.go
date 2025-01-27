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

	imageBase64, err := callVertexAI(ctx, service, prompt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate profile", "details": err.Error()})
		return
	}

	// Return the Base64-encoded image
	ctx.JSON(http.StatusOK, gin.H{
		"imageBase64": imageBase64,
	})
}

func callVertexAI(ctx context.Context, service AIService, prompt string) (string, error) {
	// Create the Vertex AI Prediction client
	client, err := aiplatform.NewPredictionClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create Vertex AI client: %w", err)
	}
	defer client.Close()

	// Define the endpoint for the model
	endpoint := fmt.Sprintf("projects/%s/locations/%s/publishers/google/models/%s", service.projectID, service.location, service.model)

	// Define the instances (input) for the request
	instances := []*structpb.Value{
		structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"prompt": structpb.NewStringValue(prompt),
			},
		}),
	}

	// Define the parameters for the request
	parameters := structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"temperature": structpb.NewNumberValue(0.8),
			"sampleCount": structpb.NewNumberValue(1),
		},
	})

	// Create the PredictRequest using the updated types
	req := &aiplatformpb.PredictRequest{
		Endpoint:   endpoint,
		Instances:  instances,
		Parameters: parameters,
	}

	// Call the Vertex AI Predict API
	resp, err := client.Predict(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to call Vertex AI: %w", err)
	}

	// Extract the Base64-encoded image from the response
	if len(resp.Predictions) == 0 {
		return "", fmt.Errorf("no predictions found in response")
	}

	// Extract the prediction details
	prediction := resp.Predictions[0].GetStructValue().Fields
	imageBase64 := prediction["bytesBase64Encoded"].GetStringValue()

	if imageBase64 == "" {
		return "", fmt.Errorf("failed to extract image data")
	}

	return imageBase64, nil
}
