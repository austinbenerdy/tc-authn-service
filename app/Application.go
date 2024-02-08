package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
	tokenManager    TokenManager
	router          *mux.Router
	databaseConnect DatabaseConnect
}

func (application *Application) init() {
	application.databaseConnect.testConnect()

	application.tokenManager = TokenManager{}
	application.tokenManager.secretKey = []byte("secret-key")

	application.router = mux.NewRouter()
	application.router.HandleFunc("/login", application.login).Methods("POST")
}

type LoginModel struct {
	Email    string
	Password string
}

func (application *Application) login(w http.ResponseWriter, r *http.Request) {

	var loginModel LoginModel

	err := json.NewDecoder(r.Body).Decode(&loginModel)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	err = application.databaseConnect.authenticate(loginModel.Email, loginModel.Password)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error())
	}

	token, err := application.tokenManager.createToken(loginModel.Email)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, token)
}
