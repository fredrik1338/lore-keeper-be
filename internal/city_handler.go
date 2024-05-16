package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"
	"net/http"
)

func getCity(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	city, err := database.GetCity(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	response, err := json.Marshal(city)
	if err != nil {
		return fmt.Sprintf("failed to marshal character: %s", err.Error()), http.StatusInternalServerError
	}

	return string(response), http.StatusOK
}

func addCity(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var city types.City
	err := json.Unmarshal(body, &city)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.AddCity(ctx, db, &city)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Added City named %s", city.Name), http.StatusOK
}

func updateCity(ctx context.Context, body []byte) (string, int) {
	db := getDBConn()
	var city types.City
	err := json.Unmarshal(body, &city)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.UpdateCity(ctx, db, &city)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Updated city named %s", city.Name), http.StatusOK
}

func deleteCity(ctx context.Context, body []byte) (string, int) {
	var name string
	db := getDBConn()
	err := json.Unmarshal(body, &name)
	if err != nil {
		return fmt.Sprintf("failed to unmarshal body: %s", err.Error()), http.StatusBadRequest
	}

	err = database.DeleteCity(ctx, db, name)
	if err != nil {
		return fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError
	}

	return fmt.Sprintf("Deleted City named %s", name), http.StatusOK
}
