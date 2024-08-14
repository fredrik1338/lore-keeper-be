package internal

import (
	"context"
	"fmt"
	"lore-keeper-be/internal/database"
	"lore-keeper-be/internal/database/sqlite"
	"net/http"
)

type Server struct {
	multiplexer *http.ServeMux
	address     string
	db          database.Database
}

func newServer() Server {
	mux := http.NewServeMux()

	// TODO add a flag to choose the database
	// database, err := pg.New()
	database, err := sqlite.New()

	if err != nil {
		panic(fmt.Sprintf("Could not crate DB %s", err.Error()))
	}

	err = database.InitDB(context.TODO()) //TODO this context should be passed fomr above to allow graceful shutdown
	if err != nil {
		panic(fmt.Sprintf("Could not init DB %s", err.Error()))
	}

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
	server.multiplexer.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { fmt.Print("hello") })

	// Wrap the entire mux with the CORS middleware
	corsMux := corsMiddleware(server.multiplexer)

	return http.ListenAndServe(server.address, corsMux)
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
