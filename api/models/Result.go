package models

// Result describes a part of search result
type Result struct {
	Style       []string  `json:"style,omitempty"`
	Thumb       string    `json:"thumb,omitempty"`
	CoverImage  string    `json:"cover_image,omitempty"`
	Title       string    `json:"title,omitempty"`
	Country     string    `json:"country,omitempty"`
	Format      []string  `json:"format,omitempty"`
	URI         string    `json:"uri,omitempty"`
	Community   Community `json:"community,omitempty"`
	Label       []string  `json:"label,omitempty"`
	Catno       string    `json:"catno,omitempty"`
	Year        string    `json:"year,omitempty"`
	Genre       []string  `json:"genre,omitempty"`
	ResourceURL string    `json:"resource_url,omitempty"`
	Type        string    `json:"type,omitempty"`
	ID          int       `json:"id,omitempty"`
	MasterID    int       `json:"master_id,omitempty"`
}
