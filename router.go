package main

import (
	home "gopubsub/pkg/handler/home"
	user "gopubsub/pkg/handler/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/", home.HomePageHandler)
	app.Post("/create", user.CreateUserAndPublishHandler)
}
