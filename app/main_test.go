package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/auth")
	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	lm := newLoginModel("austin.l.adamson@gmail.com", "test-password")
	hashedPassword := lm.HashPassword()

	insert, err := db.Query("REPLACE INTO users VALUES(?, ?, ?)", 1, lm.Email, hashedPassword)
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

	bodyJson, _ := json.Marshal(lm)

	// Create a request to the /login endpoint
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(bodyJson))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create an instance of your application
	app := Application{}
	app.init()

	// Use the router's ServeHTTP method to handle the request
	app.router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var response string
	err = json.Unmarshal([]byte(rr.Body.String()), &response)

	isValid, err := app.tokenManager.validate(response)

	if isValid != true || err != nil {
		t.Errorf("Token was not a valid JWT")
	}
}

func TestNewHandler(t *testing.T) {
	lm := newLoginModel("newUserTest@gmail.com", "test-password")

	bodyJson, _ := json.Marshal(lm)

	// Create a request to the /login endpoint
	req, err := http.NewRequest("POST", "/new", bytes.NewReader(bodyJson))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create an instance of your application
	app := Application{}
	app.init()

	// Use the router's ServeHTTP method to handle the request
	app.router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
