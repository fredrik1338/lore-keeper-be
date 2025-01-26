package internal

import (
	"context"
	"encoding/base64"
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

func (api *dbAPI) addCharacter(ctx *gin.Context) {
	var character types.Character
	if err := ctx.ShouldBindJSON(&character); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	// Validate that the Base64 string is valid, if provided
	if character.ProfilePicture != "" {
		if _, err := base64.StdEncoding.DecodeString(character.ProfilePicture); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid profile picture (Base64): %s", err.Error())})
			return
		}
	}

	// Save the character to Firestore
	if err := api.db.AddCharacter(context.Background(), &character); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add character: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, character)
}

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

	// Fetch the existing character to ensure partial updates work
	existingCharacter, err := api.db.GetCharacter(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Character not found: %s", err.Error())})
		return
	}

	// Update fields
	character.Name = name // Ensure the name in the URL matches the character name
	if character.ProfilePicture == "" {
		// Retain the existing profile picture if no new one is provided
		character.ProfilePicture = existingCharacter.ProfilePicture
	} else {
		// Validate the new Base64 string
		if _, err := base64.StdEncoding.DecodeString(character.ProfilePicture); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid profile picture (Base64): %s", err.Error())})
			return
		}
	}

	// Save the updated character
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
