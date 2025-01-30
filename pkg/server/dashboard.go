package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Dashboard(c *fiber.Ctx) error {
	loggedIn := s.Authorization.Valid(c)
	if !loggedIn {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Please login first before accessing the user dashboard.",
		})
	}

	// Create modal data based on the ID
	data := map[string]interface{}{
		"ServerName":   s.ServerName,
		"LoggedIn":     true,
		"GamesPlayed":  0,
		"FriendsMet":   0,
		"PointsEarned": 0,
		"Logs": []map[string]interface{}{
			{
				"Timestamp": "nil",
				"Action":    "Authentication (IP: 127.0.0.1)",
				"Success":   true,
				"Message":   "Passed",
			},
			{
				"Timestamp": "nil",
				"Action":    "Authentication (IP: 127.0.0.1)",
				"Warn":      true,
				"Message":   "Multifactor Step Failed",
			},
			{
				"Timestamp": "nil",
				"Action":    "Authentication (IP: 192.168.1.1)",
				"Fail":      true,
				"Message":   "Invalid Credentials",
			},
			{
				"Timestamp": "nil",
				"Action":    "Setup TOTP",
				"Success":   true,
				"Message":   "Ok",
			},
			{
				"Timestamp": "nil",
				"Action":    "Verify Email",
				"Success":   true,
				"Message":   "Ok",
			},
			{
				"Timestamp": "nil",
				"Action":    "Create Account",
				"Success":   true,
				"Message":   "Account Created",
			},
		},
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/user", data, "views/layouts/nofooter")
}
