package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// ReleaseRepository represents the contract between
type ReleaseRepository interface {
	FindReleaseByID(id int64) (*models.Release, error)
}

// ReleaseService is a struct
type ReleaseService struct {
	url       string
	uri       string
	currency  string
	useragent string
	token     string
	release   *models.Release
}

// NewReleaseService is a constructor
func NewReleaseService(url, uri, currency, useragent, token string) ReleaseRepository {
	return &ReleaseService{
		url:       url,
		uri:       uri,
		currency:  currency,
		useragent: useragent,
		token:     token,
	}
}

// FindReleaseByID is a method
func (rs *ReleaseService) FindReleaseByID(id int64) (*models.Release, error) {

	var err error

	params := url.Values{}
	params.Set("curr_abbr", rs.currency)

	url := fmt.Sprintf("%s%s%v", rs.url, rs.uri, id)

	// Add token to request headers
	headers := map[string]string{
		"User-Agent":    rs.useragent,
		"Authorization": "Discogs token=" + rs.token,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, params, headers)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &rs.release)
	return rs.release, err

}
