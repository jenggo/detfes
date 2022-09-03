package routes

import "github.com/gofiber/fiber/v2"

func Return(c *fiber.Ctx, status int, message string, faces int, error bool) error {
	return c.Status(status).JSON(fiber.Map{
		"detected": faces,
		"error":    error,
		"message":  message,
	})
}
