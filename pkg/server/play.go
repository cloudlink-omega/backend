package server

import (
	"github.com/gofiber/fiber/v2"
)

// Handler for the explore page
func (s *Server) Play(c *fiber.Ctx) error {
	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Title":      "Play",
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/play", data, "views/layouts/default")
}
