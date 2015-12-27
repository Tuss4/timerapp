package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index might not actually be necessary. Just testing out routing
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "This is the Index.")
}

// UserRegister User registration endpoint.
func UserRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO: User Register")
}

// UserLogin User login endpoint.
func UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO: User Login")
}

// UserDetail user detail endpoint.
func UserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, fmt.Sprintf("User ID: %s", userID))
	// stmt, err := db.Prepare(FetchUserQuery)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		fmt.Fprintln(w, "{'detail': 'User not found.'}")
	// 	}
	// }
}
