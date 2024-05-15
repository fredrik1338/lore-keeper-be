package types

type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	Home        string `json:"home"` //TODO make into *World once the DB is updated
}
