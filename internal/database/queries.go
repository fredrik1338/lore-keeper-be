package database

const (
	// Characters
	setupCharactersTable = `CREATE TABLE IF NOT EXISTS characters(
		name TEXT PRIMARY KEY,
		age INTEGER,
		world TEXT
	)`
	setupWorldsTable = `CREATE TABLE IF NOT EXISTS worlds(
		name TEXT PRIMARY KEY,
		description TEXT,
		cities TEXT[]
	)`

	setupCitiesTable = `CREATE TABLE IF NOT EXISTS cities(
		name TEXT PRIMARY KEY,
		description TEXT,
		founding_date INTEGER,
		notable_characters TEXT[],
		factions TEXT[]
	)`

	setupFactionsTable = `CREATE TABLE IF NOT EXISTS factions(
		name TEXT PRIMARY KEY,
		description TEXT,
		founding_date INTEGER,
		notable_characters TEXT[]
	)`

	addPersonQuery    = `INSERT INTO characters (name, age, world) VALUES ($1, $2, $3)`
	getPersonQuery    = `SELECT * FROM characters WHERE name=$1`
	deletePersonQuery = `DELETE * FROM characters WHERE name=$1`
	updatePersonQuery = `UPDATE characters SET
	name=$1, age=$2, world=$3 
	WHERE name=$1`

	addWorldQuery    = `INSERT INTO worlds (name, description, cities) VALUES ($1, $2, $3)`
	getWorldQuery    = `SELECT * FROM worlds WHERE name=$1`
	updateWorldQuery = `UPDATE worlds SET
	name=$1, description=$2, cities=$3 
	WHERE name=$1`

	addCityQuery    = `INSERT INTO cities (name, description, founding_data, notable_characters, factions) VALUES ($1, $2, $3, $4, $5)`
	getCityQuery    = `SELECT * FROM cities WHERE name=$1`
	updateCityQuery = `UPDATE cities SET
	name=$1, description=$2, founding_data=$3, notable_characters=$4, factions=$5
	WHERE name=$1`
	deleteCityQuery = `DELETE * FROM cities WHERE name=$1`

	addFactionQuery    = `INSERT INTO factions (name, description, cities) VALUES ($1, $2, $3)`
	getFactionQuery    = `SELECT from factions WHERE name=$1`
	updateFactionQuery = `UPDATE factions SET
	name=$1, description=$2, founding_data=$3, notable_characters=$4, leader=$5
	WHERE name=$1`
	deleteFactionQuery = `DELETE * FROM factions WHERE name=$1`
)
