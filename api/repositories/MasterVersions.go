package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// MasterVersionsRepository represents the contract
type MasterVersionsRepository interface {
	FindMasterVersionsByID(id int64, pagination models.PaginationRepository) (*models.MasterVersions, error)
}

// MasterVersionsService is a struct
type MasterVersionsService struct {
	url       string
	uri       string
	useragent string
	token     string
	versions  *models.MasterVersions
}

// NewMasterVersionsService is a constructor
func NewMasterVersionsService(url, uri, useragent, token string) MasterVersionsRepository {
	return &MasterVersionsService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindMasterVersionsByID is a method
func (rs *MasterVersionsService) FindMasterVersionsByID(id int64, pagination models.PaginationRepository) (*models.MasterVersions, error) {

	var err error

	url := fmt.Sprintf("%s%s%v%s", rs.url, rs.uri, id, "/versions")

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

	err = json.Unmarshal(body, &rs.versions)
	return rs.versions, err

}
