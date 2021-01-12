package models

// ArtistSource ...
type ArtistSource struct {
	Anv         string `json:"anv"`
	ID          int    `json:"id"`
	Join        string `json:"join"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
}
