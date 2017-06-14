package cockroachdb

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/mainflux/manager"
)

const uniqueErr string = "unique_violation"

var _ manager.UserRepository = (*userRepository)(nil)

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
	rec := &userRecord{
		Email:    user.Email,
		Password: user.Password,
	}

	err := ur.db.Create(rec).Error
	if pqErr, ok := err.(*pq.Error); ok && strings.Contains(pqErr.Code.Name(), uniqueErr) {
		return manager.ErrConflict
	}

	return err
}

func (ur *userRepository) Get(email string) (manager.User, error) {
	rec := &userRecord{}

	if ne := ur.db.Where("email = ?", email).First(rec).RecordNotFound(); ne {
		return manager.User{}, manager.ErrInvalidCredentials
	}

	user := manager.User{
		Email:    rec.Email,
		Password: rec.Password,
	}

	return user, nil
}