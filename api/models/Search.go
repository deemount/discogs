package models

// Search describes search response
type Search struct {
	Pagination Page     `json:"pagination"`
	Results    []Result `json:"results,omitempty"`
}
