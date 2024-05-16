package types

// TODO split into DTOs and models
// We use DTOs when creating or updating
// Then look into if there is a need for models containing all the data for a type. I.e a city that also has a list of all charcters related to it
type World struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cities      []string `json:"cities"` //Should this be a list of cities or just strings/foreign keys
	// Cities      []*City `json:"cities"`
	//TODO add image
}
