package api

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//MusnadAhmad struct kitab musnad Ahmad
type MusnadAhmad struct {
	ID         int    `db:"id"`
	Nass       string `db:"nass"`
	Terjemah   string `db:"terjemah"`
	NassGundul string `db:"nass_gundul"`
	Hno        int    `db:"hno"`
}

func hadits(resp http.ResponseWriter, req *http.Request) {
	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/takhrij")
	if err != nil {
		log.Fatalln(err)
	}

	hadits := MusnadAhmad{}
	err = db.Get(&hadits, "SELECT * FROM Hadits_Musnad_Ahmad LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}

	listHadits, _ := json.Marshal(hadits)
	log.Println(listHadits)

	resp.Write([]byte(listHadits))
}
