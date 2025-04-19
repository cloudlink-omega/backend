package v0

import (
	"github.com/cloudlink-omega/backend/pkg/structs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIv0 struct {
	App          *fiber.App
	ParentServer *structs.Server
	Database     *gorm.DB
}

type Result struct {
	Result string `json:"result"`
	Data   any    `json:"data"`
}

func New(s *structs.Server) *APIv0 {
	api := &APIv0{
		App:          fiber.New(),
		ParentServer: s,
		Database:     s.DB.DB,
	}

	// Cloud save slots feature
	api.App.Post("/save", api.Save)
	api.App.Post("/load", api.Load)

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
