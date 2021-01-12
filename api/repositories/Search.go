package repositories

import (
	"github.com/deemount/discogs/api/models"
)

// SearchServiceRepository is a contract
type SearchServiceRepository interface {
	// Search makes search request to discogs.
	// Issue a search query to database. This endpoint accepts pagination parameters.
	// Authentication (as any user) is required.
	// https://www.discogs.com/developers/#page:database,header:database-search
	Search(req models.SearchRequest) (*models.Search, error)
}

// SearchService is a struct
type SearchService struct {
	url    string
	search *models.Search
}

// NewSearchService is a constructor
func NewSearchService(url string) SearchServiceRepository {
	return &SearchService{
		url: url,
	}
}

// Search is a method
func (rs *SearchService) Search(req models.SearchRequest) (*models.Search, error) {

	var err error

	return rs.search, err

}
