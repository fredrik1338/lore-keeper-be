package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"lore-keeper-be/internal/constants"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Database struct {
	sqlite *sql.DB
}

func New() (*Database, error) {
	path := "test.db"

	os.Remove(path)              // Delete the file to avoid duplicated records.
	file, err := os.Create(path) // Create SQLite file
	if err != nil {
		return nil, err
	}
	file.Close()

	sqliteDatabase, err := sql.Open("sqlite3", fmt.Sprintf("./%s", path))
	if err != nil {
		return nil, err
	}

	return &Database{
		sqlite: sqliteDatabase,
	}, nil
}

func (db *Database) InitDB(ctx context.Context) error {
	statement, err := db.sqlite.Prepare(database.SetupCharactersTable)
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.sqlite.Prepare(database.SetupWorldsTable)
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.sqlite.Prepare(database.SetupCitiesTable)
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.sqlite.Prepare(database.SetupFactionsTable)
	if err != nil {
		return err
	}
	statement.Exec()

	return nil
}

func (db *Database) AddCharacter(ctx context.Context, character *types.Character) error {
	statement, err := db.sqlite.Prepare(database.AddCharacterQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(character.Name, character.Description, character.Age, character.World)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) ListCharacters(ctx context.Context) ([]string, error) {
	var characters []string
	rows, err := db.sqlite.Query(database.ListCharactersQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var character string
		err = rows.Scan(&character)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	rows.Close()
	return characters, nil
}

func (db *Database) DeleteCharacter(ctx context.Context, name string) error {
	statement, err := db.sqlite.Prepare(database.DeleteCharacterQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	if err != nil {
		return err
	}
	return nil

}

func (db *Database) UpdateCharacter(ctx context.Context, character *types.Character) error {
	statement, err := db.sqlite.Prepare(database.UpdateCharacterQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(character.Name, character.Description, character.Age, character.World)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetCharacter(ctx context.Context, name string) (*types.Character, error) {
	var character types.Character
	rows, err := db.sqlite.Query(database.GetCharacterQuery, name)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&character.Name, &character.Description, &character.Age, &character.World)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, constants.ErrNilCharacter
	}
	rows.Close()
	return &character, nil

}

func (db *Database) AddWorld(ctx context.Context, world *types.World) error {
	statement, err := db.sqlite.Prepare(database.AddWorldQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(world.Name, world.Description, strings.Join(world.Cities, ","))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) ListWorlds(ctx context.Context) ([]string, error) {
	var worlds []string
	rows, err := db.sqlite.Query(database.ListWorldsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var world string
		err = rows.Scan(&world)
		if err != nil {
			return nil, err
		}
		worlds = append(worlds, world)
	}
	rows.Close()
	return worlds, nil
}

func (db *Database) DeleteWorld(ctx context.Context, name string) error {
	statement, err := db.sqlite.Prepare(database.DeleteWorldQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateWorld(ctx context.Context, world *types.World) error {
	statement, err := db.sqlite.Prepare(database.UpdateWorldQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(world.Name, world.Description, strings.Join(world.Cities, ","))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetWorld(ctx context.Context, name string) (*types.World, error) {
	var world types.World
	var cities string
	rows, err := db.sqlite.Query(database.GetWorldQuery, name)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&world.Name, &world.Description, &cities)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, constants.ErrNilWorld
	}
	rows.Close()

	world.Cities = strings.Split(cities, ",")
	return &world, nil
}

func (db *Database) AddCity(ctx context.Context, city *types.City) error {
	statement, err := db.sqlite.Prepare(database.AddCityQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(city.Name, city.Description, city.FoundingDate, strings.Join(city.NotableCharacters, ","), strings.Join(city.Factions, ","))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) ListCities(ctx context.Context) ([]string, error) {
	var cities []string
	rows, err := db.sqlite.Query(database.ListCitiesQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var city string
		err = rows.Scan(&city)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	rows.Close()
	return cities, nil
}

func (db *Database) DeleteCity(ctx context.Context, name string) error {
	statement, err := db.sqlite.Prepare(database.DeleteCityQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateCity(ctx context.Context, city *types.City) error {
	statement, err := db.sqlite.Prepare(database.UpdateCityQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(city.Name, city.Description, city.FoundingDate, strings.Join(city.NotableCharacters, ","), strings.Join(city.Factions, ","))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetCity(ctx context.Context, name string) (*types.City, error) {
	var city types.City
	var factions, notableCharacters string
	rows, err := db.sqlite.Query(database.GetCityQuery, name)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&city.Name, &city.Description, &city.FoundingDate, &notableCharacters, &factions)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, constants.ErrNilCity
	}
	rows.Close()

	city.Factions = strings.Split(factions, ",")
	city.NotableCharacters = strings.Split(notableCharacters, ",")

	return &city, nil
}

func (db *Database) AddFaction(ctx context.Context, faction *types.Faction) error {
	statement, err := db.sqlite.Prepare(database.AddFactionQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(faction.Name, faction.Description, faction.FoundingDate, strings.Join(faction.NotableCharacters, ","))
	return err
}

func (db *Database) ListFactions(ctx context.Context) ([]string, error) {
	var factions []string
	rows, err := db.sqlite.Query(database.ListFactionsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var faction string
		err = rows.Scan(&faction)
		if err != nil {
			return nil, err
		}
		factions = append(factions, faction)
	}
	rows.Close()
	return factions, nil
}

func (db *Database) DeleteFaction(ctx context.Context, name string) error {
	statement, err := db.sqlite.Prepare(database.DeleteFactionQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	return err
}

func (db *Database) UpdateFaction(ctx context.Context, faction *types.Faction) error {
	statement, err := db.sqlite.Prepare(database.UpdateFactionQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(faction.Name, faction.Description, faction.FoundingDate, strings.Join(faction.NotableCharacters, ","))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetFaction(ctx context.Context, name string) (*types.Faction, error) {
	var faction types.Faction
	var notableCharacters string
	rows, err := db.sqlite.Query(database.GetFactionQuery, name)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&faction.Name, &faction.Description, &faction.FoundingDate, &notableCharacters)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, constants.ErrNilFaction
	}
	rows.Close()
	faction.NotableCharacters = strings.Split(notableCharacters, ",")

	return &faction, nil
}

func (db *Database) Shutdown() error {
	return db.sqlite.Close()
}
