package server

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
)

type explore_card_entry struct {
	Cover     *types.Image
	Image     string
	Title     string
	Developer string
	ID        string
}

// Handler for the explore page
func (s *Server) Explore(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)
	claims := s.Authorization.GetNormalClaims(c)

	var loaded_cards []*explore_card_entry
	games, _, _ := s.DB.GetAllGames(0, 20)

	for _, game := range games {
		entry := &explore_card_entry{
			Title:     game.Name,
			Developer: game.Developer.Name,
			ID:        game.ID,
			Cover:     game.Thumbnail,
		}

		loaded_cards = append(loaded_cards, entry)
	}

	var username string
	if loggedIn {
		username = claims.Username
	}

	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Title":      "Explore",
		"LoggedIn":   loggedIn,
		"Username":   username,
		"PageCards":  loaded_cards, /* []map[string]any{
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy2.png",
				"Title":     "CatChat Remastered",
				"Developer": "MikeDEV Games",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy3.png",
				"Title":     "Cloud Platformer Multiplayer Fun",
				"Developer": "griffpatch",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy4.png",
				"Title":     "SB-3: The story of the project whose name was stolen",
				"Developer": "MikeDEV Games",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
			{
				"Image":     "/assets/static/img/dummy1.png",
				"Title":     "Lorem Ipsum",
				"Developer": "Sample Text",
				"ID":        "01HNPHRWS0N0AYMM5K4HN31V4W",
			},
		},*/
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/explore", data, "views/layouts/default")
}
