package main

import (
	"pingPong/internal/db"
	"pingPong/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database := db.Connect()

	app := fiber.New()

	userHandler := handlers.NewUserHandler(database)

	app.Post("/users", userHandler.CreateUser)

	app.Listen(":3000")
}
