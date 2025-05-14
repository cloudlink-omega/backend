package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/cloudlink-omega/storage/pkg/types"
)

// Handler for the explore page
func (s *Server) Play(c *fiber.Ctx) error {

	// Read ID path
	id := c.Params("id")

	// Check if the game exists
	var game types.DeveloperGame
	if s.DB.DB.Preload("Developer").Find(&game, "id = ?", id).Error == gorm.ErrRecordNotFound {
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
