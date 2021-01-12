package config

// DB is a struct
type DB struct {
	Driver        string
	User          string
	PW            string
	Port          string
	Host          string
	SSL           string
	Schema        string
	TblPrefix     string
	Name          string
	SingularTable bool
	LogMode       bool
}
