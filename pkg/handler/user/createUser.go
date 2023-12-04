package user

import (
	"fmt"
	"gopubsub/pkg/model"
	"gopubsub/pkg/publisher"

	"github.com/gofiber/fiber/v2"
)

func CreateUserAndPublishHandler(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	publisherInstance, found := c.Locals("publisher").(*publisher.AzureServiceBusPublisher)
	if !found {
		return c.Status(fiber.StatusInternalServerError).SendString("Publisher instance not found in context")
	}

	operationTopic := "create"

	fmt.Println("Sending -----")
	fmt.Println(user)

	err := publisherInstance.PublishMessage(user, operationTopic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	fmt.Println("Success -----")

	return c.JSON(fiber.Map{
		"message": "Message published successfully",
		"content": user,
	})
}
