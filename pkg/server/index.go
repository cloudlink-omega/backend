package server

import (
	"strconv"

	"github.com/cloudlink-omega/backend/pkg/flags"
	"github.com/gofiber/fiber/v2"
)

// Handler for the index page
func (s *Server) Index(c *fiber.Ctx) error {

	// Fetch dynamic data (e.g., based on query params or IDs)
	loggedIn, err := strconv.ParseBool(c.Params("auth"))
	if err != nil {
		loggedIn = false
	}

	data := map[string]interface{}{
		"ServerName": s.ServerName,
		"Title":      "Home Page",
		"LoggedIn":   loggedIn,
		"Cards": []map[string]interface{}{
			{
				"ImageURL":   "/assets/static/img/dummy2.png",
				"Title":      "CatChat Remastered",
				"Text":       "Meow around and make new friends! Remastered with voice chat and more.",
				"Creator":    "MikeDEV Games",
				"FooterText": "Launching April 1st, 2025",
				"Enabled":    true,
				"IsNew":      true,
				"ID":         "catchat",
				"Features": []map[string]interface{}{
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
				"Features": []map[string]interface{}{
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
				"Features": []map[string]interface{}{
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
		},
		"CardsSoon": []map[string]interface{}{
			{
				"ImageURL":   "",
				"Title":      "",
				"Text":       "",
				"Creator":    "",
				"FooterText": "",
				"Features":   []map[string]interface{}{},
			},
			{
				"ImageURL":   "",
				"Title":      "",
				"Text":       "",
				"Creator":    "",
				"FooterText": "",
				"Features":   []map[string]interface{}{},
			},
			{
				"ImageURL":   "",
				"Title":      "",
				"Text":       "",
				"Creator":    "",
				"FooterText": "",
				"Features":   []map[string]interface{}{},
			},
		},
	}

	// Render the modal template
	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("index", data, "layouts/default")
}
