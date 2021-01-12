package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// ReleaseRatingRepository is a contract
type ReleaseRatingRepository interface {
	FindReleaseRatingByID(id int64) (*models.ReleaseRating, error)
}

// ReleaseRatingService is a struct
type ReleaseRatingService struct {
	url       string
	uri       string
	useragent string
	token     string
	rating    *models.ReleaseRating
}

// NewReleaseRatingService is a constructor
func NewReleaseRatingService(url, uri, useragent, token string) ReleaseRatingRepository {
	return &ReleaseRatingService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindReleaseRatingByID is a method
func (rs *ReleaseRatingService) FindReleaseRatingByID(id int64) (*models.ReleaseRating, error) {

	var err error

	url := fmt.Sprintf("%s%s%v%s", rs.url, rs.uri, id, "/rating")

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

	err = json.Unmarshal(body, &rs.rating)
	return rs.rating, err

}
