package types

type City struct {
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	FoundingDate      string       `json:"foundingDate"`
	NotableCharacters []*Character `json:"notableCharacters"`
	Factions          []*Faction   `json:"factions"`
}
