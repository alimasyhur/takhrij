package main

import (
	"log"

	"github.com/alimasyhur/takhrij/api"
	database "github.com/alimasyhur/takhrij/database"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/takhrij?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	// defer db.Close()

	database.Configure(db)

	api.Start()
}
