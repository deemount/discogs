package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SalvoSP9/discogs/api/constants"
	"github.com/SalvoSP9/discogs/api/models"
	"github.com/SalvoSP9/discogs/api/repositories"
	"github.com/SalvoSP9/discogs/api/responses"
	"github.com/SalvoSP9/discogs/api/utils"
)

// GetLabelReleases Controller
// @Summary Get label releases
// @Description Calling the discogs api and get the label releases by id
// @ID home
// @Accept json
// @Produce json
// @Success 200 {object} model.LabelReleases
// @Header 200 {string} ""
// @Failure 404 {object} http.
// @Router /labels/ [get]
func (server *Server) GetLabelReleases(w http.ResponseWriter, r *http.Request) {

	id, err := utils.ConvertQueryStringToInt64(r)
	if err != nil {
		log.Print(err)
	}

	// e.g. map[page:[] per_page:[] sort:[] sort_order:[]]
	query := r.URL.Query()

	page := query.Get("page")
	if page == "" {
		page = strconv.Itoa(1)
	}
	perpage := query.Get("per_page")
	if perpage == "" {
		perpage = strconv.Itoa(50)
	}
	sort := query.Get("sort")
	if sort == "" {
		sort = "year"
	}
	sortorder := query.Get("sort_order")
	if sortorder == "" {
		sortorder = "asc"
	}

	pagination := models.NewPaginationService(page, perpage, sort, sortorder)

	service := repositories.NewLabelReleasesService(
		server.App.Discogs.URL,
		constants.LABELSURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.LabelReleasesRepository(service)

	releases, err := repository.FindLabelReleasesByID(id, pagination)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, releases)

}
