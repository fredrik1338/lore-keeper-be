package database

const (
	CreateTODOsTableSQL = `CREATE TABLE todos (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"description"
	  );`
	GetTODOsSQL   = `SELECT * FROM todos ORDER BY id;`
	AddTODOSQL    = `INSERT INTO todos(title, description) VALUES (?, ?)`
	DeleteTODOSQL = `DELETE FROM todos WHERE ?=id`
)
