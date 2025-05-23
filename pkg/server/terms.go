package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Terms(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)

	var username string
	if loggedIn {
		username = s.Authorization.GetNormalClaims(c).Username
	}

	// Create modal data based on the ID
	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Username":   username,
		"LoggedIn":   loggedIn,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/terms", data, "views/layouts/default")
}
