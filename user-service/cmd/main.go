package main

import (
	"github.com/tawhidii/user-service/config"
	"github.com/tawhidii/user-service/db"
	"github.com/tawhidii/user-service/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	db.ConnectDatabase()

	app := fiber.New()
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
