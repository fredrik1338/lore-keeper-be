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

func (db *Database) ListCharacters(ctx context.Context) ([]string, error) {
	var characters []string
	rows, err := db.pg.QueryContext(ctx, database.ListCharactersQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var character string
		err := rows.Scan(&character)
		if err != nil {

			return nil, err
		}
		characters = append(characters, character)
	}

	rows.Close()
	return characters, nil
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
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&person.Name, &person.Age, &person.World)
		if err != nil {

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

func (db *Database) ListWorlds(ctx context.Context) ([]string, error) {
	var worlds []string
	rows, err := db.pg.QueryContext(ctx, database.ListWorldsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var world string
		err := rows.Scan(&world)
		if err != nil {

			return nil, err
		}
		worlds = append(worlds, world)
	}

	rows.Close()
	fmt.Printf("got worlds %v", worlds)
	return worlds, nil
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
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&world.Name, &world.Description, &world.Cities)
		if err != nil {

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

func (db *Database) ListCities(ctx context.Context) ([]string, error) {
	var cities []string
	rows, err := db.pg.QueryContext(ctx, database.ListCitiesQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var city string
		err := rows.Scan(&city)
		if err != nil {

			return nil, err
		}
		cities = append(cities, city)
	}

	rows.Close()
	fmt.Printf("got cities %v", cities)
	return cities, nil
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
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&city.Name,
			&city.Description,
			&city.FoundingDate,
			city.NotableCharacters,
			city.Factions)
		if err != nil {

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

func (db *Database) ListFactions(ctx context.Context) ([]string, error) {
	var factions []string
	rows, err := db.pg.QueryContext(ctx, database.ListFactionsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var faction string
		err := rows.Scan(&faction)
		if err != nil {

			return nil, err
		}
		factions = append(factions, faction)
	}

	rows.Close()
	fmt.Printf("got factions %v", factions)
	return factions, nil
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
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&faction.Name,
			&faction.Description,
			&faction.FoundingDate,
			faction.NotableCharacters)
		if err != nil {

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

func (db *Database) AddUniverse(ctx context.Context, universe *types.Universe) error {
	if universe == nil {
		return constants.ErrNilUniverse
	}
	_, err := db.pg.QueryContext(ctx, database.AddUniverseQuery, universe.Name,
		universe.Description)
	return err
}

func (db *Database) GetUniverse(ctx context.Context, name string) (*types.Universe, error) {
	var universe types.Universe
	rows, err := db.pg.QueryContext(ctx, database.GetUniverseQuery, name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&universe.Name, &universe.Description)
		if err != nil {

			return nil, err
		}
	}

	rows.Close()
	fmt.Printf("got universe %v", universe)
	return &universe, nil
}

func (db *Database) ListUniverses(ctx context.Context) ([]string, error) {
	var universes []string
	rows, err := db.pg.QueryContext(ctx, database.ListUniversesQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var universe string
		err := rows.Scan(&universe)
		if err != nil {

			return nil, err
		}
		universes = append(universes, universe)
	}

	rows.Close()
	fmt.Printf("got universes %v", universes)
	return universes, nil
}

func (db *Database) UpdateUniverse(ctx context.Context, universe *types.Universe) error {
	// TODO figure out what to update

	return nil
}

func (db *Database) DeleteUniverse(ctx context.Context, name string) error {
	_, err := db.pg.QueryContext(ctx, database.DeleteUniverseQuery, name)
	return err
}
