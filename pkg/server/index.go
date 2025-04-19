package server

import (
	"github.com/cloudlink-omega/backend/pkg/flags"
	"github.com/gofiber/fiber/v2"
)

var BlankComingSoonEntry map[string]any = map[string]any{
	"ImageURL":   "",
	"Title":      "",
	"Text":       "",
	"Creator":    "",
	"FooterText": "",
	"Features":   []map[string]any{},
}

var BlankCardEntry map[string]any = map[string]any{
	"ImageURL":   "",
	"Title":      ". . .",
	"Text":       "",
	"Creator":    "",
	"FooterText": "",
	"Enabled":    false,
	"IsNew":      false,
	"ID":         "",
	"Features":   []map[string]string{},
}

// Handler for the index page
func (s *Server) Index(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)

	data := map[string]any{
		"BaseURL":    s.ServerURL,
		"ServerName": s.ServerName,
		"Title":      "Home Page",
		"LoggedIn":   loggedIn,
		"Cards": []map[string]any{
			{
				"ImageURL":   "/assets/static/img/dummy2.png",
				"Title":      "CatChat Remastered",
				"Text":       "Meow around and make new friends! Remastered with voice chat and more.",
				"Creator":    "MikeDEV Games",
				"FooterText": "Launching April 1st, 2025",
				"Enabled":    true,
				"IsNew":      true,
				"ID":         "catchat",
				"Features": []map[string]string{
					flags.SupportsAchievements,
					flags.SupportsControllers,
					flags.SuitableForAllAges,
					flags.SupportsLegacyProtocols,
					flags.GameIsForAllDevices,
					flags.GameIsOpenSource,
					flags.OriginalOnScratch,
					flags.MadeWithTurbowarp,
					flags.UnderReview,
					flags.SupportsSaveData,
					flags.SupportsProximityChat,
				},
			},
			{
				"ImageURL":   "/assets/static/img/dummy3.png",
				"Title":      "Cloud Platformer Multiplayer Fun",
				"Text":       "Enhanced with CloudLink Omega.",
				"Creator":    "griffpatch",
				"FooterText": "Launching January 1st, 2025",
				"Enabled":    true,
				"IsNew":      true,
				"ID":         "cpmf",
				"Features": []map[string]string{
					flags.SupportsControllers,
					flags.SuitableForAllAges,
					flags.SupportsLegacyProtocols,
					flags.GameIsForAllDevices,
					flags.MadeWithPenguinMod,
					flags.OriginalOnScratch,
					flags.GameIsOpenSource,
					flags.UnderReview,
				},
			},
			{
				"ImageURL":   "/assets/static/img/dummy4.png",
				"Title":      "SB-3: The story of the project whose name was stolen",
				"Text":       "A parody of BS-X on the Super Famicom / Satellaview.",
				"Creator":    "MikeDEV Games",
				"FooterText": "Launching April 1st, 2025",
				"Enabled":    true,
				"IsNew":      true,
				"ID":         "sb3",
				"Features": []map[string]string{
					flags.SupportsControllers,
					flags.HasDownloadableContent,
					flags.SuitableForAllAges,
					flags.SupportsLegacyProtocols,
					flags.GameIsForAllDevices,
					flags.MadeWithSheeptesterMod,
					flags.OriginalOnScratch,
					flags.GameIsOpenSource,
					flags.UnderReview,
					flags.SupportsSaveData,
				},
			},
			BlankCardEntry,
		},
		"CardsSoon": []map[string]any{
			BlankComingSoonEntry,
			BlankComingSoonEntry,
			BlankComingSoonEntry,
		},
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/index", data, "views/layouts/default")
}
