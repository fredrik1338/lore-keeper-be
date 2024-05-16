package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"
	"net/http"
)

func getFaction(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string

	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	faction, err := db.GetFaction(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	response, err := json.Marshal(faction)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addFaction(ctx context.Context, body []byte, db database.Database) (string, int) {

	var faction types.Faction
	err := json.Unmarshal(body, &faction)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.AddFaction(ctx, &faction)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added faction named %s", faction.Name), http.StatusOK
}

func updateFaction(ctx context.Context, body []byte, db database.Database) (string, int) {

	var faction types.Faction
	err := json.Unmarshal(body, &faction)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.UpdateFaction(ctx, &faction)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated faction named %s", faction.Name), http.StatusOK
}

func deleteFaction(ctx context.Context, body []byte, db database.Database) (string, int) {
	var name string

	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = db.DeleteFaction(ctx, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted faction named %s", name), http.StatusOK
}
