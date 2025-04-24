package server

import (
	"github.com/gofiber/fiber/v2"
)

// Handler for the explore page
func (s *Server) Play(c *fiber.Ctx) error {

	// Read ID path
	id := c.Params("id")
	if id == "" {
		return s.ErrorPage(c, &fiber.Error{Code: fiber.StatusNotFound, Message: "Whoops! This page doesn't exist."})
	}

	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Title":      "Play",
		"ID":         id,
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/play", data, "views/layouts/default")
}
