package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"lore-keeper-be/internal/types"

	_ "github.com/lib/pq"
)

var (
	errNilPerson  = errors.New("character is nil")
	errNilCity    = errors.New("city is nil")
	errNilWorld   = errors.New("world is nil")
	errNilFaction = errors.New("faction is nil")
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
	_, err := db.pg.QueryContext(ctx, setupCharactersTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, setupWorldsTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, setupCitiesTable)
	if err != nil {
		return err
	}
	_, err = db.pg.QueryContext(ctx, setupFactionsTable)
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
		return errNilPerson
	}
	_, err := db.pg.QueryContext(ctx, addPersonQuery, character.Name, character.Age, character.Home)
	return err
}

func (db *Database) DeleteCharacter(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, deletePersonQuery, name)
	return err
}

func (db *Database) UpdateCharacter(ctx context.Context, character *types.Character) error {
	_, err := db.pg.QueryContext(ctx, updatePersonQuery,
		character.Name,
		character.Age,
		character.Home)
	return err
}

func (db *Database) GetCharacter(ctx context.Context, name string) (*types.Character, error) {
	var person types.Character
	rows, err := db.pg.QueryContext(ctx, getPersonQuery, name)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&person.Name, &person.Age, &person.Home)
		if err != nil {
			fmt.Printf("got error %s", err.Error())
			return nil, err
		}
	} else {
		return nil, errNilPerson
	}

	rows.Close()
	fmt.Printf("got person %v", person)
	return &person, nil
}

func (db *Database) AddWorld(ctx context.Context, world *types.World) error {
	if world == nil {
		return errNilWorld
	}
	_, err := db.pg.QueryContext(ctx,
		addWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func (db *Database) UpdateWorld(ctx context.Context, world *types.World) error {
	if world == nil {
		return errNilWorld
	}
	_, err := db.pg.QueryContext(ctx, updateWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func (db *Database) GetWorld(ctx context.Context, name string) (*types.World, error) {
	var world types.World
	rows, err := db.pg.QueryContext(ctx, getWorldQuery)
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
	_, err := db.pg.QueryContext(ctx, getWorldQuery)
	return err
}

func (db *Database) AddCity(ctx context.Context, city *types.City) error {
	if city == nil {
		return errNilCity
	}
	_, err := db.pg.QueryContext(ctx, addCityQuery, city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func (db *Database) UpdateCity(ctx context.Context, city *types.City) error {
	if city == nil {
		return errNilCity
	}
	_, err := db.pg.QueryContext(ctx, updateCityQuery,
		city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func (db *Database) GetCity(ctx context.Context, name string) (*types.City, error) {
	var city types.City
	rows, err := db.pg.QueryContext(ctx, getCityQuery, name)
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
	_, err := db.pg.QueryContext(ctx, getCityQuery)
	return err
}

func (db *Database) AddFaction(ctx context.Context, faction *types.Faction) error {
	if faction == nil {
		return errNilFaction
	}
	_, err := db.pg.QueryContext(ctx, addFactionQuery, faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters,
		faction.Leader)
	return err
}

func (db *Database) UpdateFaction(ctx context.Context, faction *types.Faction) error {
	if faction == nil {
		return errNilFaction
	}
	_, err := db.pg.QueryContext(ctx, updateFactionQuery,
		faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters,
		faction.Leader)
	return err
}

func (db *Database) GetFaction(ctx context.Context, name string) (*types.Faction, error) {
	var faction types.Faction
	rows, err := db.pg.QueryContext(ctx, getFactionQuery, name)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&faction.Name,
			&faction.Description,
			&faction.FoundingDate,
			faction.NotableCharacters,
			faction.Leader)
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
	_, err := db.pg.QueryContext(ctx, getFactionQuery, name)
	return err
}
