package api

import (
	"log"
	"net/http"
)

//Start Server
func Start() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
