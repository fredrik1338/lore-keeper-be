package internal

import (
	"database/sql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	basePath   = "api/v1/"
	apiPath    = basePath + "lore-keeper"
	characters = "characters"
	worlds     = "worlds"
	cities     = "cities"
	factions   = "factions"
)

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

func handleCharacters(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		message, status = getCharacter(request.Context(), body)
	case http.MethodPost:
		message, status = addCharacter(request.Context(), body)
	case http.MethodDelete:
		message, status = deleteCharacter(request.Context(), body)
	case http.MethodPut:
		message, status = updateCharacter(request.Context(), body)
	default:
		message = "Method not allowed on Character endpoint"
		status = http.StatusMethodNotAllowed
	}
	writeResponse(writer, status, message)
}

func handleCities(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		message, status = getCity(request.Context(), body)
	case http.MethodPost:
		message, status = addCity(request.Context(), body)
	case http.MethodDelete:
		message, status = deleteCity(request.Context(), body)
	case http.MethodPut:
		message, status = updateCity(request.Context(), body)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)

}

func handleWorlds(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		message, status = getWorld(request.Context(), body)
	case http.MethodPost:
		message, status = addWorld(request.Context(), body)
	case http.MethodDelete:
		message, status = deleteWorld(request.Context(), body)
	case http.MethodPut:
		message, status = updateWorld(request.Context(), body)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)
}

func handleFactions(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		message, status = getFaction(request.Context(), body)
	case http.MethodPost:
		message, status = addFaction(request.Context(), body)
	case http.MethodDelete:
		message, status = deleteFaction(request.Context(), body)
	case http.MethodPut:
		message, status = updateFaction(request.Context(), body)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)
}

func writeResponse(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Set("Content-Type", "application/text")
	writer.Write([]byte(message))
}
