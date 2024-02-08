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
	application.router.HandleFunc("/new", application.new).Methods("POST")
}

func (application *Application) login(w http.ResponseWriter, r *http.Request) {

	var loginModel LoginModel

	err := json.NewDecoder(r.Body).Decode(&loginModel)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	user, err := application.databaseConnect.getUser(loginModel.Email)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	isAuthed := user.auth(loginModel.Password)

	if !isAuthed {
		respondWithJSON(w, http.StatusBadRequest, "Auth Failed")
	}

	token, err := application.tokenManager.createToken(loginModel.Email)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, token)
}

func (application *Application) new(w http.ResponseWriter, r *http.Request) {
	var loginModel LoginModel
	err := json.NewDecoder(r.Body).Decode(&loginModel)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	application.databaseConnect.createUser(loginModel.Email, loginModel.HashPassword())

	respondWithJSON(w, http.StatusCreated, "User Created")
}
