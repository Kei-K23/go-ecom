package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kei-K23/go-ecom/cmd/api"
	"github.com/Kei-K23/go-ecom/config"
	"github.com/Kei-K23/go-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	// init database
	db, err := db.NewDB(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPass,
		Addr:                 config.Env.DBAddr,
		DBName:               config.Env.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initDB(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully initialized database")
}
