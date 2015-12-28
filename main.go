package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = ":5000"

// runServer runs the server
func runServer(db *sql.DB) {
	router.HandleFunc("/", Index)
	router.HandleFunc("/user/login", UserLogin).Methods("POST")
	router.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) { UserRegister(w, r, db) }).Methods("POST")
	router.HandleFunc("/user/{userID}", func(w http.ResponseWriter, r *http.Request) { UserDetail(w, r, db) }).Methods("GET")
	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	db := dbConn()
	err := db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database.")
		log.Println(err)
	}
	switch {
	case os.Args[1] == "runserver":
		runServer(db)
	case os.Args[1] == "createdb":
		createTables(db)
	}
	defer db.Close()
}
