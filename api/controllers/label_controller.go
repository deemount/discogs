package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/repositories"
	"github.com/deemount/discogs/api/responses"
	"github.com/deemount/discogs/api/utils"
)

// GetLabel Controller
// @Summary Get label
// @Description Calling the discogs api and get the label by id
// @ID home
// @Accept json
// @Produce json
// @Success 200 {object} model.Label
// @Header 200 {string} ""
// @Failure 404 {object} http.
// @Router /labels/ [get]
func (server *Server) GetLabel(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	service := repositories.NewLabelService(
		server.App.Discogs.URL,
		constants.LABELSURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.LabelRepository(service)

	label, err := repository.FindLabelByID(id)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, label)

}
