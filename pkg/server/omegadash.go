package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) OmegaDash(c *fiber.Ctx) error {
	client := s.Authorization.GetClaims(c)

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"Username":   client.Username,
		"Points":     1234,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/omega", data, "views/layouts/omegadash")
}
