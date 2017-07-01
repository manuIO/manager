package cockroachdb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect establishes connection to the CockroachDB cluster.
func Connect(addr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&userRecord{})
	db.AutoMigrate(&deviceRecord{})
	db.LogMode(false)

	return db, nil
}
