package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhidii/user-service/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	users := api.Group("/users")

	users.Post("/", handlers.CreateUser)
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
}
