package server

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var embedded_templates embed.FS

//go:embed assets/*
var embedded_static embed.FS

// TODO: add fields for the frontend server
type Server struct {
	ServerName string
	App        *fiber.App
}

// New creates a new Server instance.
//
// Server is a collection of API endpoints, pages, and other things that make up the
// backend service. It is meant to be used as the main entrypoint to the service.
//
// The created instance is returned.
func New(

	// Server Name is used for labeling the server. Format: [Country Code]-[Server Nickname]-[Designation].
	server_name string,

) *Server {
	srv := &Server{
		ServerName: server_name,
	}

	// Initialize template engine
	engine := html.NewFileSystem(http.FS(embedded_templates), ".html")

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

	// Initialize assets path
	srv.App.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(embedded_static),
		PathPrefix: "assets",
		Browse:     true,
	}))

	// Initialize middleware
	srv.App.Use(recover.New())
	srv.App.Use(logger.New())

	// Return created instance
	return srv
}
