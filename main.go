package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = ":5000"

// runServer runs the server
func runServer() {
	router.HandleFunc("/", Index)
	router.HandleFunc("/user/login", UserLogin)
	router.HandleFunc("/user/register", UserRegister)
	router.HandleFunc("/user/{userID}", UserDetail)
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
		runServer()
	case os.Args[1] == "createdb":
		createTables(db)
	}
}
