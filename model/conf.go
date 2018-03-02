package model

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//Configure database connection of takhrij
func Configure(DB *sqlx.DB) {
	db = DB
}
