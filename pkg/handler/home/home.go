package home

import "github.com/gofiber/fiber/v2"

func HomePageHandler(c *fiber.Ctx) error {

	response := map[string]string{"success": "home"}

	return c.Status(fiber.StatusCreated).JSON(response)
}
