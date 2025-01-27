package types

// Define a struct to bind the incoming JSON request body
type GenerateProfileRequest struct {
	Name       string `json:"name" binding:"required"`
	HairColor  string `json:"hairColor" binding:"required"`
	Profession string `json:"profession" binding:"required"`
	Build      string `json:"build" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
}
