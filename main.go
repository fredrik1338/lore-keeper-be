package main

import (
	"context"
	"fmt"
	"log"
	"lore-keeper-be/internal"
	"lore-keeper-be/internal/database/firestore"
	"lore-keeper-be/internal/env"
)

const (
	// Sets your Google Cloud Platform project ID.
	projectID = "lore-keeper-project"
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

	server := internal.NewAPI(database)

	port := env.GetEnvOrDefault(env.Port, env.DefaultPort)
	host := env.GetEnvOrDefault(env.Host, env.DefaultHost)
	server.Run(host, port)

	//TODO add graceful shutdown
	// This should include listening for shutdown signals
	defer database.Shutdown()
}
