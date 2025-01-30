package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) DeveloperDashboard(c *fiber.Ctx) error {
	loggedIn := s.Authorization.Valid(c)
	if !loggedIn {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Please login first before accessing the developer dashboard.",
		})
	}

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"LoggedIn":   true,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/developer", data, "views/layouts/nofooter")
}
