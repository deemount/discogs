package models

// Identifier ...
type Identifier struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}
