package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api dbAPI) getCharacter(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Character name is required"})
		return
	}

	character, err := api.db.GetCharacter(ctx, name)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, character)
}

func (api dbAPI) listCharacters(ctx *gin.Context) {
	characters, err := api.db.ListCharacters(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list characters: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, characters)
}

// Handler for adding a new character
func (api *dbAPI) addCharacter(ctx *gin.Context) {
	var character types.Character
	if err := ctx.ShouldBindJSON(&character); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	if err := api.db.AddCharacter(context.Background(), &character); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add character: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, character)
}

// Handler for updating an existing character
func (api *dbAPI) updateCharacter(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Character name is required"})
		return
	}

	var character types.Character
	if err := ctx.ShouldBindJSON(&character); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	character.Name = name // Ensure the name in the URL matches the name in the payload
	if err := api.db.UpdateCharacter(context.Background(), &character); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update character: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, character)
}

// Handler for deleting a character
func (api *dbAPI) deleteCharacter(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Character name is required"})
		return
	}

	if err := api.db.DeleteCharacter(context.Background(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete character: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Character '%s' deleted successfully", name)})
}
