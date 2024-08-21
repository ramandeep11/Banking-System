package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/db/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if(err != nil) {
		log.Fatal("connot load config")
	}
	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to the db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
