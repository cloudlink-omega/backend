package v0

import (
	"github.com/gofiber/fiber/v2"
)

func (a *APIv0) Index(c *fiber.Ctx) error {
	return APIResult(c, fiber.StatusOK, "OK", "Hello, world!")
}
