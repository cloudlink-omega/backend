package v0

import (
	"github.com/cloudlink-omega/backend/pkg/database"
	"github.com/cloudlink-omega/backend/pkg/structs"
	"github.com/gofiber/fiber/v2"
)

type APIv0 struct {
	App          *fiber.App
	ParentServer *structs.Server
	Database     *database.Database
}

type Result struct {
	Result string `json:"result"`
	Data   any    `json:"data,omitempty"`
}

func New(s *structs.Server) *APIv0 {
	api := &APIv0{
		App:          fiber.New(),
		ParentServer: s,
		Database:     s.DB,
	}

	// Cloud save slots feature
	api.App.Post("/save", api.Save)
	api.App.Post("/load", api.Load)

	// Index
	api.App.Get("/", api.Index)

	return api
}
