package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"lore-keeper-be/internal/constants"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/types"

	_ "github.com/lib/pq"
)

type Database struct {
	pg *sql.DB
}

func New() (*Database, error) {
	pg := getDBConn()
	database := Database{
		pg: pg,
	}

	err := database.InitDB(context.Background())
	if err != nil {
		// panic(fmt.Sprintf("Could not setup DB %s", err.Error()))
		return nil, err
	}
	return &database, nil
}

// TODO get the connstr from env variables
const (
	connStr = "postgresql://admin:pgadmin@localhost/MyfirstDB?sslmode=disable"
)

func getDBConn() *sql.DB {
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) //TODO change this to return error
	}
	return db
}

func (db *Database) InitDB(ctx context.Context) error {
	_, err := db.pg.QueryContext(ctx, database.SetupCharactersTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, database.SetupWorldsTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, database.SetupCitiesTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, database.SetupFactionsTable)
	if err != nil {
		return err
	}

	return nil
}

// TODO Investigate what we need in this method
func (db *Database) Shutdown() error {
	return db.pg.Close()
}

func (db *Database) AddCharacter(ctx context.Context, character *types.Character) error {
	if character == nil {
		return constants.ErrNilCharacter
	}
	_, err := db.pg.QueryContext(ctx, database.AddCharacterQuery, character.Name, character.Age, character.World)
	return err
}

func (db *Database) DeleteCharacter(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, database.DeleteCharacterQuery, name)
	return err
}

func (db *Database) UpdateCharacter(ctx context.Context, character *types.Character) error {
	_, err := db.pg.QueryContext(ctx, database.UpdateCharacterQuery,
		character.Name,
		character.Age,
		character.World)
	return err
}

func (db *Database) GetCharacter(ctx context.Context, name string) (*types.Character, error) {
	var person types.Character
	rows, err := db.pg.QueryContext(ctx, database.GetCharacterQuery, name)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&person.Name, &person.Age, &person.World)
		if err != nil {
			fmt.Printf("got error %s", err.Error())
			return nil, err
		}
	} else {
		return nil, constants.ErrNilCharacter
	}

	rows.Close()
	fmt.Printf("got person %v", person)
	return &person, nil
}

func (db *Database) AddWorld(ctx context.Context, world *types.World) error {
	if world == nil {
		return constants.ErrNilWorld
	}
	_, err := db.pg.QueryContext(ctx,
		database.AddWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func (db *Database) UpdateWorld(ctx context.Context, world *types.World) error {
	if world == nil {
		return constants.ErrNilWorld
	}
	_, err := db.pg.QueryContext(ctx, database.UpdateWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func (db *Database) GetWorld(ctx context.Context, name string) (*types.World, error) {
	var world types.World
	rows, err := db.pg.QueryContext(ctx, database.GetWorldQuery)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&world.Name, &world.Description, &world.Cities)
		if err != nil {
			fmt.Printf("got error %s", err.Error())
			return nil, err
		}
	}

	rows.Close()
	fmt.Printf("got person %v", world)
	return &world, nil
}

func (db *Database) DeleteWorld(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, database.DeleteWorldQuery)
	return err
}

func (db *Database) AddCity(ctx context.Context, city *types.City) error {
	if city == nil {
		return constants.ErrNilCity
	}
	_, err := db.pg.QueryContext(ctx, database.AddCityQuery, city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func (db *Database) UpdateCity(ctx context.Context, city *types.City) error {
	if city == nil {
		return constants.ErrNilCity
	}
	_, err := db.pg.QueryContext(ctx, database.UpdateCityQuery,
		city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func (db *Database) GetCity(ctx context.Context, name string) (*types.City, error) {
	var city types.City
	rows, err := db.pg.QueryContext(ctx, database.GetCityQuery, name)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&city.Name,
			&city.Description,
			&city.FoundingDate,
			city.NotableCharacters,
			city.Factions)
		if err != nil {
			fmt.Printf("got error %s", err.Error())
			return nil, err
		}
	}

	rows.Close()
	fmt.Printf("got city %v", city)
	return &city, nil
}

func (db *Database) DeleteCity(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, database.DeleteCityQuery)
	return err
}

func (db *Database) AddFaction(ctx context.Context, faction *types.Faction) error {
	if faction == nil {
		return constants.ErrNilFaction
	}
	_, err := db.pg.QueryContext(ctx, database.AddFactionQuery, faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters,
	)
	return err
}

func (db *Database) UpdateFaction(ctx context.Context, faction *types.Faction) error {
	if faction == nil {
		return constants.ErrNilFaction
	}
	_, err := db.pg.QueryContext(ctx, database.UpdateFactionQuery,
		faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters)
	return err
}

func (db *Database) GetFaction(ctx context.Context, name string) (*types.Faction, error) {
	var faction types.Faction
	rows, err := db.pg.QueryContext(ctx, database.GetFactionQuery, name)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&faction.Name,
			&faction.Description,
			&faction.FoundingDate,
			faction.NotableCharacters)
		if err != nil {
			fmt.Printf("got error %s", err.Error())
			return nil, err
		}
	}

	rows.Close()
	fmt.Printf("got faction %v", faction)
	return &faction, nil
}

func (db *Database) DeleteFaction(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, database.DeleteFactionQuery, name)
	return err
}
