package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
