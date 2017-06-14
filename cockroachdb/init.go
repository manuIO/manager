package cockroachdb

import (
	"github.com/jinzhu/gorm"
)

// Connect establishes connection to the CockroachDB cluster.
func Connect(addr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&userRecord{})
	db.LogMode(false)

	return db, nil
}
