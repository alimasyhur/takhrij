package api

import (
	"encoding/json"
	"net/http"

	"github.com/alimasyhur/takhrij/database"
	_ "github.com/go-sql-driver/mysql"
)

func hadits(resp http.ResponseWriter, req *http.Request) {
	data := database.GetOneHadits()

	listHadits, _ := json.Marshal(data)

	resp.Write([]byte(listHadits))
}
