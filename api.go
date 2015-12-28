package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// BadRequest struct for error message in case of 400
type BadRequest struct {
	Detail string `json:"detail"`
}

// Index might not actually be necessary. Just testing out routing
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "This is the Index.")
}

// UserRegister User registration endpoint.
func UserRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	cu := CreateUser{}
	BR := BadRequest{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&cu)
	if err != nil {
		fmt.Println(err)
	}
	if cu.isConflict(db) {
		w.WriteHeader(http.StatusConflict)
		BR.Detail = "Email must be unique."
		json.NewEncoder(w).Encode(BR)
	} else {
		if cu.isValid() {
			w.WriteHeader(http.StatusCreated)
			nu, err := cu.craete(db)
			if err != nil {
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(nu)
		} else {
			BR.Detail = "Bad request."
			json.NewEncoder(w).Encode(BR)
		}
	}
}

// UserLogin User login endpoint.
func UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO: User Login")
}

// UserDetail user detail endpoint.
func UserDetail(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	user := new(User)
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare(FetchUserQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Server error.")
	}
	err = stmt.QueryRow(userID).Scan(&user.ID, &user.Email, &user.Created)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		NotFound := BadRequest{"User not found."}
		er := json.NewEncoder(w).Encode(NotFound)
		if er != nil {
			panic(er)
		}
	} else {
		er := json.NewEncoder(w).Encode(user)
		if er != nil {
			panic(er)
		}
	}
}
