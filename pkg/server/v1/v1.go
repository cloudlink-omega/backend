package v1

import (
	"time"

	account_structs "github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/backend/pkg/database"
	"github.com/cloudlink-omega/backend/pkg/structs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type APIv1 struct {
	App          *fiber.App
	ParentServer *structs.Server
	Database     *database.Database
	EmailConfig  *account_structs.MailConfig
}

type Result struct {
	Result string `json:"result"`
	Data   any    `json:"data,omitempty"`
}

func New(s *structs.Server) *APIv1 {
	api := &APIv1{
		App:          fiber.New(),
		ParentServer: s,
		Database:     s.DB,
		EmailConfig:  s.MailConfig,
	}

	// Cloud save slots feature
	api.App.Post("/save", api.Save)
	api.App.Post("/load", api.Load)
	// api.App.Post("/comment", api.Comment)

	// Developers API. Throttle to 5 requests per minute
	dev_group := fiber.New()

	dev_group.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: time.Minute,
	}))

	// Register a new developer account
	dev_group.Post("/register", api.RegisterDeveloper)

	// Register a new game
	dev_group.Post("/newgame", api.RegisterGame)

	// Mount the developer group
	api.App.Mount("/developer", dev_group)

	// Index
	api.App.Get("/", api.Index)

	return api
}

func APIResult(c *fiber.Ctx, status int, result string, data any) error {
	c.Set("Content-Type", "application/json")
	c.SendStatus(status)
	message, _ := json.Marshal(&Result{Result: result, Data: data})
	return c.SendString(string(message))
}
