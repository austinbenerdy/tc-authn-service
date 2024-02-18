package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

type DatabaseConnect struct {
	DB *sql.DB
}

func (databaseConnect *DatabaseConnect) TestConnect() {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}

func (databaseConnect *DatabaseConnect) Open() {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	databaseConnect.DB = db
}

func (databaseConnect *DatabaseConnect) Close() {
	databaseConnect.DB.Close()
}

func (databaseConnect *DatabaseConnect) Migrate() {
	databaseConnect.Open()
	defer databaseConnect.Close()

	err := goose.SetDialect("mysql")
	if err != nil {
		panic(err)
	}

	err = goose.Up(databaseConnect.DB, "migrations")
	if err != nil {
		panic(err)
	}
}
