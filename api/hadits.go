package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alimasyhur/takhrij/model"
	"github.com/gorilla/mux"
)

func getListHaditses(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idKitab, _ := strconv.Atoi(vars["id_kitab"])
	count, _ := strconv.Atoi(vars["count"])
	start, _ := strconv.Atoi(vars["start"])

	haditses, err := model.GetListHadits(idKitab, count, start)
	response, _ := json.Marshal(haditses)

	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Write([]byte(response))
}

func getHadits(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	hadits := model.Hadits{ID: id}

	response, _ := json.Marshal(hadits)
	if err := hadits.Get(); err != nil {
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Write([]byte(response))
}
