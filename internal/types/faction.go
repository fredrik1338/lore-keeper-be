package types

// TODO split into DTOs and models
// We use DTOs when creating or updating
// Then look into if there is a need for models containing all the data for a type. I.e a city that also has a list of all charcters related to it
type Faction struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	FoundingDate      string   `json:"foundingDate"`
	NotableCharacters []string `json:"notableCharacters"`
	// NotableCharacters []*Character `json:"notableCharacters"`
	// Leader *Character `json:"leader"` TODO is it really necessary ot have a leader?
}
