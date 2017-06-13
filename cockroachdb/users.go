package cockroachdb

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/mainflux/manager"
)

const (
	uniqueErr string = "unique_violation"
)

type userRepository struct {
	db *gorm.DB
}

type userRecord struct {
	gorm.Model
	Email    string `gorm:"type:varchar(254);not null;unique"`
	Password string `gorm:"type:char(60);not null"`
}

// TableName sets user entity table name.
func (u userRecord) TableName() string {
	return "users"
}

// NewUserRepository instantiates CockroachDB-specific user repository.
func NewUserRepository(db *gorm.DB) manager.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Save(user manager.User) error {
	r := &userRecord{
		Email:    user.Email,
		Password: user.Password,
	}

	err := ur.db.Create(r).Error
	if pqErr, ok := err.(*pq.Error); ok && strings.Contains(pqErr.Code.Name(), uniqueErr) {
		return manager.ErrConflict
	}

	return err
}

func (ur *userRepository) Exists(user manager.User) bool {
	u := &userRecord{}

	if ne := ur.db.Where("email = ?", user.Email).First(u).RecordNotFound(); ne {
		return false
	}

	return u.Password == user.Password
}
