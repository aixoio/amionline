package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dblogin string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dblogin)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func Disconnect(db *sql.DB) {
	db.Close()
}
