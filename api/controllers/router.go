package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// assign server
var server *Server

// NewRouter ...
func NewRouter() *mux.Router {

	server.App.Router = mux.NewRouter().StrictSlash(true)

	for _, route := range server.App.Routes {

		var handler http.Handler

		// if route.Secure {
		// 	handler = middlewares.AuthMiddleware(route.HandlerFunc)
		// } else {
		// 	handler = route.HandlerFunc
		// }

		handler = route.HandlerFunc
		handler = handlers.LoggingHandler(os.Stderr, handler)

		server.App.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return server.App.Router

}

// HealthCheck checks api health status
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Status: "up",
	})
}
