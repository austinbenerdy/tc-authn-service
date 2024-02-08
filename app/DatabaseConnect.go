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

func (databaseConnect *DatabaseConnect) getUser(email string) (User, error) {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	results, err := db.Query("SELECT * FROM users WHERE email = ?", email)

	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Email, &user.Password)

		return user, nil
	}

	return User{
		Id:       "",
		Email:    "",
		Password: "",
	}, &AuthFailedError{}
}
