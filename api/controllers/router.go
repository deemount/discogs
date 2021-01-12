package controllers

import (
	"net/http"
	"os"

	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// assign server
var server *Server

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secure      bool
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Home",
		Method:      "GET",
		Pattern:     constants.HOMEURI,
		HandlerFunc: server.Home,
		Secure:      false,
	},

	Route{
		Name:        "CheckData",
		Method:      "GET",
		Pattern:     constants.CHECKDATAURI,
		HandlerFunc: server.CheckData,
		Secure:      false,
	},
}

// NewRouter ...
func NewRouter() *mux.Router {

	server.App.Router = mux.NewRouter().StrictSlash(true)
	server.App.Router.NotFoundHandler = http.HandlerFunc(middlewares.NotFound)
	server.App.Router.MethodNotAllowedHandler = http.HandlerFunc(middlewares.NotAllowed)

	for _, route := range routes {

		var handler http.Handler
		if route.Secure {
			handler = middlewares.AuthMiddleware(route.HandlerFunc)
		} else {
			handler = route.HandlerFunc
		}

		handler = handlers.LoggingHandler(os.Stderr, handler)

		server.App.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return server.App.Router
}
