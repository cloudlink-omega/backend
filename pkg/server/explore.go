package server

import (
	"github.com/gofiber/fiber/v2"
)

// Handler for the explore page
func (s *Server) Explore(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)

	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Title":      "Explore",
		"LoggedIn":   loggedIn,
		"PageCards": []map[string]any{
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy2.png",
				"Title":     "CatChat Remastered",
				"Developer": "MikeDEV Games",
			},
			{
				"Image":     "/assets/static/img/dummy3.png",
				"Title":     "Cloud Platformer Multiplayer Fun",
				"Developer": "griffpatch",
			},
			{
				"Image":     "/assets/static/img/dummy4.png",
				"Title":     "SB-3: The story of the project whose name was stolen",
				"Developer": "MikeDEV Games",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
			},
		},
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/explore", data, "views/layouts/default")
}
