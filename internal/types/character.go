package types

type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	World       string `json:"world"` //TODO make into *World once the DB is updated
}
