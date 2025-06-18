package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	models "github.com/tawhidii/user-service/models"
	"github.com/tawhidii/user-service/services"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (authController *AuthController) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := authController.AuthService.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"user": user})
}

func (authController *AuthController) Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	token, err := authController.AuthService.Login(data.Email, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": token})
}

func (authController *AuthController) ValidateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
	}
	// The header should be in the format "Bearer <token>"
	// We split the header string to get the token part.
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header format"})
	}
	tokenString := parts[1]

	// Call the service to validate the token
	claims, err := authController.AuthService.ValidateToken(tokenString)
	if err != nil {
		// If the token is invalid, return an unauthorized error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token", "details": err.Error()})
	}

	// If the token is valid, return the claims
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"valid": true, "claims": claims})
}
