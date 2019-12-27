package datastore_connectors

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"os"
)

const Database_Connection_String = "DATABASE_CONNECTION_STRING"

func ConnectPostgres() (*sql.DB, error) {
	datasource := os.Getenv(Database_Connection_String)
	if datasource == "" {
		return nil, errors.New("internal error no database connection string")
	}
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
