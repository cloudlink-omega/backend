package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (a *APIv1) Index(c *fiber.Ctx) error {
	return APIResult(c, fiber.StatusOK, "OK", "Hello, world!")
}
