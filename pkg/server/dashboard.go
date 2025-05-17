package server

import (
	"fmt"

	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/mileusna/useragent"

	"github.com/gofiber/fiber/v2"
)

type ParsedSession struct {
	*types.UserSession
	Device string
}

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

	// Parse sessions list
	var parsed_sessions []*ParsedSession
	for _, session := range sessions {
		ua := useragent.Parse(session.UserAgent)

		dev := fmt.Sprintf("%s v%s on %s", ua.Name, ua.Version, ua.OS)

		if ua.Desktop {
			dev += " (Desktop)"
		} else if ua.Mobile {
			dev += " (Mobile)"
		} else if ua.Tablet {
			dev += " (Tablet)"
		} else if ua.Bot {
			dev += " (Bot)"
		} else {
			dev += " (Other/Unknown)"
		}

		parsed_session := &ParsedSession{
			UserSession: session,
			Device:      dev,
		}
		parsed_sessions = append(parsed_sessions, parsed_session)
	}

	// Get the top 20 logs
	logs, pages, err := s.Accounts.DB.GetUserLogs(claims.ULID, 0)
	if err != nil {
		return s.ErrorPage(c, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to retrieve user logs.",
		})
	}

	// Create modal data based on the ID
	data := map[string]any{
		"BaseURL":      s.ServerURL,
		"ServerName":   s.ServerName,
		"LoggedIn":     true,
		"GamesPlayed":  "WIP",
		"FriendsMet":   "WIP",
		"PointsEarned": "WIP",
		"Sessions":     parsed_sessions,
		"Empty":        len(logs) == 0,
		"Logs":         logs,
		"Page":         1,
		"Pages":        pages,
	}

	c.Context().SetContentType("text/html; charset=utf-8")
	return c.Render("views/user", data, "views/layouts/nofooter")
}
