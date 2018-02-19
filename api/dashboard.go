package api

import "net/http"

func index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Welcome to Takhrif Apps"))
}
