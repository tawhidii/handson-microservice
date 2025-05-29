package repository

import (
	models "github.com/tawhidii/user-service/model"
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
	err := respository.db.Where("email = ?", email).First(&user).Error
	return nil, err
}
