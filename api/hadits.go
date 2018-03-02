package api

import (
	_ "github.com/go-sql-driver/mysql"
)

// func getHaditses(resp http.ResponseWriter, req *http.Request) {
// 	vars := mux.Vars(req)

// 	data := database.GetHaditsByIDKitab(vars["id_kitab"])

// 	hadits, _ := json.Marshal(data)

// 	resp.Write([]byte(hadits))
// }

// func getHadits(resp http.ResponseWriter, req *http.Request) {
// 	vars := mux.Vars(req)

// 	data := database.GetOneHadits(vars["id"], vars["id_kitab"])

// 	log.Println(data)

// 	hadits, _ := json.Marshal(data)

// 	resp.Write([]byte(hadits))
// }
