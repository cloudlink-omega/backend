package structs

import "github.com/gofiber/fiber/v2"

type Server struct {
	// TODO: add fields for the frontend server
	ServerName string
	App        *fiber.App
}
