package createuser

import (
	"gopubsub/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUserAndPublishHandler(c *fiber.Ctx) error {
	var user model.User // make sure to import your "model" package
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// err := publishMessage(user)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// }

	return c.SendString("Message published successfully")
}
