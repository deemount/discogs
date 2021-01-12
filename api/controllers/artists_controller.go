package controllers

import (
	"log"
	"net/http"

	"github.com/SalvoSP9/discogs/api/constants"
	"github.com/SalvoSP9/discogs/api/repositories"
	"github.com/SalvoSP9/discogs/api/responses"
	"github.com/SalvoSP9/discogs/api/utils"
)

// GetArtist Controller
// @Summary Get artist by query
// @Description Calling the discogs api and get the artist by id
// @ID get-artist
// @Accept json
// @Produce json
// @Success 200 {object} models.Artist
// @Header 200 {string} Token "ok"
// @Failure 404 {object} http.
// @Router /artists/ [get]
func (server *Server) GetArtist(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	service := repositories.NewArtistService(
		server.App.Discogs.URL,
		constants.ARTISTSURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.ArtistRepository(service)

	artist, err := repository.FindArtistByID(id)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, artist)

}
