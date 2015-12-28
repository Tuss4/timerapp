package main

import (
	"database/sql"
	"log"
)

// CreateUserTable is a query to create the User Table
var CreateUserTable = "CREATE TABLE users_user (id SERIAL PRIMARY KEY, email varchar(40) NOT NULL UNIQUE, password varchar(30) NOT NULL, created date DEFAULT current_date);"

// CreateUserQuery creates a new user
var CreateUserQuery = "INSERT INTO users_user (email, password) VALUES($1, $2);"

// FetchUserQuery returns the user by ID
var FetchUserQuery = "SELECT u.id, u.email, u.created FROM users_user AS u WHERE u.id = $1;"

// FetchUserByEmailQuery returns the user by email
var FetchUserByEmailQuery = "SELECT u.id, u.email, u.created FROM users_user AS u WHERE u.email = $1;"

// UserExistsQuery checks to see if user exists
var UserExistsQuery = "SELECT COUNT(*) FROM users_user WHERE email = $1;"

// createTables creates the tables in the database.
func createTables(conn *sql.DB) {
	q, err := conn.Exec(CreateUserTable)
	if err != nil {
		log.Println(err)
	}
	println(q)
}
