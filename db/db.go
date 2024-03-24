package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	return db, nil
}
