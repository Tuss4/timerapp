package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DatabaseURL ENV
var DatabaseURL = os.Getenv("DATABASE_URL")

func dbConn() *sql.DB {
	dbURL := DatabaseURL
	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://postgres:postgres@%s:5432/postgres", os.Getenv("DB_PORT_5432_TCP_ADDR"))
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
