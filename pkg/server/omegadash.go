package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) OmegaDash(c *fiber.Ctx) error {
	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"Username":   "MikeDEV",
		"Points":     1234,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("omega", data, "layouts/omegadash")
}
