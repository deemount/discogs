package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/repositories"
	"github.com/deemount/discogs/api/responses"
	"github.com/deemount/discogs/api/utils"
)

// GetMaster Controller
func (server *Server) GetMaster(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	service := repositories.NewMasterService(
		server.App.Discogs.URL,
		constants.MASTERSURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.MasterRepository(service)

	releases, err := repository.FindMasterByID(id)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, releases)

}
