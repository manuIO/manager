package cockroachdb

import "github.com/jinzhu/gorm"

const (
	uniqueErr string = "unique_violation"
	fkErr     string = "foreign key violation"
)

// Connect establishes connection to the CockroachDB cluster.
func Connect(addr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&userRecord{}, &deviceRecord{})
	db.LogMode(false)

	return db, nil
}
