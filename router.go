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
	Queries     [2]string
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		Name:        "GetTourPr",
		Method:      "GET",
		Pattern:     "/tournaments/{id}",
		HandlerFunc: hb.GetTournamentPronos,
	},
	Route{
		Name:        "GetTourPrUser",
		Method:      "GET",
		Pattern:     "/tournaments",
		HandlerFunc: hb.GetTournamentPronosByUser,
		Queries:     [2]string{"user", "{user}"},
	},
	Route{
		Name:        "AddTourn",
		Method:      "POST",
		Pattern:     "/tournaments",
		HandlerFunc: hb.AddTournamentPronos,
	},
	Route{
		Name:        "AddProno",
		Method:      "POST",
		Pattern:     "/tournaments/{id_tp}/prono",
		HandlerFunc: hb.AddProno,
	},
}

var adminRoutes = Routes{
	Route{
		Name:        "AdmAddTourn",
		Method:      "POST",
		Pattern:     "/tournaments",
		HandlerFunc: hb.AddTournament,
	},
	Route{
		Name:        "AdmGetTourn",
		Method:      "GET",
		Pattern:     "/tournaments/{id}",
		HandlerFunc: hb.GetTournament,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		if route.Queries[0] != "" {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler).
				Queries(route.Queries[0], route.Queries[1])
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}

	}
	s := router.PathPrefix("/admin").Subrouter()

	for _, route := range adminRoutes {
		var handler http.Handler
		handler = route.HandlerFunc

		s.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
