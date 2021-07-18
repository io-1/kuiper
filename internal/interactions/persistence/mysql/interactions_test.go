// +build !unit,integration

package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db *MysqlPersistence
)

func Test_InteractionDetails(t *testing.T) {
	var (
		id = "e83fcdcf-6c94-4b9a-9001-6ba0380a814b"
	)

	_, err := db.GetInteractionDetails(id)
	assert.NoError(t, err)
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
