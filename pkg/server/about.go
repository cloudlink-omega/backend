package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) About(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)

	// Create modal data based on the ID
	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"LoggedIn":   loggedIn,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/about", data, "views/layouts/default")
}
