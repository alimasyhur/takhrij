package main

import (
	"log"

	"github.com/alimasyhur/takhrij/api"
	"github.com/alimasyhur/takhrij/model"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/takhrij?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	model.Configure(db)

	api.Start()
}
