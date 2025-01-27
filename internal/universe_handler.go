package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) getUniverse(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Universe name is required"})
		return
	}

	universe, err := api.db.GetUniverse(ctx, name)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, universe)
}

func (api API) listUniverses(ctx *gin.Context) {
	universes, err := api.db.ListUniverses(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list universes: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, universes)
}

// Handler for adding a new universe
func (api *API) addUniverse(ctx *gin.Context) {
	var universe types.Universe
	if err := ctx.ShouldBindJSON(&universe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	if err := api.db.AddUniverse(context.Background(), &universe); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add universe: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, universe)
}

// Handler for updating an existing universe
func (api *API) updateUniverse(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Universe name is required"})
		return
	}

	var universe types.Universe
	if err := ctx.ShouldBindJSON(&universe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	universe.Name = name // Ensure the name in the URL matches the name in the payload
	if err := api.db.UpdateUniverse(context.Background(), &universe); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update universe: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, universe)
}

// Handler for deleting a universe
func (api *API) deleteUniverse(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Universe name is required"})
		return
	}

	if err := api.db.DeleteUniverse(context.Background(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete universe: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Universe '%s' deleted successfully", name)})
}
