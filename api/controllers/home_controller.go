package controllers

import (
	"net/http"

	"github.com/deemount/discogs/api/responses"
)

// Home Controller
// @Summary Get home
// @Description Homepage
// @ID home
// @Accept json
// @Produce json
// @Header 200 {string} Token "ok"
// @Failure 404 {object} utils.HTTPError404
// @Router /home [get]
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Reporting Distributor API v2")
}
