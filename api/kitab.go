package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alimasyhur/takhrij/model"
	"github.com/gorilla/mux"
)

//kitab handler return all list available kitab
func getListKitab(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	count, _ := strconv.Atoi(vars["count"])
	start, _ := strconv.Atoi(vars["start"])

	kitabs, err := model.GetListKitab(start, count)
	response, _ := json.Marshal(kitabs)

	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	resp.Write([]byte(response))
}

//GetKitab handler return one selected kitab by id
func getKitab(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	kitab := model.Kitab{ID: id}
	response, _ := json.Marshal(kitab)
	if err := kitab.Get(); err != nil {
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Write([]byte(response))
}

//createKitab endpoint
func createKitab(resp http.ResponseWriter, req *http.Request) {
	var k model.Kitab
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&k); err != nil {
		errResp := ErrorResp{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		response, _ := json.Marshal(errResp)
		resp.Write([]byte(response))
		return
	}

	defer req.Body.Close()

	if err := k.Create(); err != nil {
		errResp := ErrorResp{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

		response, _ := json.Marshal(errResp)

		resp.Write([]byte(response))
		return
	}

	errResp := ErrorResp{
		Code:    http.StatusOK,
		Message: "Success Create New Kitab",
	}

	response, _ := json.Marshal(errResp)
	resp.Write([]byte(response))
}
