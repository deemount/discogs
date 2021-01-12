package models

// ReleaseRating serves response for community release rating request.
type ReleaseRating struct {
	ID     int    `json:"release_id"`
	Rating Rating `json:"rating"`
}
