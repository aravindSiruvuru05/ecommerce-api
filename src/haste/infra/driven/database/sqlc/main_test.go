package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDBConn *sql.DB

func TestMain(m *testing.M) {
	DBDriver := "postgres"
	DBSource := "postgresql://user:1234@localhost:5432/haste_db?sslmode=disable"

	var err error
	testDBConn, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	} else {
		log.Println("Successfully connected to db.")
	}

	testQueries = New(testDBConn)
	os.Exit(m.Run())
}
