package main

import (
	"os"

	"github.com/Abyss-fallers/ITB-go-back/database"
	"github.com/Abyss-fallers/ITB-go-back/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Post("/api/registration", handlers.Registration)
	app.Post("api/auth", handlers.Authentication)

	app.Listen(":" + os.Getenv("PORT"))
}
