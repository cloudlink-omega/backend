package server

import (
	"github.com/gofiber/fiber/v2"
)

// Handler for the explore page
func (s *Server) Play(c *fiber.Ctx) error {

	// Read ID path
	id := c.Params("id")

	// Check if the game exists
	game := s.DB.GetGame(id)

	if game == nil {
		return s.ErrorPage(c, &fiber.Error{Code: fiber.StatusNotFound, Message: "Whoops! Game not found."})
	}

	data := map[string]any{
		"BaseURL":         s.ServerURL,
		"ServerName":      s.ServerName,
		"Title":           game.Name,
		"GameName":        game.Name,
		"DeveloperName":   game.Developer.Name,
		"GameDescription": game.Description,
		"ID":              game.ID,
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/play", data, "views/layouts/default")
}
