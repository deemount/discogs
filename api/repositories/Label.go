package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// LabelRepository represents the contract between
type LabelRepository interface {
	FindLabelByID(id int64) (*models.Label, error)
}

// LabelService is a struct
type LabelService struct {
	url       string
	uri       string
	useragent string
	token     string
	label     *models.Label
}

// NewLabelService is a object
func NewLabelService(url, uri, useragent, token string) LabelRepository {
	return &LabelService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindLabelByID is a method
func (rs *LabelService) FindLabelByID(id int64) (*models.Label, error) {

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

	err = json.Unmarshal(body, &rs.label)
	return rs.label, err

}
