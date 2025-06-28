package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhidii/user-service/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/health", controllers.HealthCheck)
}
