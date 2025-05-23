package v0

import (
	"github.com/gofiber/fiber/v2"
)

func (a *APIv0) Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Hello, world!")
}
