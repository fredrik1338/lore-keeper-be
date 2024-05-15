package internal

import (
	"context"
	"fmt"
	"net/http"
	"serverapp/internal/database"
)

type Server struct {
	multiplexer *http.ServeMux
	address     string
}

func newServer() Server {
	mux := http.NewServeMux()

	//TODO set handler for any new functions
	// To test just use curl localhost:8080/api/v1/lore-keeper/<characters or worlds> TODOD create better examples

	mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, characters), handleCharacters)
	mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, worlds), handleWorlds)
	mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, cities), handleCities)
	mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, factions), handleFactions)
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { fmt.Print("hello") })

	return Server{
		multiplexer: mux,
		address:     ":8080",
	}
}

func (server Server) Start() error {
	db := getDBConn()

	err := database.SetupDB(context.Background(), db)
	if err != nil {
		panic(fmt.Sprintf("Could not setup DB %s", err.Error()))
	}
	return http.ListenAndServe(server.address, server.multiplexer)
}

// Example queries

// In postman:
// method localhost:8080/api/v1/lore-keeper/characters

// {
//     "name": "Bob",
//     "description": "very bobbish",
//     "age": 25,
//     "home": "narnia"
// }
