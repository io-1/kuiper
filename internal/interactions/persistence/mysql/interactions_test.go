// +build !unit,integration

package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
)

var (
	db *MysqlPersistence
)

func Test_InteractionDetails(t *testing.T) {
	var (
		id = uuid.New().String()
	)

	_, err := db.GetInteractionDetails(id)
	// assert.Equal(t, rowsAffectedExpected, rowsAffectedActual)
}

func TestMain(m *testing.M) {
	var (
		err error
	)

	dbConn := os.Getenv("DB_CONN")
	db, err = NewMysqlPersistence(dbConn)
	if err != nil {
		log.Fatal(err.Error())
	}
	code := m.Run()
	os.Exit(code)
}
