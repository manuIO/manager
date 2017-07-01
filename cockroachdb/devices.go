package cockroachdb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mainflux/manager"
)

var _ manager.DeviceRepository = (*deviceRepository)(nil)

type deviceRepository struct {
	db *gorm.DB
}

type deviceRecord struct {
	gorm.Model
	Owner       string `gorm:"type:varchar(254);not null"`
	Name        string `gorm:"type:varchar(50);not null"`
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
		Owner:       device.Owner,
		Name:        device.Name,
		Description: device.Description,
		Channels:    toString(device.Channels),
	}

	err := dr.db.Create(rec).Error
	return rec.ID, err
}

func (dr *deviceRepository) One(id uint, owner string) (manager.Device, error) {
	rec := &deviceRecord{}

	if dr.db.Where("id = ? AND owner = ?", id, owner).First(rec).RecordNotFound() {
		return manager.Device{}, manager.ErrNotFound
	}

	device := manager.Device{
		ID:          rec.ID,
		Owner:       rec.Owner,
		Name:        rec.Name,
		Description: rec.Description,
		Channels:    fromString(rec.Channels),
	}

	return device, nil
}
