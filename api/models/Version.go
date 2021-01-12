package models

// Version ...
type Version struct {
	Catno       string `json:"catno"`
	Country     string `json:"country"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Released    string `json:"released"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
}
