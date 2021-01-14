package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres dialect
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/deemount/discogs/api/app"
	"github.com/deemount/discogs/api/config/driver"
	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/middlewares"
)

// Server is a struct
type Server struct {
	App app.App
}

// Initialize is a method
// @Summary init db connection and set router
// @Description initialize database connection and set multiplexer router
// @ID init-db-and-set-router
func (server *Server) Initialize() error {

	var err error

	db := driver.NewDataService(*server.App.DB.Config)
	idle, err := db.Connect()
	if err != nil {
		log.Printf("Could not open database connection: %v", err)
	}

	log.Print(idle)

	// set new router instance
	server.App.Router = mux.NewRouter()

	// build swagger ui
	server.App.Router.PathPrefix(constants.SWAGGERURI).
		Handler(
			httpSwagger.Handler(
				httpSwagger.URL(server.App.Swagger.Host+":"+server.App.Swagger.Port+"/swagger/doc.json"), // The url pointing to API definition
				httpSwagger.DeepLinking(true),
				httpSwagger.DocExpansion("none"),
				httpSwagger.DomID("#swagger-ui"),
			))

	// Register new routes with matcher for path
	server.App.V2 = server.App.Router.PathPrefix(server.App.API.Path).Subrouter()
	server.App.V2.Use(middlewares.JSON)

	server.initializeRoutes()

	return err

}

// Run calls listen-and-serve and implements logging handler
// @Summary Runs the listener on tcp and serves handler for incoming connections
// @Description run listen and serve on given port
// @ID run-listen-and-serve-on-give-port
func (server *Server) Run() {

	go log.Printf("Discogs API v%d is ready to listen and serve on port %s", server.App.API.Version, server.App.API.Port)
	log.Fatal(http.ListenAndServe(":"+server.App.API.Port, handlers.LoggingHandler(os.Stdout, server.App.Router)))

}
