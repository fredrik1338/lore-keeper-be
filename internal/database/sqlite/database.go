package sqlite

// import (
// 	"database/sql"
// 	"fmt"
// 	"lore-keeper-be/internal/database"
// 	"lore-keeper-be/internal/dto"
// 	"os"
// 	"strconv"

// 	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
// )

// type Database struct {
// 	sqlite *sql.DB
// }

// func New(path string) (*Database, error) {
// 	os.Remove(path)              // Delete the file to avoid duplicated records.
// 	file, err := os.Create(path) // Create SQLite file
// 	if err != nil {
// 		return nil, err
// 	}
// 	file.Close()

// 	sqliteDatabase, err := sql.Open("sqlite3", fmt.Sprintf("./%s", path))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Database{
// 		sqlite: sqliteDatabase,
// 	}, nil
// }

// func (db *Database) InitDB() error {
// 	statement, err := db.sqlite.Prepare(database.CreateTODOsTableSQL)
// 	if err != nil {
// 		return err
// 	}
// 	statement.Exec()
// 	return nil
// }

// func (db *Database) GetTODOs() ([]dto.TODO, error) {
// 	var todos []dto.TODO
// 	statement, err := db.sqlite.Prepare(database.GetTODOsSQL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err := statement.Query()
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		todo := dto.TODO{}
// 		rows.Scan(&todo.ID, &todo.Title, &todo.Description)
// 		todos = append(todos, todo)
// 	}
// 	return todos, nil
// }

// func (db *Database) AddTODO(todo dto.TODO) error {
// 	statement, err := db.sqlite.Prepare(database.AddTODOSQL)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = statement.Exec(todo.Title, todo.Description)
// 	if err != nil {
// 		return err

// 	}
// 	return nil
// }

// func (db *Database) DeleteTODO(id string) (bool, error) {
// 	statement, err := db.sqlite.Prepare(database.DeleteTODOSQL)
// 	if err != nil {
// 		return false, err
// 	}
// 	nr, err := strconv.Atoi(id)
// 	if err != nil {
// 		return false, err

// 	}

// 	result, err := statement.Exec(nr)
// 	if err != nil {
// 		return false, err

// 	}
// 	affectedRows, err := result.RowsAffected()
// 	if err != nil {
// 		return true, err
// 	}
// 	if affectedRows > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func (db *Database) Shutdown() error {
// 	return db.sqlite.Close()
// }
