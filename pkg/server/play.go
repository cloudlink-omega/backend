package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Handler for the explore page
func (s *Server) Play(c *fiber.Ctx) error {

	// Read ID path
	id := c.Params("id")

	// Check if the game exists
	game := s.DB.GetGame(id)

	_, file_error := os.Stat(s.HostedPath + "/projects_public/" + id)

	claims := s.Authorization.GetNormalClaims(c)
	loggedIn := s.Authorization.ValidFromNormal(c)
	var username string
	if loggedIn {
		username = claims.Username
	}

	if file_error != nil || game == nil {
		if file_error != nil {
			log.Error(file_error)
		}
		data := map[string]any{
			"BaseURL":    s.ServerURL,
			"LoggedIn":   loggedIn,
			"Username":   username,
			"ServerName": s.ServerName,
			"Title":      "Whoops!",
		}
		c.Context().SetContentType("text/html; charset=utf-8")
		c.Status(fiber.StatusNotFound)
		return c.Render("views/play_not_found", data, "views/layouts/default")
	}

	data := map[string]any{
		"BaseURL":         s.ServerURL,
		"LoggedIn":        loggedIn,
		"Username":        username,
		"ServerName":      s.ServerName,
		"Title":           game.Name,
		"GameName":        game.Name,
		"DeveloperName":   game.Developer.Name,
		"GameDescription": game.Description,
		"ID":              game.ID,
		"Features":        game.Features,
		"Comments": []map[string]string{
			{
				"ID":       "1",
				"Username": "MikeDEV",
				"Comment":  `Hello world! This is an example comment. **Very cool!** *This should render as markdown.*`,
				"Date":     "1/1/2023",
			},
		},
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/play", data, "views/layouts/default")
}
