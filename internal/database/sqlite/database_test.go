package sqlite

import (
	"context"
	"lore-keeper-be/internal/constants"
	"lore-keeper-be/internal/types"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testdb      = "test.db"
	name        = "test"
	description = "test description"
	age         = 1
	world       = "test world"
)

var cities = []string{"city1", "city2"}

func TestCreateDB(t *testing.T) {
	defer os.Remove(testdb)
	db, err := New()
	assert.Nil(t, err)
	db.sqlite.Close()
}

func TestCRUDCharacter(t *testing.T) {
	defer os.Remove(testdb)
	db, err := New()
	assert.Nil(t, err)
	defer db.sqlite.Close()

	db.InitDB(context.TODO())
	err = db.AddCharacter(context.TODO(), &types.Character{Name: name, Description: description, Age: age, World: world})
	assert.NoError(t, err)

	character, err := db.GetCharacter(context.TODO(), name)
	assert.NoError(t, err)

	assert.Equal(t, name, character.Name)
	assert.Equal(t, description, character.Description)
	assert.Equal(t, 1, character.Age)
	assert.Equal(t, world, character.World)

	err = db.UpdateCharacter(context.TODO(), &types.Character{Name: name, Description: "new description", Age: 2, World: "new world"})
	assert.NoError(t, err)

	character, err = db.GetCharacter(context.TODO(), name)
	assert.NoError(t, err)
	assert.Equal(t, name, character.Name)
	assert.Equal(t, "new description", character.Description)
	assert.Equal(t, 2, character.Age)
	assert.Equal(t, "new world", character.World)

	err = db.DeleteCharacter(context.TODO(), name)
	assert.NoError(t, err)

	character, err = db.GetCharacter(context.TODO(), name)
	assert.ErrorIs(t, err, constants.ErrNilCharacter)
	assert.Nil(t, character)
}

func TestCRUDWorld(t *testing.T) {
	defer os.Remove(testdb)
	db, err := New()
	assert.Nil(t, err)
	defer db.sqlite.Close()

	db.InitDB(context.TODO())
	err = db.AddWorld(context.TODO(), &types.World{Name: name, Description: description, Cities: cities})
	assert.NoError(t, err)

	world, err := db.GetWorld(context.TODO(), name)
	assert.NoError(t, err)

	assert.Equal(t, name, world.Name)
	assert.Equal(t, description, world.Description)
	assert.Equal(t, cities, world.Cities)

	err = db.UpdateWorld(context.TODO(), &types.World{Name: name, Description: "new description", Cities: cities})
	assert.NoError(t, err)

	world, err = db.GetWorld(context.TODO(), name)
	assert.NoError(t, err)
	assert.Equal(t, name, world.Name)
	assert.Equal(t, "new description", world.Description)
	assert.Equal(t, cities, world.Cities)

	err = db.DeleteWorld(context.TODO(), name)
	assert.NoError(t, err)

	world, err = db.GetWorld(context.TODO(), name)
	assert.ErrorIs(t, err, constants.ErrNilWorld)
	assert.Nil(t, world)
}

// TODO check over below copilot generated tests: Check variables, valuesm consts etc
func TestCRUDCity(t *testing.T) {
	defer os.Remove(testdb)
	db, err := New()
	assert.Nil(t, err)
	defer db.sqlite.Close()

	db.InitDB(context.TODO())
	err = db.AddCity(context.TODO(), &types.City{Name: name, Description: description, FoundingDate: "1", NotableCharacters: []string{name}, Factions: []string{name}})
	assert.NoError(t, err)

	city, err := db.GetCity(context.TODO(), name)
	assert.NoError(t, err)

	assert.Equal(t, name, city.Name)
	assert.Equal(t, description, city.Description)
	assert.Equal(t, "1", city.FoundingDate)
	assert.Equal(t, []string{name}, city.NotableCharacters)
	assert.Equal(t, []string{name}, city.Factions)

	err = db.UpdateCity(context.TODO(), &types.City{Name: name, Description: "new description", FoundingDate: "2", NotableCharacters: []string{"new"}, Factions: []string{"new"}})
	assert.NoError(t, err)

	city, err = db.GetCity(context.TODO(), name)
	assert.NoError(t, err)
	assert.Equal(t, name, city.Name)
	assert.Equal(t, "new description", city.Description)
	assert.Equal(t, "2", city.FoundingDate)
	assert.Equal(t, []string{"new"}, city.NotableCharacters)
	assert.Equal(t, []string{"new"}, city.Factions)

	err = db.DeleteCity(context.TODO(), name)
	assert.NoError(t, err)

	city, err = db.GetCity(context.TODO(), name)
	assert.ErrorIs(t, err, constants.ErrNilCity)
	assert.Nil(t, city)
}

func TestCRUDFaction(t *testing.T) {
	defer os.Remove(testdb)
	db, err := New()
	assert.Nil(t, err)
	defer db.sqlite.Close()

	db.InitDB(context.TODO())
	err = db.AddFaction(context.TODO(), &types.Faction{Name: name, Description: description, FoundingDate: "1", NotableCharacters: []string{name}})
	assert.NoError(t, err)

	faction, err := db.GetFaction(context.TODO(), name)
	assert.NoError(t, err)

	assert.Equal(t, name, faction.Name)
	assert.Equal(t, description, faction.Description)
	assert.Equal(t, "1", faction.FoundingDate)
	assert.Equal(t, []string{name}, faction.NotableCharacters)

	err = db.UpdateFaction(context.TODO(), &types.Faction{Name: name, Description: "new description", FoundingDate: "2", NotableCharacters: []string{"new"}})
	assert.NoError(t, err)

	faction, err = db.GetFaction(context.TODO(), name)
	assert.NoError(t, err)
	assert.Equal(t, name, faction.Name)
	assert.Equal(t, "new description", faction.Description)
	assert.Equal(t, "2", faction.FoundingDate)
	assert.Equal(t, []string{"new"}, faction.NotableCharacters)

	err = db.DeleteFaction(context.TODO(), name)
	assert.NoError(t, err)

	faction, err = db.GetFaction(context.TODO(), name)
	assert.ErrorIs(t, err, constants.ErrNilFaction)
	assert.Nil(t, faction)
}
