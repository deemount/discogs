package app

import (
	"github.com/deemount/discogs/api/config"
	"github.com/deemount/discogs/api/config/driver"
	"github.com/gorilla/mux"
)

// App ...
type App struct {
	// refer
	DB     driver.DataService
	Routes config.Routes

	// pointer
	API     *config.API
	Options *config.Options
	Swagger *config.Swagger

	// legal
	Router *mux.Router
	V2     *mux.Router
}
