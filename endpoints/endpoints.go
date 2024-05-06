package endpoints

import "github.com/gofiber/fiber/v2"

func ProtectedEndpoint(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Reaching protected route is sucessfull!",
	})
}
