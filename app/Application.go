package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Application struct {
	tokenManager TokenManager
	router       *mux.Router
}

func (application *Application) login(w http.ResponseWriter, r *http.Request) {
	token, err := application.tokenManager.createToken("austinbenerdy")

	validity, err := application.tokenManager.validate(token)
	isValid := strconv.FormatBool(validity)

	if err != nil {
		respondWithJSON(w, 500, err.Error())
	}

	data := []string{"login", string(application.tokenManager.secretKey[:]), token, isValid}
	respondWithJSON(w, 200, data)
}
