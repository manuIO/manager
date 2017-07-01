package cockroachdb

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/mainflux/manager"
)

var _ manager.DeviceRepository = (*deviceRepository)(nil)

type deviceRepository struct {
	db *gorm.DB
}

type deviceRecord struct {
	gorm.Model
	OwnerID     string `sql:"type:varchar(254) REFERENCES users(email)"`
	Name        string `gorm:"type:varchar(50);not null"`
	Key         string `gorm:"type:text;not null"`
	Description string
	Channels    string
}

// TableName sets device entity table name.
func (dr deviceRecord) TableName() string {
	return "devices"
}

// NewDeviceRepository instantiates CockroachDB-specific device repository.
func NewDeviceRepository(db *gorm.DB) manager.DeviceRepository {
	return &deviceRepository{db}
}

func (dr *deviceRepository) Save(device manager.Device) (uint, error) {
	rec := &deviceRecord{
		OwnerID:     device.Owner,
		Name:        device.Name,
		Key:         device.Key,
		Description: device.Description,
		Channels:    toString(device.Channels),
	}

	err := dr.db.Create(rec).Error

	if pqErr, ok := err.(*pq.Error); ok && strings.Contains(pqErr.Message, fkErr) {
		return 0, manager.ErrUnauthorizedAccess
	}

	return rec.ID, err
}

func (dr *deviceRepository) One(id uint, owner string) (manager.Device, error) {
	rec := &deviceRecord{}

	if dr.db.Where("id = ? AND owner_id = ?", id, owner).First(rec).RecordNotFound() {
		return manager.Device{}, manager.ErrNotFound
	}

	device := manager.Device{
		ID:          rec.ID,
		Owner:       rec.OwnerID,
		Name:        rec.Name,
		Key:         rec.Key,
		Description: rec.Description,
		Channels:    fromString(rec.Channels),
	}

	return device, nil
}
