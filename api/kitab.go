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
		responseWithMessage(resp, http.StatusBadRequest, err.Error())
		return
	}

	defer req.Body.Close()

	if err := k.Create(); err != nil {
		responseWithMessage(resp, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithMessage(resp, http.StatusOK, "Kitab Successfully Created")
}

//updateKitab endpoint
func updateKitab(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithMessage(resp, http.StatusBadRequest, err.Error())
		return
	}

	var k model.Kitab
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&k); err != nil {
		responseWithMessage(resp, http.StatusBadRequest, err.Error())
		return
	}

	defer req.Body.Close()
	k.ID = id

	if err := k.Update(); err != nil {
		responseWithMessage(resp, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithMessage(resp, http.StatusOK, "Kitab "+k.Name+" Successfully Updated")
}

//deleteKitab endpoint
func deleteKitab(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithMessage(resp, http.StatusBadRequest, err.Error())
		return
	}
	defer req.Body.Close()
	k := model.Kitab{ID: id}

	if err := k.Delete(); err != nil {
		responseWithMessage(resp, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithMessage(resp, http.StatusOK, "Kitab "+k.Name+" Successfully Deleted")
}

//getArchivedKitab handler return all list available kitab
func getArchivedKitab(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	count, _ := strconv.Atoi(vars["count"])
	start, _ := strconv.Atoi(vars["start"])

	kitabs, err := model.GetArchivedKitab(start, count)
	response, _ := json.Marshal(kitabs)

	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	resp.Write([]byte(response))
}
