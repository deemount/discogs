package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/deemount/discogs/api/config"
	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/controllers"
)

// server is a global var
var server = controllers.Server{}

// Run is a func
// returns nil on success
func Run() error {

	// assign error
	var err error

	// Notice:
	// Put in comments when in production mode
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	//
	version, _ := strconv.Atoi(os.Getenv("API_DISCOGS_VERSION"))
	limitClient, _ := strconv.Atoi(os.Getenv("API_SERVER_CLIENT_LIMIT"))
	singularTable, _ := strconv.ParseBool(os.Getenv("API_DB_SINGULARTABLE"))
	logMode, _ := strconv.ParseBool(os.Getenv("API_DB_LOGMODE"))

	server.App.DB.Config = &config.DB{
		Driver:        os.Getenv("API_DB_DRIVER"),
		User:          os.Getenv("API_DB_USER"),
		PW:            os.Getenv("API_DB_PASSWORD"),
		Port:          os.Getenv("API_DB_PORT"),
		Host:          os.Getenv("API_DB_HOST"),
		SSL:           os.Getenv("API_DB_SSLMODE"),
		Schema:        os.Getenv("API_DB_SCHEMA"),
		TblPrefix:     os.Getenv("API_DB_TABLE_PREFIX"),
		Name:          os.Getenv("API_DB_NAME"),
		SingularTable: singularTable,
		LogMode:       logMode,
	}

	server.App.API = &config.API{
		Version:     version,
		Host:        os.Getenv("API_SERVER_HOST"),
		Path:        os.Getenv("API_SERVER_PATH_PREFIX"),
		Port:        os.Getenv("API_SERVER_PORT"),
		Destination: os.Getenv("API_SERVER_PATH_SRC"),
		LimitClient: limitClient,
	}

	server.App.Swagger = &config.Swagger{
		Host: os.Getenv("API_SWAGGER_HOST"),
		Port: os.Getenv("API_SWAGGER_PORT"),
	}

	// Initialize routes
	server.App.Routes = config.Routes{
		config.Route{
			Name:        "Home",            // Name
			Method:      http.MethodGet,    // HTTP method
			Pattern:     constants.HOMEURI, // Route pattern
			HandlerFunc: server.Home,       // Handler function
		},
		config.Route{
			Name:        "HealthCheck",
			Method:      http.MethodGet,
			Pattern:     constants.HEALTHURI,
			HandlerFunc: server.HealthCheck,
		},
	}

	// initialize orm idle and router
	if err = server.Initialize(); err != nil {
		return err
	}

	// run server
	server.Run()

	return nil

}
