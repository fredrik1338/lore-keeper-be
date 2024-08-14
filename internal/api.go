package internal

import (
	"io"
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

func (api Server) handleCharacters(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := io.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		// Check for a specific list parameter
		if _, ok := request.URL.Query()["list"]; ok {
			message, status = listCharacters(request.Context(), api.db)
		} else {
			message, status = getCharacter(request.Context(), body, api.db)
		}
	case http.MethodPost:
		message, status = addCharacter(request.Context(), body, api.db)
	case http.MethodDelete:
		message, status = deleteCharacter(request.Context(), body, api.db)
	case http.MethodPut:
		message, status = updateCharacter(request.Context(), body, api.db)
	default:
		message = "Method not allowed on Character endpoint"
		status = http.StatusMethodNotAllowed
	}
	writeResponse(writer, status, message)
}

func (api Server) handleCities(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := io.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		// Check for a specific list parameter
		if _, ok := request.URL.Query()["list"]; ok {
			message, status = listCities(request.Context(), api.db)
		} else {
			message, status = getCity(request.Context(), body, api.db)
		}
	case http.MethodPost:
		message, status = addCity(request.Context(), body, api.db)
	case http.MethodDelete:
		message, status = deleteCity(request.Context(), body, api.db)
	case http.MethodPut:
		message, status = updateCity(request.Context(), body, api.db)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)

}

func (api Server) handleWorlds(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := io.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		// Check for a specific list parameter
		if _, ok := request.URL.Query()["list"]; ok {
			message, status = listWorlds(request.Context(), api.db)
		} else {
			message, status = getWorld(request.Context(), body, api.db)
		}
	case http.MethodPost:
		message, status = addWorld(request.Context(), body, api.db)
	case http.MethodDelete:
		message, status = deleteWorld(request.Context(), body, api.db)
	case http.MethodPut:
		message, status = updateWorld(request.Context(), body, api.db)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)
}

func (api Server) handleFactions(writer http.ResponseWriter, request *http.Request) {
	var message string
	var status int

	body, err := io.ReadAll(request.Body)
	if err != nil {
		message = "Could not read request"
		status = http.StatusBadRequest
	}

	switch request.Method {
	case http.MethodGet:
		// Check for a specific list parameter
		if _, ok := request.URL.Query()["list"]; ok {
			message, status = listFactions(request.Context(), api.db)
		} else {
			message, status = getFaction(request.Context(), body, api.db)
		}
	case http.MethodPost:
		message, status = addFaction(request.Context(), body, api.db)
	case http.MethodDelete:
		message, status = deleteFaction(request.Context(), body, api.db)
	case http.MethodPut:
		message, status = updateFaction(request.Context(), body, api.db)
	default:
		io.WriteString(writer, "Method not allowed")
	}
	writeResponse(writer, status, message)
}

func writeResponse(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Set("Content-Type", "application/text")
	writer.Write([]byte(message))
}
