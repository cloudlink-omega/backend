package server

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func StaticHandler(s *Server, group fiber.Router, hosted_path string) {
	group.Static("/projects_public", hosted_path+"/projects_public", fiber.Static{Compress: true, CacheDuration: time.Minute, Browse: true})
	group.Static("/developer_art", hosted_path+"/developer_art", fiber.Static{Compress: true, Browse: true})
	group.Static("/game_art", hosted_path+"/game_art", fiber.Static{Compress: true, Browse: true})
	group.Static("/thumbnail", hosted_path+"/thumbnails", fiber.Static{Compress: true, Browse: true})
	group.Static("/projects_private", hosted_path+"/projects_private", fiber.Static{
		Compress:      true,
		CacheDuration: time.Minute,
		Next: func(c *fiber.Ctx) bool {

			// TODO: Properly implement checks to make sure the requesting user has permissions to access a private project file
			path := c.Request().URI().Path()
			segments := strings.Split(string(path), "/")
			fileName := segments[len(segments)-1]
			log.Debug(fileName)
			return !s.Authorization.ValidFromNormal(c)
		},
	})
}
