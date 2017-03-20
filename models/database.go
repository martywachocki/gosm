package models

import (
	"database/sql"
	"encoding/json"
	"os"

	"io/ioutil"

	"github.com/jmoiron/sqlx"
	// Add sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	databasePath = "data/gosm.db"
	setupFile    = "data/setup.sql"
)

// Database The sqlite3 database
var Database *sqlx.DB

// Connect Connects to the sqlite3 database, and creates the database if it does not already exist
func Connect() {
	var needsSetup = false
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		needsSetup = true
		_, err := os.Create(databasePath)
		if err != nil {
			panic(err)
		}
	}
	database, err := sqlx.Open("sqlite3", databasePath)
	if err != nil {
		panic(err)
	}
	Database = database
	if needsSetup {
		setupFile, err := ioutil.ReadFile(setupFile)
		_, err = Database.Exec(string(setupFile))
		if err != nil {
			panic(err)
		}
	}
}

type jsonNullInt64 struct {
	sql.NullInt64
}

func (v jsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *jsonNullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}
