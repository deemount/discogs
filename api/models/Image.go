package models

// Image ...
type Image struct {
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
}
