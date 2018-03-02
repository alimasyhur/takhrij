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
	//Route of kitab
	Route{"kitab", "GET", "/kitab/{start:[0-9]+}/{count:[0-9]}", getListKitab},
	Route{"kitab", "GET", "/kitab/{id:[0-9]+}", getKitab},
	Route{"kitab", "POST", "/kitab", createKitab},
	//Route of hadits
	// Route{"hadits", "GET", "/hadits/{id_kitab:[0-9]+}", getHaditses},
	// Route{"hadits", "GET", "/hadits/{id:[0-9]+}/{id_kitab:[0-9]+}", getHadits},
}
