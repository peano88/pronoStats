package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"GetTourPr",
		"GET",
		"/tournament/{id}",
		hb.GetTournamentPronos,
	},
	Route{
		"AddTourn",
		"POST",
		"/tournament",
		hb.AddTournamentPronos,
	},
	Route{
		"AddProno",
		"POST",
		"/tournament/{id_tp}/prono",
		hb.AddTournamentPronos,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
