package main

import (
	publisher "gopubsub/pkg/publisher"
	subscriber "gopubsub/pkg/subscribers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		AppName:               "Go Pub Sub",
		DisableStartupMessage: false,
		EnablePrintRoutes:     true,
		BodyLimit:             5 * 1024 * 1024 * 1024,
		Concurrency:           256 * 1024,
	})

	ConnStrSubscriber := os.Getenv("CONN_STR_SUBSCRIBER")
	ConnStrPublisher := os.Getenv("CONN_STR_PUBLISHER")
	TopicName := os.Getenv("TOPIC_NAME")

	// Initialize Service Bus subscribers
	subscriber.InitializeServiceBus(ConnStrSubscriber)

	// Initialize the publisher and store it in Fiber's c.Locals
	publisherInstance := publisher.InitializeServiceBus(ConnStrPublisher, TopicName)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("publisher", publisherInstance)
		return c.Next()
	})

	SetupRoutes(app)

	// Gracefully shutdown Service Bus resources when the application exits.
	defer subscriber.ShutdownServiceBus()
	defer publisher.ShutdownServiceBus()

	log.Fatal(app.Listen(":3000"))
}
