package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestApiKey := c.Get("API-Key")
		if requestApiKey != apiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API Key",
			})
		}
		return c.Next()
	}
}
