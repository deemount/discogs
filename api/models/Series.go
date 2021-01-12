package models

// Series ...
type Series struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
}
