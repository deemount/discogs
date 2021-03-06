package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/repositories"
	"github.com/deemount/discogs/api/responses"
	"github.com/deemount/discogs/api/utils"
)

// GetReleaseRating Controller
func (server *Server) GetReleaseRating(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	service := repositories.NewReleaseRatingService(
		server.App.Discogs.URL,
		constants.RELEASESURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.ReleaseRatingRepository(service)

	rating, err := repository.FindReleaseRatingByID(id)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, rating)

}
