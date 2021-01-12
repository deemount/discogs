package config

// API ...
type API struct {
	Version     int
	Host        string
	Path        string
	Port        string
	StrictSlash bool
	LimitClient int
	Destination string
	Msg         Message
}

// Message ...
type Message struct {
	Status int
	Notice string
	Out    string
}
