package models

// MasterVersions retrieves a list of all releases that are versions of this master.
type MasterVersions struct {
	Pagination Page      `json:"pagination"`
	Versions   []Version `json:"versions"`
}
