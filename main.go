package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":5000"

func main() {
	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
	}
}
