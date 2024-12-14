package getprefix

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) string {
	var result string

	// If we are inside of a router, we need to fetch the name of what router we are in
	if c.Route().Name != "" {
		result = c.Route().Name
	} else {
		result = strings.Split(c.Path(), "/")[1]
	}

	log.Print("Here's the router name: " + result)

	// If we are outside of a router, we need to fetch the name of the route we are on
	return result
}
