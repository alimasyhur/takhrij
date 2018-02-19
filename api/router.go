package api

import (
	"net/http"

	"github.com/alimasyhur/takhrij/api/helpers"
	"github.com/gorilla/mux"
)

//NewRouter handle and register available router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = helpers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
