package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"serverapp/internal/database"
	"serverapp/internal/types"
)

func getWorld(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	world, err := database.GetWorld(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	response, err := json.Marshal(world)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addWorld(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var world types.World
	err := json.Unmarshal(body, &world)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.AddWorld(ctx, db, &world)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added world named %s", world.Name), http.StatusOK
}

func updateWorld(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var world types.World
	err := json.Unmarshal(body, &world)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.UpdateWorld(ctx, db, &world)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated world named %s", world.Name), http.StatusOK
}

func deleteWorld(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.DeleteWorld(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted world named %s", name), http.StatusOK
}
