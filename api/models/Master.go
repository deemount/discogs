package models

// Master resource represents a set of similar releases.
// Masters (also known as `master releases`) have a `main release` which is often the chronologically earliest.
// More information https://www.discogs.com/developers#page:database,header:database-master-release
type Master struct {
	ID                   int            `json:"id"`
	Styles               []string       `json:"styles"`
	Genres               []string       `json:"genres"`
	Title                string         `json:"title"`
	Year                 int            `json:"year"`
	Tracklist            []Track        `json:"tracklist"`
	Notes                string         `json:"notes"`
	Artists              []ArtistSource `json:"artists"`
	Images               []Image        `json:"images"`
	Videos               []Video        `json:"videos"`
	NumForSale           int            `json:"num_for_sale"`
	LowestPrice          float64        `json:"lowest_price"`
	URI                  string         `json:"uri"`
	MainRelease          int            `json:"main_release"`
	MainReleaseURL       string         `json:"main_release_url"`
	MostRecentRelease    int            `json:"most_recent_release"`
	MostRecentReleaseURL string         `json:"most_recent_release_url"`
	VersionsURL          string         `json:"versions_url"`
	ResourceURL          string         `json:"resource_url"`
	DataQuality          string         `json:"data_quality"`
}
