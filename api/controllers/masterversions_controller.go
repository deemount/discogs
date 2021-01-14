package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/deemount/discogs/api/constants"
	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/repositories"
	"github.com/deemount/discogs/api/responses"
	"github.com/deemount/discogs/api/utils"
)

// GetMasterVersions Controller
func (server *Server) GetMasterVersions(w http.ResponseWriter, r *http.Request) {

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

	service := repositories.NewMasterVersionsService(
		server.App.Discogs.URL,
		constants.MASTERSURI,
		server.App.Discogs.UserAgent,
		server.App.Discogs.Token)

	repository := repositories.MasterVersionsRepository(service)

	releases, err := repository.FindMasterVersionsByID(id, pagination)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, releases)

}
