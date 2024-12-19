package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Admin(c *fiber.Ctx) error {
	loggedIn := s.Authorization.Valid(c)

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"LoggedIn":   loggedIn,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/admin", data, "views/layouts/nofooter")
}
