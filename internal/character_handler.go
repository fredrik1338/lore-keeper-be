package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"serverapp/internal/database"
	"serverapp/internal/types"
)

func getCharacter(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	character, err := database.GetCharacter(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}
	response, err := json.Marshal(character)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addCharacter(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var person types.Character
	err := json.Unmarshal(body, &person)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.AddCharacter(ctx, db, &person)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added character named %s", person.Name), http.StatusOK
}

func updateCharacter(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var person types.Character
	err := json.Unmarshal(body, &person)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.UpdateCharacter(ctx, db, &person)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated character named %s", person.Name), http.StatusOK
}

func deleteCharacter(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.DeleteCharacter(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted character named %s", name), http.StatusOK
}
