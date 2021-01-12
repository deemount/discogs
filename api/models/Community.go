package models

// Community ...
type Community struct {
	Contributors []Contributor `json:"contributors"`
	DataQuality  string        `json:"data_quality"`
	Have         int           `json:"have"`
	Rating       Rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    Submitter     `json:"submitter"`
	Want         int           `json:"want"`
}
