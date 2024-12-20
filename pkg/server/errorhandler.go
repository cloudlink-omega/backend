package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) ErrorHandler(c *fiber.Ctx, err error) error {

	var error_code int
	if e, ok := err.(*fiber.Error); ok {
		error_code = e.Code
	} else {
		error_code = fiber.StatusInternalServerError
	}

	c.SendStatus(error_code)

	request_content_type := string(c.Request().Header.Peek("Content-Type"))
	var match bool
	for _, t := range []string{"html", "plain", "form"} {
		match = !match && strings.Contains(request_content_type, t)
	}

	if match {
		return c.SendString(err.Error())
	}

	if render_error := c.Render("views/error", &map[string]string{
		"Message":    err.Error(),
		"Status":     fmt.Sprint(error_code),
		"ServerName": s.ServerName}); render_error != nil {
		panic(render_error)
	}

	return nil
}
