package api

import "net/http"

//Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes array of Route
type Routes []Route

var routes = Routes{
	Route{"index", "GET", "/", index},
	Route{"kitab", "GET", "/kitab", kitab},
	Route{"hadits", "GET", "/hadits", hadits},
}
