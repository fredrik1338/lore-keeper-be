package database

const (
	// Characters
	SetupCharactersTable = `CREATE TABLE IF NOT EXISTS characters(
		name TEXT PRIMARY KEY,
		description TEXT,
		age INTEGER,
		world TEXT
	)`
	SetupWorldsTable = `CREATE TABLE IF NOT EXISTS worlds(
		name TEXT PRIMARY KEY,
		description TEXT,
		cities TEXT
	)`

	SetupCitiesTable = `CREATE TABLE IF NOT EXISTS cities(
		name TEXT PRIMARY KEY,
		description TEXT,
		founding_date INTEGER,
		notable_characters TEXT,
		factions TEXT
	)`

	SetupFactionsTable = `CREATE TABLE IF NOT EXISTS factions(
		name TEXT PRIMARY KEY,
		description TEXT,
		founding_date INTEGER,
		notable_characters TEXT
	)`

	AddCharacterQuery    = `INSERT INTO characters (name, description, age, world) VALUES ($1, $2, $3, $4)`
	ListCharactersQuery  = `SELECT name FROM characters`
	GetCharacterQuery    = `SELECT * FROM characters WHERE name=$1`
	DeleteCharacterQuery = `DELETE FROM characters WHERE name=$1`
	UpdateCharacterQuery = `UPDATE characters SET
	name=$1, description=$2, age=$3, world=$4 
	WHERE name=$1`

	AddWorldQuery    = `INSERT INTO worlds (name, description, cities) VALUES ($1, $2, $3)`
	ListWorldsQuery  = `SELECT name FROM worlds`
	GetWorldQuery    = `SELECT * FROM worlds WHERE name=$1`
	UpdateWorldQuery = `UPDATE worlds SET
	name=$1, description=$2, cities=$3 
	WHERE name=$1`
	DeleteWorldQuery = `DELETE FROM worlds WHERE name=$1`

	AddCityQuery    = `INSERT INTO cities (name, description, founding_date, notable_characters, factions) VALUES ($1, $2, $3, $4, $5)`
	ListCitiesQuery = `SELECT name FROM cities`
	GetCityQuery    = `SELECT * FROM cities WHERE name=$1`
	UpdateCityQuery = `UPDATE cities SET
	name=$1, description=$2, founding_date=$3, notable_characters=$4, factions=$5
	WHERE name=$1`
	DeleteCityQuery = `DELETE FROM cities WHERE name=$1`

	AddFactionQuery    = `INSERT INTO factions (name, description, founding_date, notable_characters) VALUES ($1, $2, $3, $4)`
	ListFactionsQuery  = `SELECT name FROM factions`
	GetFactionQuery    = `SELECT * FROM factions WHERE name=$1`
	UpdateFactionQuery = `UPDATE factions SET
	name=$1, description=$2, founding_date=$3, notable_characters=$4
	WHERE name=$1`
	DeleteFactionQuery = `DELETE FROM factions WHERE name=$1`
)
