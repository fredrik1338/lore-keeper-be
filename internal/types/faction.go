package types

type Faction struct {
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	FoundingDate      string       `json:"foundingDate"`
	NotableCharacters []*Character `json:"notableCharacters"`
	Leader            *Character   `json:"leader"`
}
