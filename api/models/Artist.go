package models

// Artist ...
type Artist struct {
	ID             int      `gorm:"column:id;primary_key;unique" json:"id" example:"1"`
	Name           string   `gorm:"column:name;type:varchar(255)" json:"name"`
	Realname       string   `gorm:"column:realname;type:varchar(255)" json:"realname"`
	Members        []Member `json:"members,omitempty"`
	Aliases        []Alias  `json:"aliases,omitempty"`
	Namevariations []string `json:"namevariations"`
	Images         []Image  `json:"images"`
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	URLs           []string `json:"urls"`
	DataQuality    string   `json:"data_quality"`
}
