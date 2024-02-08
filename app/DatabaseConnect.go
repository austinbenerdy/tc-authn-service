package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnect struct {
}

func (databaseConnect *DatabaseConnect) testConnect() {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}

type User struct {
	Id       string
	Email    string
	password string
}

type AuthFailed struct{}

func (err *AuthFailed) Error() string {
	return "Auth Failed"
}

func (databaseConnect *DatabaseConnect) authenticate(email string, password string) error {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	results, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", email, password)
	count := 0

	for results.Next() {
		count++
		var user User
		err = results.Scan(&user.Id, &user.Email, &user.password)

		if err != nil {
			panic(err.Error())
		}
	}

	if count > 0 {
		return nil
	}

	return &AuthFailed{}
}
