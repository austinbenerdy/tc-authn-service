package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
	tokenManager TokenManager
	router       *mux.Router
}

func (application *Application) init() {
	application.tokenManager = TokenManager{}
	application.tokenManager.secretKey = []byte("secret-key")

	application.router = mux.NewRouter()
	application.router.HandleFunc("/login", application.login).Methods("POST")
}

func (application *Application) login(w http.ResponseWriter, r *http.Request) {
	token, err := application.tokenManager.createToken("austinbenerdy")

	if err != nil {
		respondWithJSON(w, 500, err.Error())
	}

	respondWithJSON(w, 200, token)
}
