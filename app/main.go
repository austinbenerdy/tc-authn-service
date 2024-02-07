package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	app := Application{}
	app.tokenManager = TokenManager{}
	app.tokenManager.secretKey = []byte("secret-key")

	app.router = mux.NewRouter()
	app.router.HandleFunc("/login", app.login).Methods("GET")

	log.Println("Listening for requests")
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", app.router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}