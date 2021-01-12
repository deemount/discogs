package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// ArtistReleasesRepository represents the contract
type ArtistReleasesRepository interface {
	FindArtistReleasesByID(id int64, pagination models.PaginationRepository) (*models.ArtistReleases, error) // , pagination models.PaginationRepository
}

// ArtistReleasesService is a struct
type ArtistReleasesService struct {
	url       string
	uri       string
	useragent string
	token     string
	releases  *models.ArtistReleases
}

// NewArtistReleasesService is a constructor
func NewArtistReleasesService(url, uri, useragent, token string) ArtistReleasesRepository {
	return &ArtistReleasesService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindArtistReleasesByID is a method
func (rs *ArtistReleasesService) FindArtistReleasesByID(id int64, pagination models.PaginationRepository) (*models.ArtistReleases, error) {

	var err error

	url := fmt.Sprintf("%s%s%v%s", rs.url, rs.uri, id, "/releases")

	// Add token to request headers
	headers := map[string]string{
		"User-Agent":    rs.useragent,
		"Authorization": "Discogs token=" + rs.token,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, pagination.Params(), headers) // pagination.Params()
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &rs.releases)
	return rs.releases, err

}
