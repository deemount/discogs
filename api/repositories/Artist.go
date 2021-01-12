package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deemount/discogs/api/models"
	"github.com/deemount/discogs/api/requests"
)

// ArtistRepository represents the contract between
type ArtistRepository interface {
	FindArtistByID(id int64) (*models.Artist, error)
}

// ArtistService is a struct
type ArtistService struct {
	url       string
	uri       string
	useragent string
	token     string
	artist    *models.Artist
}

// NewArtistService is a object
func NewArtistService(url, uri, useragent, token string) ArtistRepository {
	return &ArtistService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		token:     token,
	}
}

// FindArtistByID is a method
func (rs *ArtistService) FindArtistByID(id int64) (*models.Artist, error) {

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

	err = json.Unmarshal(body, &rs.artist)
	return rs.artist, err

}
