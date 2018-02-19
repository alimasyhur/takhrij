package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//Start Server
func Start() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
