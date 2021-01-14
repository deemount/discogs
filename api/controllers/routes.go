package controllers

import (
	"net/http"

	"github.com/deemount/discogs/api/constants"
)

// initializeRoutes is a method
func (server *Server) initializeRoutes() {

	// regular expression
	id := constants.IDREGEX
	page := constants.PAGEREGEX
	perpage := constants.PERPAGEREGEX
	sort := constants.SORTREGEX
	sortorder := constants.SORTORDERREGEX

	// uri
	home := constants.HOMEURI
	artists := constants.ARTISTSURI
	artistreleases := constants.ARTISTRELEASESURI
	labels := constants.LABELSURI
	labelreleases := constants.LABELRELEASESURI
	masters := constants.MASTERSURI
	masterversions := constants.MASTERVERSIONSURI
	releases := constants.RELEASESURI
	releaserating := constants.RELEASERATINGURI

	//**************** Home Route

	server.App.V2.HandleFunc(home, server.Home).Methods(http.MethodGet)

	//**************** Artists Routes

	// single request
	server.App.V2.HandleFunc(artists, server.GetArtist).Queries("id", id).Methods(http.MethodGet)

	//**************** Artist Releases Routes

	// single request
	server.App.V2.HandleFunc(artistreleases, server.GetArtistReleases).Queries("id", id).Methods(http.MethodGet)
	server.App.V2.HandleFunc(artistreleases, server.GetArtistReleases).Queries("id", id).Queries("page", page).Queries("per_page", perpage).Queries("sort", sort).Queries("sort_order", sortorder).Methods(http.MethodGet)

	//**************** Label Routes

	// single request
	server.App.V2.HandleFunc(labels, server.GetLabel).Queries("id", id).Methods(http.MethodGet)

	//**************** Label Releases Routes

	// single request
	server.App.V2.HandleFunc(labelreleases, server.GetLabelReleases).Queries("id", id).Methods(http.MethodGet)
	server.App.V2.HandleFunc(labelreleases, server.GetLabelReleases).Queries("id", id).Queries("page", page).Queries("per_page", perpage).Queries("sort", sort).Queries("sort_order", sortorder).Methods(http.MethodGet)

	//**************** Master Routes

	// single request
	server.App.V2.HandleFunc(masters, server.GetMaster).Queries("id", id).Methods(http.MethodGet)

	//**************** Master Versions Routes

	// single request
	server.App.V2.HandleFunc(masterversions, server.GetMaster).Queries("id", id).Methods(http.MethodGet)

	//**************** Releases Routes

	// single request
	server.App.V2.HandleFunc(releases, server.GetRelease).Queries("id", id).Methods(http.MethodGet)

	//**************** Release Rating Routes

	// single request
	server.App.V2.HandleFunc(releaserating, server.GetReleaseRating).Queries("id", id).Methods(http.MethodGet)

}
