package main

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"
)

// User model struct
type User struct {
	ID      int       `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

// CreateUser struct for creating user
type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cu *CreateUser) isValid() bool {
	switch {
	case !cu.isEmailValid() && !cu.isPassValid():
		print("Email and pass")
		return false
	case !cu.isEmailValid():
		print("Email")
		return false
	case !cu.isPassValid():
		print("Password")
		return false
	default:
		return true
	}
}

func (cu *CreateUser) isPassValid() bool {
	if len(cu.Password) < 8 {
		return false
	}
	return true
}

func (cu *CreateUser) isEmailValid() bool {
	validEmail := regexp.MustCompile(`.+@.+`)
	switch {
	case cu.Email == "":
		return false
	case !validEmail.MatchString(cu.Email):
		return false
	default:
		return true
	}
}

func (cu *CreateUser) craete(db *sql.DB) (User, error) {
	u := User{}
	e := errors.New("")
	stmt, err := db.Prepare(CreateUserQuery)
	if err != nil {
		e = err
	}
	res, err := stmt.Exec(cu.Email, cu.Password)
	if err != nil {
		e = err
	}
	fmt.Println(res)
	stmt, err = db.Prepare(FetchUserByEmailQuery)
	if err != nil {
		e = err
	}
	err = stmt.QueryRow(cu.Email).Scan(&u.ID, &u.Email, &u.Created)
	if err != nil {
		e = err
	}
	return u, e
}
