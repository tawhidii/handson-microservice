package utils

import (
	"github.com/golang-jwt/jwt/v5"
	models "github.com/tawhidii/user-service/model"
	"time"
)

var JwtSecret = []byte("your_secret_key")

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
