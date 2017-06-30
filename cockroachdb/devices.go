package cockroachdb

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mainflux/manager"
)

const sep string = ","

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

func toString(slice []uint) string {
	text := []string{}
	for _, v := range slice {
		text = append(text, fmt.Sprintf("%d", v))
	}

	return strings.Join(text, sep)
}
