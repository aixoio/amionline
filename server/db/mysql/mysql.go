package mysql

import (
	"database/sql"

	"github.com/aixoio/amionline/logger"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(dblogin string) (*sql.DB, error) {
	logger.Info().Println("Connecting to MySQL")
	db, err := sql.Open("mysql", dblogin)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func Disconnect(db *sql.DB) {
	db.Close()
}
