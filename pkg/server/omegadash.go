package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) OmegaDash(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)
	if !loggedIn {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Please login first.",
		})
	}

	client := s.Authorization.GetNormalClaims(c)

	// Create modal data based on the ID
	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Username":   client.Username,
		"Points":     1234,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/omega", data, "views/layouts/omegadash")
}
