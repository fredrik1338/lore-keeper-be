package types

type World struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cities      []*City `json:"cities"` //Should this be a list of cities or just strings/foreign keys
	//TODO add image
}
