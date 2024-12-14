package main

import (
	"os"
	"strings"

	"git.mikedev101.cc/MikeDEV/accounts"
	"git.mikedev101.cc/MikeDEV/backend/pkg/server"
	"git.mikedev101.cc/MikeDEV/signaling"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Initialize frontend server
	backend := server.New(
		"./templates",
		os.Getenv("SERVER_NAME"))

	// Initialize fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: backend.ErrorPage,
	})

	// Initialize the Signaling server
	signaling_server := signaling.New(
		strings.Split(os.Getenv("ALLOWED_DOMAINS"), " "),
		os.Getenv("TURN_ONLY") == "true",
	)

	// Initialize the Accounts service
	auth := accounts.New(
		os.Getenv("SERVER_URL"),
		os.Getenv("API_DOMAIN"),
		os.Getenv("API_URL"),
		os.Getenv("SERVER_NAME"),
		os.Getenv("PRIMARY_WEBSITE"),
		os.Getenv("SESSION_KEY"),
		os.Getenv("ENFORCE_HTTPS") == "true",
		"./templates/accounts",
	)

	// Initialize the OAuth providers
	auth.OAuth.Discord(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"))
	auth.OAuth.Google(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"))
	auth.OAuth.GitHub(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"))

	// Initialize middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Mount the routes
	app.Mount("/signaling", signaling_server.App)
	app.Mount("/accounts", auth.App)

	// Serve assets and hosted files
	app.Static("/assets", "./assets")
	app.Static("/hosted", "./hosted")

	// Mount the metrics endpoint
	app.Get("/metrics", monitor.New())

	// Mount the frontend server
	app.Mount("/", backend.App)

	// Run the app
	app.Listen(os.Getenv("API_URL"))
}
