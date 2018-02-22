package database

import (
	"log"

	model "github.com/alimasyhur/takhrij/model"
)

//GetOneHadits query to get on available hadits
func GetOneHadits() (hadits *model.MusnadAhmad) {
	err := db.Get(&hadits, "SELECT * FROM Hadits_Musnad_Ahmad LIMIT 1")

	if err != nil {
		log.Fatalln(err)
	}

	return hadits
}
