package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) getFaction(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faction name is required"})
		return
	}

	faction, err := api.db.GetFaction(ctx, name)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, faction)
}

func (api API) listFactions(ctx *gin.Context) {
	factions, err := api.db.ListFactions(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list factions: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, factions)
}

// Handler for adding a new faction
func (api *API) addFaction(ctx *gin.Context) {
	var faction types.Faction
	if err := ctx.ShouldBindJSON(&faction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	if err := api.db.AddFaction(context.Background(), &faction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add faction: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, faction)
}

// Handler for updating an existing faction
func (api *API) updateFaction(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faction name is required"})
		return
	}

	var faction types.Faction
	if err := ctx.ShouldBindJSON(&faction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	faction.Name = name // Ensure the name in the URL matches the name in the payload
	if err := api.db.UpdateFaction(context.Background(), &faction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update faction: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, faction)
}

// Handler for deleting a faction
func (api *API) deleteFaction(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faction name is required"})
		return
	}

	if err := api.db.DeleteFaction(context.Background(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete faction: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Faction '%s' deleted successfully", name)})
}
