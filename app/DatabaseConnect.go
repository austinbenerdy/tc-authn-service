package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	Password string
}

func (u *User) auth(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type AuthFailed struct{}

func (err *AuthFailed) Error() string {
	return "Auth Failed"
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
	}, &AuthFailed{}
}
