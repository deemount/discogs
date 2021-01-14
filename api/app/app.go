package app

import (
	"github.com/gorilla/mux"

	"github.com/deemount/discogs/api/config"
	"github.com/deemount/discogs/api/config/driver"
)

// App ...
type App struct {
	// refer
	DB     driver.DataService
	Routes config.Routes

	// pointer
	API     *config.API
	Discogs *config.Discogs
	Options *config.Options
	Swagger *config.Swagger

	// legal
	Router *mux.Router
	V2     *mux.Router
}
