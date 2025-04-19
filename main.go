package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/cloudlink-omega/accounts"
	"github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/backend/pkg/server"
	"github.com/cloudlink-omega/signaling"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func main() {
	// Initialize variables for flags and environment
	var err error
	var email_port int
	var https_mode, use_email, turn_only, enforce_https, enable_google, enable_github, enable_discord bool

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Read flags from environment
	https_mode = os.Getenv("HTTPS_MODE") == "true"
	use_email = os.Getenv("USE_EMAIL") == "true"
	turn_only = os.Getenv("TURN_ONLY") == "true"
	enforce_https = os.Getenv("ENFORCE_HTTPS") == "true"
	enable_google = os.Getenv("ENABLE_GOOGLE") == "true"
	enable_github = os.Getenv("ENABLE_GITHUB") == "true"
	enable_discord = os.Getenv("ENABLE_DISCORD") == "true"

	// Initialize database
	db, err := gorm.Open(sqlite.Open("mydb.db"), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// Compile authorized domains for CORS
	allowed_domains := strings.ReplaceAll(os.Getenv("ALLOWED_DOMAINS"), " ", ", ")

	// Read port from environment
	if use_email {
		email_port, err = strconv.Atoi(os.Getenv("EMAIL_PORT"))
		if err != nil {
			panic(err)
		}
	}

	// Initialize the Accounts server
	auth := accounts.New(
		"/accounts",
		os.Getenv("SERVER_URL"),
		os.Getenv("API_DOMAIN"),
		os.Getenv("API_URL"),
		os.Getenv("SERVER_NAME"),
		os.Getenv("PRIMARY_WEBSITE"),
		os.Getenv("SERVER_SECRET"),
		enforce_https,
		db,
		&structs.MailConfig{
			Enabled:  use_email,
			Port:     email_port,
			Server:   os.Getenv("EMAIL_SERVER"),
			Username: os.Getenv("EMAIL_USERNAME"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
	)

	// Initialize the Signaling server
	signaling_server := signaling.New(
		strings.Split(os.Getenv("ALLOWED_DOMAINS"), " "),
		turn_only,
		auth.APIv1.Auth,
		db,
	)

	// Initialize the Frontend server
	backend := server.New(
		os.Getenv("SERVER_NAME"),
		os.Getenv("SERVER_URL"),
		db,
		auth,
	)

	// Passthrough authorization server to the backend server
	backend.Authorization = auth.APIv1.Auth

	// Initialize the OAuth providers
	if enable_discord {
		auth.OAuth.Discord(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"))
	}
	if enable_google {
		auth.OAuth.Google(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"))
	}
	if enable_github {
		auth.OAuth.GitHub(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"))
	}

	// Initialize overall Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: backend.ErrorPage,
	})

	// Initialize Fiber middleware
	app.Use(fiber_logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowed_domains,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Mount servers in the Fiber app
	app.Mount("/signaling", signaling_server.App)
	app.Mount("/accounts", auth.App)
	app.Mount("/", backend.App)

	// Serve hosted files
	app.Static("/hosted", os.Getenv("HOSTED_PATH"), fiber.Static{Compress: true})

	// Mount metrics middleware
	app.Get("/metrics", monitor.New())

	// Run the app
	if https_mode {
		app.ListenTLS(os.Getenv("API_URL"), os.Getenv("HTTPS_CERT"), os.Getenv("HTTPS_KEY"))
	} else {
		app.Listen(os.Getenv("API_URL"))
	}
}
