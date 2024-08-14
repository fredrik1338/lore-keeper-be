package main

import (
	"context"
	"fmt"
	"lore-keeper-be/internal"
	"lore-keeper-be/internal/database/sqlite"
	"lore-keeper-be/internal/env"
)

func main() {
	//TODO implement
	// mode := env.GetEnvOrDefault("RUN_MODE", env.DefaultMode)
	// dbName := env.GetEnvOrDefault("DB_NAME", env.DefaultDB)
	// db, err := firestore.New(projectID, dbName, mode)
	// if err != nil {
	// 	log.Fatalf("Could not open DB due to err: %v", err)
	// 	return
	// }
	database, err := sqlite.New()

	if err != nil {
		panic(fmt.Sprintf("Could not crate DB %s", err.Error()))
	}

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
