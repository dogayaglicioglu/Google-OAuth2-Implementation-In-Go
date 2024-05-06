package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Middleware(c *fiber.Ctx) error {
	token := ExtractTokenFromContext(c)
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}

func ExtractTokenFromContext(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return ""
	}

	token := authHeader[len(prefix):]
	return token
}
