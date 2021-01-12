package models

// Video ...
type Video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
}
