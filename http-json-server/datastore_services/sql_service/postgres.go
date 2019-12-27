package sql_service

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type postgresservice struct {
	db *sql.DB
}

func NewPostgresService(datasource string) (*postgresservice, error) {
	if datasource == "" {
		return nil, errors.New("internal error no database connection string")
	}
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}
	return &postgresservice{db:db}, nil
}


func (p *postgresservice) Query(objecttoscan interface{}, querystring string, args ...string) error {
	rows, err := p.db.Query(querystring, args)
	if err != nil {
		return err
	}
	err = rows.Scan(objecttoscan)
	if err != nil {
		return err
	}
	return nil
}


