package controllers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/tawhidii/user-service/model"
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
