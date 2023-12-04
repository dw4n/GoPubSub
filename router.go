package main

import (
	"log"

	createuser "gopubsub/pkg/handler/CreateUser"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Post("/create", createuser.CreateUserAndPublishHandler)
}
