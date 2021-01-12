package models

// LabelReleases is a list of Releases associated with the label.
type LabelReleases struct {
	Pagination Page            `json:"pagination"`
	Releases   []ReleaseSource `json:"releases"`
}
