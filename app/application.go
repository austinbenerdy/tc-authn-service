package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tinycloudtv/authn-service/app/internal/models"
	"github.com/tinycloudtv/authn-service/app/internal/repositories"
	"net/http"
)

type Application struct {
	tokenManager    TokenManager
	router          *mux.Router
	usersRepository repositories.UserRepository
}

func (application *Application) init() {
	application.usersRepository.Init()

	application.tokenManager = TokenManager{}
	application.tokenManager.secretKey = []byte("secret-key")

	application.router = mux.NewRouter()
	application.router.HandleFunc("/login", application.login).Methods("POST")
	application.router.HandleFunc("/new", application.new).Methods("POST")
	application.router.HandleFunc("/validate", application.new).Methods("POST")
	application.router.HandleFunc("/", application.home).Methods("GET")
}

func (application *Application) login(w http.ResponseWriter, r *http.Request) {

	var loginModel models.LoginModel

	err := json.NewDecoder(r.Body).Decode(&loginModel)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	user, err := application.usersRepository.GetUser(loginModel.Email)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	isAuthed := user.Auth(loginModel.Password)

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
	var loginModel models.LoginModel
	err := json.NewDecoder(r.Body).Decode(&loginModel)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
	}

	application.usersRepository.CreateUser(loginModel.Email, loginModel.HashPassword())

	respondWithJSON(w, http.StatusCreated, "User Created")
}

func (application *Application) validate(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		respondWithJSON(w, http.StatusUnauthorized, "")
	}
	tokenString = tokenString[len("Bearer "):]

	validate, err := application.tokenManager.validate(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, "")
	}

	if validate {
		respondWithJSON(w, http.StatusOK, "")
	}

	respondWithJSON(w, http.StatusUnauthorized, "")
}

func (application *Application) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	type homepage struct {
		Page string
	}

	payload, _ := json.Marshal(&homepage{
		Page: "Homepage",
	})

	w.Write(payload)
}
