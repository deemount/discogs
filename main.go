package main

import (
	"log"

	// _ "github.com/SalvoSP9/discogs/docs"

	"github.com/deemount/discogs/api"
)

// @title Discogs API
// @version 0.1.0
// @description Fetches data from Discogs and stores it in database
// @termsOfService https://github.com/SalvoSP9/discogs/terms/index.html

// @contact.name API Support
// @contact.url https://github.com/deemount
// @contact.email salvatore.gonda@web.de

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8787
// @BasePath /discogs/v2
func main() {

	// assign error
	var err error

	// run application interface
	if err = api.Run(); err != nil {
		log.Fatalf("Discogs API Error %+s", err)
	}

}
