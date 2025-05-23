package v1

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type CommentPost struct {
	Comment string `json:"comment" form:"comment"`
	Parent  string `json:"parent" form:"parent"`
	GameID  string `json:"gameid" form:"gameid"`
}

func (a *APIv1) Comment(c *fiber.Ctx) error {
	if !a.ParentServer.Authorization.ValidFromNormal(c) {
		return APIResult(c, fiber.StatusUnauthorized, "Not logged in!", nil)
	}

	// Read form
	var comment_post CommentPost
	if err := c.BodyParser(&comment_post); err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if !a.Database.GameExists(comment_post.GameID) {
		return APIResult(c, fiber.StatusBadRequest, "Invalid game ID!", nil)
	}

	if comment_post.Parent != "" && !a.Database.CommentExists(comment_post.Parent) {
		return APIResult(c, fiber.StatusBadRequest, "Invalid parent ID!", nil)
	}

	comment_id := ulid.Make().String()
	comment := &types.GameComment{
		ID:              comment_id,
		DeveloperGameID: comment_post.GameID,
		Content:         comment_post.Comment,
	}

	if comment_post.Parent != "" {
		comment.ParentID = &comment_post.Parent
	}

	a.Database.DB.Create(&comment)

	return APIResult(c, fiber.StatusOK, "OK", comment_id)
}
