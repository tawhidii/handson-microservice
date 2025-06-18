package repository

import (
	"errors"

	models "github.com/tawhidii/user-service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) CreateUser(user *models.User) error {
	return repository.db.Create(user).Error
}

func (respository *userRepository) GetUserByEmail(email string) (*models.User, error) {

	var user models.User
	// GORM's .First() method populates the 'user' struct
	err := respository.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		// If the error is that the record was not found, return a clearer error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		// For any other error (like the scan error), return it as is
		return nil, err
	}

	// If there was no error, it means the user was found. Return it.
	return &user, nil
}
