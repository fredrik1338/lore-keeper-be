package types

// TODO split into DTOs and models
// We use DTOs when creating or updating
// Then look into if there is a need for models containing all the data for a type. I.e a city that also has a list of all charcters related to it
type City struct {
	Name              string   `json:"name"`
	Description       string   `json:"description,omitempty"`
	FoundingDate      string   `json:"foundingDate,omitempty"`
	NotableCharacters []string `json:"notableCharacters,omitempty"`
	Factions          []string `json:"factions,omitempty"`
	// NotableCharacters []*Character `json:"notableCharacters"`
	// Factions          []*Faction   `json:"factions"`
}
