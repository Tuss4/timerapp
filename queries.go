package main

// CreateUserTable is a query to create the User Table
var CreateUserTable = "CREATE TABLE user (id SERIAL PRIMARY KEY, email varchar(40) NOT NULL UNIQUE, password varchar(30) NOT NULL CHECK (length(password) > 8) , created date DEFAULT current_date);"

// CreateTables creates the tables in the database.
func CreateTables() {}
