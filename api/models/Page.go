package models

// Page ...
type Page struct {
	PerPage int      `json:"per_page"`
	Items   int      `json:"items"`
	Page    int      `json:"page"`
	URLs    URLsList `json:"urls"`
	Pages   int      `json:"pages"`
}
