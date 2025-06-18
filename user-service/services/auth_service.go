package services

import (
	"errors"

	models "github.com/tawhidii/user-service/models"
	"github.com/tawhidii/user-service/repository"
	"github.com/tawhidii/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email string, password string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (service *authService) Register(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return service.userRepo.CreateUser(user)
}

func (service *authService) Login(email string, password string) (string, error) {
	user, err := service.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	return utils.GenerateToken(user)
}

func (service *authService) ValidateToken(token string) (map[string]interface{}, error) {
	return utils.ValidateToken(token)
}
