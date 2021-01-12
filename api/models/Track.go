package models

// Track ...
type Track struct {
	Duration     string         `json:"duration"`
	Position     string         `json:"position"`
	Title        string         `json:"title"`
	Type         string         `json:"type_"`
	Extraartists []ArtistSource `json:"extraartists,omitempty"`
	Artists      []ArtistSource `json:"artists,omitempty"`
}
