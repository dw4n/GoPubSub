package main

import (
	publisher "gopubsub/pkg/publisher"
	subscriber "gopubsub/pkg/subscribers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	SetupRoutes(app)

	ConnStrSubscriber := ""
	ConnStrPublisher := ""
	TopicName := ""

	// Initialize Service Bus subscribers
	subscriber.InitializeServiceBus(ConnStrSubscriber)
	publisher.InitializeServiceBusPublisher(ConnStrPublisher, TopicName)

	// Gracefully shutdown Service Bus resources when the application exits.
	defer subscriber.ShutdownServiceBusSubscriber()
	defer publisher.ShutdownServiceBusPublisher()

	log.Fatal(app.Listen(":3000"))
}
