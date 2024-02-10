package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/tinycloudtv/authn-service/app/internal/errors"
	"github.com/tinycloudtv/authn-service/app/internal/models"
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

func (databaseConnect *DatabaseConnect) getUser(email string) (models.User, error) {
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
		var user models.User
		err = results.Scan(&user.Id, &user.Email, &user.Password)

		return user, nil
	}

	return models.User{}, &errors.AuthFailedError{}
}

func (database *DatabaseConnect) createUser(email string, password string) models.User {
	id := uuid.New()

	user := models.User{
		Id:       id.String(),
		Email:    email,
		Password: password,
	}

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	insert, err := db.Query("REPLACE INTO users VALUES(?, ?, ?)", user.Id, user.Email, user.Password)
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

	return user
}
