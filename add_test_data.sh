#!/bin/bash

# Universe
curl -X POST http://localhost:8080/api/v1/lore-keeper/universes -d '{"name": "Lindas universum", "description": "Lindas universum"}'

# Character
curl -X POST http://localhost:8080/api/v1/lore-keeper/characters -d '{"name": "Linda", "description": "A baaaaad biish", "age": 32, "home": "earth"}'

# Faction
curl -X POST http://localhost:8080/api/v1/lore-keeper/factions -d '{"name": "Lindas crew", "description": "Lindas crew", "notableCharacters": ["David", "Mikaela"], "foundingDate": "0"}'

# City
curl -X POST http://localhost:8080/api/v1/lore-keeper/cities -d '{"name": "Lindsville", "notableCharacters": ["David", "Mikaela"], "factions": ["name"], "description": "Lindas crew", "foundingDate": "väldigt länge sedan"}'

# World
curl -X POST http://localhost:8080/api/v1/lore-keeper/worlds -d '{"name": "Lindas värld", "description": "Lindas värld", "cities": ["Lindsville"]}'