package main

import (
	"context"
	"fmt"
	"log"
	"lore-keeper-be/internal"
	"lore-keeper-be/internal/aiservice"
	"lore-keeper-be/internal/database/firestore"
	"lore-keeper-be/internal/env"
)

const (
	// Sets your Google Cloud Platform project ID.
	projectID = "august-journey-434715-u0"
	location  = "us-central1"
	model     = "imagen-3.0-generate-001"
)

func main() {
	//TODO implement
	mode := env.GetEnvOrDefault("RUN_MODE", env.DefaultMode)
	dbName := env.GetEnvOrDefault("DB_NAME", env.DefaultDB)
	database, err := firestore.New(dbName, mode, projectID)
	if err != nil {
		log.Fatalf("Could not open DB due to err: %v", err)
		return
	}

	// database, err := sqlite.New()

	// if err != nil {
	// 	panic(fmt.Sprintf("Could not create DB %s", err.Error()))
	// }

	err = database.InitDB(context.TODO()) //TODO this context should be passed fomr above to allow graceful shutdown
	if err != nil {
		panic(fmt.Sprintf("Could not init DB %s", err.Error()))
	}

	// // Check Application Default Credentials (ADC)
	// creds, err := transport.Creds(ctx, option.WithScopes("https://www.googleapis.com/auth/cloud-platform"))
	// if err != nil {
	// 	log.Fatalf("Failed to find default credentials: %v", err)
	// }
	// // Debug log. should be removed in production
	// token, err := creds.TokenSource.Token()
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token: %v", err)
	// }
	// log.Printf("Retrieved token: %s", token.AccessToken)

	aiService := aiservice.New(projectID, location, model, mode)

	server := internal.NewAPI(database, aiService)

	port := env.GetEnvOrDefault(env.Port, env.DefaultPort)
	host := env.GetEnvOrDefault(env.Host, env.DefaultHost)
	server.Run(host, port)

	//TODO add graceful shutdown
	// This should include listening for shutdown signals
	defer database.Shutdown()
}
