package models

// Label resource represents a label, company, recording studio, location,
// or other entity involved with artists and releases.
type Label struct {
	Profile     string     `json:"profile"`
	ReleasesURL string     `json:"releases_url"`
	Name        string     `json:"name"`
	ContactInfo string     `json:"contact_info"`
	URI         string     `json:"uri"`
	Sublabels   []Sublable `json:"sublabels"`
	URLs        []string   `json:"urls"`
	Images      []Image    `json:"images"`
	ResourceURL string     `json:"resource_url"`
	ID          int        `json:"id"`
	DataQuality string     `json:"data_quality"`
}
