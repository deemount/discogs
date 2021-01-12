package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// LabelReleasesRepository represents the contract
type LabelReleasesRepository interface {
	FindLabelReleasesByID(id int64, pagination models.PaginationRepository) (*models.LabelReleases, error)
}

// LabelReleasesService is a struct
type LabelReleasesService struct {
	url       string
	uri       string
	useragent string
	token     string
	releases  *models.LabelReleases
}

// NewLabelReleasesService is a constructor
func NewLabelReleasesService(url, uri, useragent, token string) LabelReleasesRepository {
	return &LabelReleasesService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindLabelReleasesByID is a method
func (rs *LabelReleasesService) FindLabelReleasesByID(id int64, pagination models.PaginationRepository) (*models.LabelReleases, error) {

	var err error

	url := fmt.Sprintf("%s%s%v%s", rs.url, rs.uri, id, "/releases")

	// Add token to request headers
	headers := map[string]string{
		"User-Agent":    rs.useragent,
		"Authorization": "Discogs token=" + rs.token,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, pagination.Params(), headers)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &rs.releases)
	return rs.releases, err

}
