package server

import (
	"github.com/cloudlink-omega/backend/pkg/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

type Server structs.Server

// New creates a new Server instance.
//
// Server is a collection of API endpoints, pages, and other things that make up the
// backend service. It is meant to be used as the main entrypoint to the service.
//
// The created instance is returned.
func New(

	// Point to where the template pages will be located.
	template_path string,

	// Server Name is used for labeling the server. Format: [Country Code]-[Server Nickname]-[Designation].
	server_name string,

) *Server {
	srv := &Server{
		ServerName: server_name,
	}

	// Initialize template engine
	engine := html.New(template_path, ".html")

	// Initialize app
	srv.App = fiber.New(fiber.Config{Views: engine, ErrorHandler: srv.ErrorPage})

	// Configure routes
	srv.App.Get("/admin", srv.Admin)
	srv.App.Get("/dashboard", srv.Dashboard)
	srv.App.Get("/developer", srv.DeveloperDashboard)
	srv.App.Get("/omegadash", srv.OmegaDash)
	srv.App.Get("/terms", srv.Terms)
	srv.App.Get("/modal", srv.Modal)
	srv.App.Get("/", srv.Index)

	// Initialize middleware
	srv.App.Use(recover.New())
	srv.App.Use(logger.New())

	// Return created instance
	return srv
}
