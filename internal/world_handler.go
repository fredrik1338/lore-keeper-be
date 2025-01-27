package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) getWorld(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "World name is required"})
		return
	}

	world, err := api.db.GetWorld(ctx, name)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, world)
}

func (api API) listWorlds(ctx *gin.Context) {
	worlds, err := api.db.ListWorlds(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list worlds: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, worlds)
}

// Handler for adding a new world
func (api *API) addWorld(ctx *gin.Context) {
	var world types.World
	if err := ctx.ShouldBindJSON(&world); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	if err := api.db.AddWorld(context.Background(), &world); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add world: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, world)
}

// Handler for updating an existing world
func (api *API) updateWorld(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "World name is required"})
		return
	}

	var world types.World
	if err := ctx.ShouldBindJSON(&world); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	world.Name = name // Ensure the name in the URL matches the name in the payload
	if err := api.db.UpdateWorld(context.Background(), &world); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update world: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, world)
}

// Handler for deleting a world
func (api *API) deleteWorld(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "World name is required"})
		return
	}

	if err := api.db.DeleteWorld(context.Background(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete world: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("World '%s' deleted successfully", name)})
}
