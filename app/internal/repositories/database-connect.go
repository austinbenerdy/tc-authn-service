package repositories

import (
	"database/sql"
	"fmt"
)

type DatabaseConnect struct {
	DB *sql.DB
}

func (databaseConnect *DatabaseConnect) testConnect() {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}

func (databaseConnect *DatabaseConnect) Open() {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}

	databaseConnect.DB = db
}

func (databaseConnect *DatabaseConnect) Close() {
	databaseConnect.DB.Close()
}
