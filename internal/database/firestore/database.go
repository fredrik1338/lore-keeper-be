package firestore

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"lore-keeper-be/internal/dto"
// 	"lore-keeper-be/internal/env"

// 	"cloud.google.com/go/firestore"
// 	"google.golang.org/api/iterator"
// )

// const (
// 	TODOs = "todos"
// )

// type Database struct {
// 	client *firestore.Client
// }

// func New(projectID, dbName, mode string) (*Database, error) {
// 	client := createClient(context.Background(), projectID, dbName, mode)
// 	return &Database{
// 		client: client,
// 	}, nil
// }

// func createClient(ctx context.Context, projectID, dbName, mode string) *firestore.Client {
// 	// If we run dev build we run against an emulated db
// 	if mode == "dev" {
// 		dbhost := env.GetEnvOrDefault("FIRESTORE_EMULATOR_HOST", "localhost:8082")
// 		log.Printf("dbhost %s", dbhost)
// 	}
// 	if mode == "CI" {

// 	}
// 	if mode == "prod" {
// 	}

// 	client, err := firestore.NewClientWithDatabase(ctx, projectID, dbName)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}
// 	log.Println("setup local firestore client")
// 	return client
// }

// func (db *Database) InitDB() error {
// 	if !dbSetup(db.client) {
// 		db.client.Collection(TODOs).NewDoc()
// 	}
// 	return nil
// }

// func dbSetup(client *firestore.Client) bool {
// 	iter := client.Collection(TODOs).Documents(context.Background())
// 	// Iterate through the collections to find a match
// 	for {
// 		_, err := iter.Next()
// 		if err == iterator.Done {
// 			// Collection not found
// 			return false
// 		}
// 		if err != nil {
// 			// Error occurred while iterating
// 			return false
// 		}
// 		// Collection found
// 		return true
// 	}
// }

// func (db *Database) GetTODOs() ([]dto.TODO, error) {
// 	iter := db.client.Collection(TODOs).Documents(context.Background())
// 	todos := make([]dto.TODO, 0)
// 	for {
// 		doc, err := iter.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 		jsonData, err := json.Marshal(doc.Data())
// 		if err != nil {
// 			log.Fatalf("Error marshaling Firestore document data: %v", err)
// 		}

// 		var todo dto.TODO
// 		if err := json.Unmarshal(jsonData, &todo); err != nil {
// 			log.Fatalf("Error unmarshaling JSON data into MyData struct: %v", err)
// 		}
// 		todos = append(todos, todo)
// 	}
// 	return todos, nil
// }
// func (db *Database) AddTODO(data dto.TODO) error {
// 	_, err := db.client.Collection(TODOs).Doc(data.Title).Set(context.TODO(), map[string]interface{}{
// 		"title":       data.Title,
// 		"description": data.Description,
// 	})
// 	if err != nil {
// 		log.Printf("An error has occurred: %s", err)
// 	}

// 	return err
// }
// func (db *Database) DeleteTODO(title string) (bool, error) {
// 	_, err := db.client.Collection(TODOs).Doc(title).Delete(context.Background())
// 	if err != nil {
// 		// Handle any errors in an appropriate way, such as returning them.
// 		log.Printf("An error has occurred: %s", err)
// 		return false, err
// 	}
// 	return true, nil
// }

// func (db *Database) Shutdown() error {
// 	db.client.Close()
// 	return nil
// }
