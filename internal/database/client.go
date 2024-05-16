package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"lore-keeper-be/internal/types"

	_ "github.com/lib/pq"
)

var (
	errNilPerson  = errors.New("character is nil")
	errNilCity    = errors.New("city is nil")
	errNilWorld   = errors.New("world is nil")
	errNilFaction = errors.New("faction is nil")
)

func SetupDB(ctx context.Context, db *sql.DB) error {
	_, err := db.QueryContext(ctx, setupCharactersTable)
	if err != nil {
		return err
	}
	_, err = db.QueryContext(ctx, setupWorldsTable)
	if err != nil {
		return err
	}
	_, err = db.QueryContext(ctx, setupCitiesTable)
	if err != nil {
		return err
	}
	_, err = db.QueryContext(ctx, setupFactionsTable)
	if err != nil {
		return err
	}

	return nil
}

func AddCharacter(ctx context.Context, db *sql.DB, character *types.Character) error {
	if character == nil {
		return errNilPerson
	}
	_, err := db.QueryContext(ctx, addPersonQuery, character.Name, character.Age, character.Home)
	return err
}

func DeleteCharacter(ctx context.Context, db *sql.DB, name string) error {
	_, err := db.QueryContext(ctx, deletePersonQuery, name)
	return err
}

func UpdateCharacter(ctx context.Context, db *sql.DB, character *types.Character) error {
	_, err := db.QueryContext(ctx, updatePersonQuery,
		character.Name,
		character.Age,
		character.Home)
	return err
}

func GetCharacter(ctx context.Context, db *sql.DB, name string) (*types.Character, error) {
	var person types.Character
	rows, err := db.QueryContext(ctx, getPersonQuery, name)
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

func AddWorld(ctx context.Context, db *sql.DB, world *types.World) error {
	if world == nil {
		return errNilWorld
	}
	_, err := db.QueryContext(ctx,
		addWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func UpdateWorld(ctx context.Context, db *sql.DB, world *types.World) error {
	if world == nil {
		return errNilWorld
	}
	_, err := db.QueryContext(ctx, updateWorldQuery,
		world.Name,
		world.Description,
		world.Cities)
	return err
}

func GetWorld(ctx context.Context, db *sql.DB, name string) (*types.World, error) {
	var world types.World
	rows, err := db.QueryContext(ctx, getWorldQuery)
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

func DeleteWorld(ctx context.Context, db *sql.DB, name string) error {
	_, err := db.QueryContext(ctx, getWorldQuery)
	return err
}

func AddCity(ctx context.Context, db *sql.DB, city *types.City) error {
	if city == nil {
		return errNilCity
	}
	_, err := db.QueryContext(ctx, addCityQuery, city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func UpdateCity(ctx context.Context, db *sql.DB, city *types.City) error {
	if city == nil {
		return errNilCity
	}
	_, err := db.QueryContext(ctx, updateCityQuery,
		city.Name,
		city.Description,
		city.FoundingDate,
		city.NotableCharacters,
		city.Factions)
	return err
}

func GetCity(ctx context.Context, db *sql.DB, name string) (*types.City, error) {
	var city types.City
	rows, err := db.QueryContext(ctx, getCityQuery, name)
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

func DeleteCity(ctx context.Context, db *sql.DB, name string) error {
	_, err := db.QueryContext(ctx, getCityQuery)
	return err
}

func AddFaction(ctx context.Context, db *sql.DB, faction *types.Faction) error {
	if faction == nil {
		return errNilFaction
	}
	_, err := db.QueryContext(ctx, addFactionQuery, faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters,
		faction.Leader)
	return err
}

func UpdateFaction(ctx context.Context, db *sql.DB, faction *types.Faction) error {
	if faction == nil {
		return errNilFaction
	}
	_, err := db.QueryContext(ctx, updateFactionQuery,
		faction.Name,
		faction.Description,
		faction.FoundingDate,
		faction.NotableCharacters,
		faction.Leader)
	return err
}

func GetFaction(ctx context.Context, db *sql.DB, name string) (*types.Faction, error) {
	var faction types.Faction
	rows, err := db.QueryContext(ctx, getFactionQuery, name)
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

func DeleteFaction(ctx context.Context, db *sql.DB, name string) error {
	_, err := db.QueryContext(ctx, getFactionQuery, name)
	return err
}
