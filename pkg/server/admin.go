package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Admin(c *fiber.Ctx) error {

	// Fetch dynamic data (e.g., based on query params or IDs)
	loggedIn := c.QueryBool("auth")

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"LoggedIn":   loggedIn,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("admin", data, "layouts/default")
}
