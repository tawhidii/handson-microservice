package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhidii/user-service/db"
	"github.com/tawhidii/user-service/repository"
	"github.com/tawhidii/user-service/routes"
	"github.com/tawhidii/user-service/services"
	"log"
)

func main() {
	app := fiber.New()
	db := db.ConnectDatabase()
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)

	routes.Setup(app, authService)

	log.Fatal(app.Listen(":3000"))
}
