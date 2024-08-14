package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api dbAPI) getCity(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "City name is required"})
		return
	}

	city, err := api.db.GetCity(ctx, name)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, city)
}

func (api dbAPI) listCities(ctx *gin.Context) {
	citys, err := api.db.ListCities(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list citys: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, citys)
}

// Handler for adding a new city
func (api *dbAPI) addCity(ctx *gin.Context) {
	var city types.City
	if err := ctx.ShouldBindJSON(&city); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	if err := api.db.AddCity(context.Background(), &city); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add city: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, city)
}

// Handler for updating an existing city
func (api *dbAPI) updateCity(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "City name is required"})
		return
	}

	var city types.City
	if err := ctx.ShouldBindJSON(&city); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	city.Name = name // Ensure the name in the URL matches the name in the payload
	if err := api.db.UpdateCity(context.Background(), &city); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update city: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, city)
}

// Handler for deleting a city
func (api *dbAPI) deleteCity(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "City name is required"})
		return
	}

	if err := api.db.DeleteCity(context.Background(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete city: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("City '%s' deleted successfully", name)})
}
