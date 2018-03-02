package api

import (
	"encoding/json"
	"net/http"
)

//ErrorResp struct returns code and message of error
type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func responseWithMessage(resp http.ResponseWriter, code int, message string) {
	responseWithJSON(resp, code, ErrorResp{
		Code:    code,
		Message: message,
	})
}

func responseWithJSON(resp http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	resp.Write(response)
}
