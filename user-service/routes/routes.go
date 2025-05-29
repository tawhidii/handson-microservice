package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhidii/user-service/controllers"
	"github.com/tawhidii/user-service/services"
)

func Setup(app *fiber.App, authService services.AuthService) {
	authController := controllers.NewAuthController(authService)
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
}
