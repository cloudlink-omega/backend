package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) Dashboard(c *fiber.Ctx) error {
	loggedIn := s.Authorization.ValidFromNormal(c)
	if !loggedIn {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Please login first before accessing the user dashboard.",
		})
	}

	claims := s.Authorization.GetNormalClaims(c)

	// Read all active sessions
	sessions, err := s.Accounts.DB.GetAllSessions(claims.ULID)
	if err != nil {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to retrieve user sessions.",
		})
	}

	log.Println(sessions)

	// Create modal data based on the ID
	data := map[string]any{
		"BaseURL":      s.ServerURL,
		"ServerName":   s.ServerName,
		"LoggedIn":     true,
		"GamesPlayed":  0,
		"FriendsMet":   0,
		"PointsEarned": 0,
		"Sessions":     sessions,
		"Logs": []map[string]any{
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
