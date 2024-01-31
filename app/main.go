package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, request *http.Request) {
		io.WriteString(w, "Hello World")
	}

	http.HandleFunc("/", handler)
	log.Println("Listening for requests")
	log.Fatal(http.ListenAndServe(":8000",nil))
}
