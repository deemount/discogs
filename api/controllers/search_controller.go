package controllers

import (
	"net/http"

	"github.com/deemount/discogs/api/repositories"
	"github.com/deemount/discogs/api/responses"
)

// GetSearchResult Controller
func (server *Server) GetSearchResult(w http.ResponseWriter, r *http.Request) {

	// page, perpage, sort, sortorder
	search := repositories.NewSearchService(
		server.App.Discogs.URL,
	)

	responses.JSON(w, http.StatusOK, search)

}
