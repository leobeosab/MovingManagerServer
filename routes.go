//Manages all the routes for the web server
package main

import "github.com/gorilla/mux"

//Routes : Slice of routes for the server handler
type Routes []Route

//CreateRouter : Returns a Mux Router
func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

//Set our routes valuesm, functions are in routehandler.go
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Add Move",
		"PUT",
		"/moves/add",
		AddMove,
	},
	Route{
		"Delete Move",
		"DELETE",
		"/moves/delete/{id}",
		DeleteMove,
	},
}
