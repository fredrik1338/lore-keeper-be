package internal

import (
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/database/pg"
	"net/http"
)

type Server struct {
	multiplexer *http.ServeMux
	address     string
	db          database.Database
}

func newServer() Server {
	mux := http.NewServeMux()

	database, err := pg.New()

	if err != nil {
		panic(fmt.Sprintf("Could not setup DB %s", err.Error()))
	}

	//TODO set handler for any new functions
	// To test just use curl localhost:8080/api/v1/lore-keeper/<characters or worlds> TODOD create better examples

	// mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, characters), handleCharacters)
	// mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, worlds), handleWorlds)
	// mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, cities), handleCities)
	// mux.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, factions), handleFactions)
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { fmt.Print("hello") })

	return Server{
		multiplexer: mux,
		address:     ":8080",
		db:          database,
	}
}

func (server Server) Start() error {

	server.multiplexer.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, characters), server.handleCharacters)
	server.multiplexer.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, worlds), server.handleWorlds)
	server.multiplexer.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, cities), server.handleCities)
	server.multiplexer.HandleFunc(fmt.Sprintf("/%s/%s", apiPath, factions), server.handleFactions)

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
