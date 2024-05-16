package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"
	"net/http"
)

func getCharacter(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	character, err := db.GetCharacter(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}
	response, err := json.Marshal(character)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addCharacter(ctx context.Context, body []byte, db database.Database) (string, int) {

	var person types.Character
	err := json.Unmarshal(body, &person)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.AddCharacter(ctx, &person)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added character named %s", person.Name), http.StatusOK
}

func updateCharacter(ctx context.Context, body []byte, db database.Database) (string, int) {
	var person types.Character
	err := json.Unmarshal(body, &person)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.UpdateCharacter(ctx, &person)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated character named %s", person.Name), http.StatusOK
}

func deleteCharacter(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string

	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.DeleteCharacter(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted character named %s", name), http.StatusOK
}
