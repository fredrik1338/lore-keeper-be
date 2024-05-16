package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"
	"net/http"
)

func getWorld(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string

	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	world, err := db.GetWorld(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	response, err := json.Marshal(world)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addWorld(ctx context.Context, body []byte, db database.Database) (string, int) {

	var world types.World
	err := json.Unmarshal(body, &world)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.AddWorld(ctx, &world)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added world named %s", world.Name), http.StatusOK
}

func updateWorld(ctx context.Context, body []byte, db database.Database) (string, int) {

	var world types.World
	err := json.Unmarshal(body, &world)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.UpdateWorld(ctx, &world)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated world named %s", world.Name), http.StatusOK
}

func deleteWorld(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string

	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.DeleteWorld(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted world named %s", name), http.StatusOK
}
