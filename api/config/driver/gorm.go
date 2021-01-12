package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/deemount/discogs/api/config"
	"github.com/deemount/discogs/api/models"
)

// DataServiceRepository represents the contract
type DataServiceRepository interface {
	Connect() (*DataService, error)
}

// DataService is a struct
type DataService struct {
	Config *config.DB
	ORM    *gorm.DB
}

// NewDataService is a constructor
func NewDataService(config config.DB) DataServiceRepository {
	return &DataService{
		Config: &config,
	}
}

// Connect is a method
func (db *DataService) Connect() (*DataService, error) {

	var err error

	source := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s search_path=%s password=%s",
		db.Config.Host, db.Config.Port, db.Config.User, db.Config.Name, db.Config.SSL, db.Config.Schema, db.Config.PW)

	db.ORM, err = gorm.Open(db.Config.Driver, source)

	db.config()
	db.migrate()

	gorm.DefaultTableNameHandler = func(sql *gorm.DB, defaultTableName string) string {
		return db.Config.TblPrefix + "_" + defaultTableName
	}

	return db, err

}

func (db *DataService) config() {

	// Do not automatically convert to plural table names
	db.ORM.SingularTable(db.Config.SingularTable)
	db.ORM.LogMode(db.Config.LogMode)

}

func (db *DataService) migrate() {

	// Migrate missing fields
	// It does not change/ delete types or values
	db.ORM.Debug().AutoMigrate(&models.Artist{})

}
