package controllers

import (
	"log"
	"net/http"

	"github.com/SalvoSP9/discogs/api/constants"
	"github.com/SalvoSP9/discogs/api/repositories"
	"github.com/SalvoSP9/discogs/api/responses"
	"github.com/SalvoSP9/discogs/api/utils"
)

// GetRelease Controller
func (server *Server) GetRelease(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	service := repositories.NewReleaseService(
		server.App.Discogs.URL,
		constants.RELEASESURI,
		server.App.Options.Currency,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.ReleaseRepository(service)

	release, err := repository.FindReleaseByID(id)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, release)

}
