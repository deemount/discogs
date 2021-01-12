package models

// Member ...
type Member struct {
	Active      bool   `json:"active"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}
