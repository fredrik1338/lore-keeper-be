package internal

import (
	"fmt"
	"lore-keeper-be/internal/database"

	"github.com/gin-gonic/gin"
)

const (
	basePath   = "api/v1/"
	lorePath   = "lore-keeper"
	characters = "characters"
	worlds     = "worlds"
	cities     = "cities"
	factions   = "factions"
)

type dbAPI struct {
	db database.Database // TODO gotta move stuff around, this is not clear
}

func NewAPI(db database.Database) *dbAPI {
	return &dbAPI{
		db: db,
	}
}

func (api dbAPI) Run(host, port string) {
	router := gin.Default()
	router.Use(Options)
	v1 := router.Group(basePath)
	lore := v1.Group(lorePath)
	characters := lore.Group(characters)
	characters.GET("", api.listCharacters)
	characters.GET("/:name", api.getCharacter)
	characters.POST("", api.addCharacter)
	characters.DELETE("/:name", api.deleteCharacter)
	characters.PUT("/:name", api.updateCharacter)

	worlds := lore.Group(worlds)
	worlds.GET("", api.listWorlds)
	worlds.GET("/:name", api.getWorld)
	worlds.POST("", api.addWorld)
	worlds.DELETE("/:name", api.deleteWorld)
	worlds.PUT("/:name", api.updateWorld)

	cities := lore.Group(cities)
	cities.GET("", api.listCities)
	cities.GET("/:name", api.getCity)
	cities.POST("", api.addCity)
	cities.DELETE("/:name", api.deleteCity)
	cities.PUT("/:name", api.updateCity)

	factions := lore.Group(factions)
	factions.GET("", api.listFactions)
	factions.GET("/:name", api.getFaction)
	factions.POST("", api.addFaction)
	factions.DELETE("/:name", api.deleteFaction)
	factions.PUT("/:name", api.updateFaction)

	router.Run(fmt.Sprintf("%s:%s", host, port))

}
