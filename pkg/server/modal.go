package server

import (
	"github.com/gofiber/fiber/v2"
)

// Handler to render the modal with dynamic content
func (s *Server) Modal(c *fiber.Ctx) error {

	// Fetch dynamic data (e.g., based on query params or IDs)
	id := c.Query("id")

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"Title":      "Dynamic Modal for Item " + id,
		"Content":    "This is dynamically loaded content for item " + id,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/layouts/modal", data)
}
