package sqlite

// import (
// 	"lore-keeper-be/internal/dto"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// const (
// 	testdb = "test.db"
// )

// func TestCreateDB(t *testing.T) {
// 	defer os.Remove(testdb)
// 	db, err := New("test.db")
// 	assert.Nil(t, err)
// 	db.sqlite.Close()
// }

// func TestCRUDTODO(t *testing.T) {
// 	defer os.Remove(testdb)
// 	db, err := New("test.db")
// 	assert.Nil(t, err)
// 	defer db.sqlite.Close()

// 	db.InitDB()
// 	db.AddTODO(dto.TODO{Title: "test", Description: "test"})
// 	todos, err := db.GetTODOs()
// 	assert.NoError(t, err)

// 	assert.Equal(t, 1, len(todos))

// 	success, err := db.DeleteTODO("1")
// 	assert.NoError(t, err)
// 	assert.True(t, success)

// 	success, err = db.DeleteTODO("1")
// 	assert.NoError(t, err)
// 	assert.False(t, success)

// 	todos, err = db.GetTODOs()
// 	assert.NoError(t, err)
// 	assert.Equal(t, 0, len(todos))
// }
