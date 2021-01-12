package models

// ArtistReleases ...
type ArtistReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}
