package firestore

import (
	"context"
	"log"
	"lore-keeper-be/internal/env"
	"lore-keeper-be/internal/types"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	// Sets your Google Cloud Platform project ID.
	projectID  = "lore-keeper"
	characters = "characters"
	worlds     = "worlds"
	cities     = "cities"
	factions   = "factions"
	universes  = "universes"
)

var (
	tables = []string{characters, worlds, cities, factions}
)

type Database struct {
	client *firestore.Client
}

func New(dbName, mode string) (*Database, error) {
	client := createClient(context.Background(), projectID, dbName, mode)
	return &Database{
		client: client,
	}, nil
}

func createClient(ctx context.Context, projectID, dbName, mode string) *firestore.Client {
	// If we run dev build we run against an emulated db
	if mode == "dev" {
		dbhost := env.GetEnvOrDefault("FIRESTORE_EMULATOR_HOST", "localhost:8082")
		log.Printf("dbhost %s", dbhost)
	}
	if mode == "CI" {

	}
	if mode == "prod" {
	}

	client, err := firestore.NewClientWithDatabase(ctx, projectID, dbName)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	log.Println("setup local firestore client")
	return client
}

func (db *Database) InitDB(ctx context.Context) error {

	for _, table := range tables {
		if !tableExists(db.client, table) {
			db.client.Collection(table).NewDoc()
		}
		// Any need for error handling?
	}
	return nil
}

func tableExists(client *firestore.Client, table string) bool {
	iter := client.Collection(table).Documents(context.Background())
	// Iterate through the collections to find a match
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			// Collection not found
			return false
		}
		if err != nil {
			// Error occurred while iterating
			return false
		}
		// Collection found
		return true
	}
}

func (db *Database) AddCharacter(ctx context.Context, character *types.Character) error {
	_, err := db.client.Collection(characters).Doc(character.Name).Set(ctx, map[string]interface{}{
		"name":        character.Name,
		"description": character.Description,
		"age":         character.Age,
		"world":       character.World,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (db *Database) ListCharacters(ctx context.Context) ([]string, error) {
	return db.list(ctx, characters)
}

func (db *Database) DeleteCharacter(ctx context.Context, name string) error {
	return db.delete(ctx, characters, name)
}
func (db *Database) UpdateCharacter(ctx context.Context, character *types.Character) error {
	return db.AddCharacter(ctx, character)
}

func (db *Database) GetCharacter(ctx context.Context, name string) (*types.Character, error) {
	doc, err := db.client.Collection(characters).Doc(name).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}
	var character types.Character
	doc.DataTo(&character)
	return &character, nil
}

func (db *Database) AddWorld(ctx context.Context, world *types.World) error {
	_, err := db.client.Collection(worlds).Doc(world.Name).Set(ctx, map[string]interface{}{
		"name":        world.Name,
		"description": world.Description,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (db *Database) ListWorlds(ctx context.Context) ([]string, error) {
	return db.list(ctx, worlds)
}

func (db *Database) DeleteWorld(ctx context.Context, name string) error {
	return db.delete(ctx, worlds, name)
}

func (db *Database) UpdateWorld(ctx context.Context, world *types.World) error {
	return db.AddWorld(ctx, world)
}

func (db *Database) GetWorld(ctx context.Context, name string) (*types.World, error) {
	doc, err := db.client.Collection(worlds).Doc(name).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}
	var world types.World
	doc.DataTo(&world)
	return &world, nil
}

func (db *Database) AddCity(ctx context.Context, city *types.City) error {
	_, err := db.client.Collection(cities).Doc(city.Name).Set(ctx, map[string]interface{}{
		"name":              city.Name,
		"description":       city.Description,
		"foundingDate":      city.FoundingDate,
		"notableCharacters": city.NotableCharacters,
		factions:            city.Factions,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (db *Database) ListCities(ctx context.Context) ([]string, error) {
	return db.list(ctx, cities)
}

func (db *Database) DeleteCity(ctx context.Context, name string) error {
	return db.delete(ctx, cities, name)
}

func (db *Database) UpdateCity(ctx context.Context, city *types.City) error {
	return db.AddCity(ctx, city)
}

func (db *Database) GetCity(ctx context.Context, name string) (*types.City, error) {
	doc, err := db.client.Collection(cities).Doc(name).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}

	var city types.City
	doc.DataTo(&city)
	return &city, nil
}

func (db *Database) AddFaction(ctx context.Context, faction *types.Faction) error {
	_, err := db.client.Collection(factions).Doc(faction.Name).Set(ctx, map[string]interface{}{
		"name":              faction.Name,
		"description":       faction.Description,
		"foundingDate":      faction.FoundingDate,
		"notableCharacters": faction.NotableCharacters,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (db *Database) ListFactions(ctx context.Context) ([]string, error) {
	return db.list(ctx, factions)
}

func (db *Database) DeleteFaction(ctx context.Context, name string) error {

	return db.delete(ctx, factions, name)
}

func (db *Database) UpdateFaction(ctx context.Context, faction *types.Faction) error {
	return db.AddFaction(ctx, faction)
}

func (db *Database) GetFaction(ctx context.Context, name string) (*types.Faction, error) {
	doc, err := db.client.Collection(factions).Doc(name).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}

	var faction types.Faction
	doc.DataTo(&faction)
	return &faction, nil
}

func (db *Database) AddUniverse(ctx context.Context, universe *types.Universe) error {
	_, err := db.client.Collection(universes).Doc(universe.Name).Set(ctx, map[string]interface{}{
		"name":        universe.Name,
		"description": universe.Description,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func (db *Database) ListUniverses(ctx context.Context) ([]string, error) {
	return db.list(ctx, universes)
}

func (db *Database) DeleteUniverse(ctx context.Context, name string) error {
	return db.delete(ctx, universes, name)
}

func (db *Database) UpdateUniverse(ctx context.Context, universe *types.Universe) error {
	return db.AddUniverse(ctx, universe)
}

func (db *Database) GetUniverse(ctx context.Context, name string) (*types.Universe, error) {
	doc, err := db.client.Collection(universes).Doc(name).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return nil, err
	}

	var universe types.Universe
	doc.DataTo(&universe)
	return &universe, nil
}

func (db *Database) list(ctx context.Context, table string) ([]string, error) {
	iter := db.client.Collection(table).Documents(ctx)
	items := make([]string, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		items = append(items, doc.Ref.ID)
	}
	return items, nil
}

func (db *Database) delete(ctx context.Context, table, name string) error {
	_, err := db.client.Collection(table).Doc(name).Delete(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return err

}

func (db *Database) Shutdown() error {
	db.client.Close()
	return nil
}
