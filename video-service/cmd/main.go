package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tawhidii/user-service/routes"

	"log"
)

func main() {
	app := fiber.New()
	//db := db.ConnectDatabase()
	//userRepo := repository.NewUserRepository(db)
	//authService := services.NewAuthService(userRepo)

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
