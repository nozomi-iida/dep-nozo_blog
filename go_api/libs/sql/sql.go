package sql

import "database/sql"

func Init() (*sql.DB, error) {
	db := sql.DB{}
	return &db, nil	
}
