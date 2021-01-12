package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// MasterRepository represents the contract between
type MasterRepository interface {
	FindMasterByID(id int64) (*models.Master, error)
}

// MasterService is a struct
type MasterService struct {
	url       string
	uri       string
	useragent string
	token     string
	master    *models.Master
}

// NewMasterService is a object
func NewMasterService(url, uri, useragent, token string) MasterRepository {
	return &MasterService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindMasterByID is a method
func (rs *MasterService) FindMasterByID(id int64) (*models.Master, error) {

	var err error

	url := fmt.Sprintf("%s%s%v", rs.url, rs.uri, id)

	// Add token to request headers
	headers := map[string]string{
		"User-Agent":    rs.useragent,
		"Authorization": "Discogs token=" + rs.token,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, nil, headers)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &rs.master)
	return rs.master, err

}
