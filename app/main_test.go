package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	// Create a request to the /login endpoint
	req, err := http.NewRequest("POST", "/login", nil)
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
