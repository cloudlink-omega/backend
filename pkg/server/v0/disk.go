package v0

import (
	"log"
	"strconv"

	authorization_structs "github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SaveArgs struct {
	Slot  uint8  `json:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	Data  string `json:"save_data" validate:"required,max=10000" label:"save_data"`
	UGI   string `json:"ugi" validate:"ulid" label:"ugi"`
	Token string `json:"token" validate:"" label:"token"`
}

type LoadArgs struct {
	Slot  uint8  `json:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	UGI   string `json:"ugi" validate:"ulid" label:"ugi"`
	Token string `json:"token" validate:"" label:"token"`
}

func (a *APIv0) Save(c *fiber.Ctx) error {

	// Attempt to get claims based on token or cookie
	var claims *authorization_structs.Claims
	var save SaveArgs
	if c.Body() != nil {
		if err := json.Unmarshal(c.Body(), &save); err != nil {
			return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
		}
	}

	// Require authorization
	if save.Token != "" {
		if !a.ParentServer.Authorization.ValidFromToken(c, save.Token) {
			return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
		}
		claims = a.ParentServer.Authorization.GetClaimsFromToken(c, save.Token)
	} else {
		if !a.ParentServer.Authorization.ValidFromNormal(c) {
			return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
		}
		claims = a.ParentServer.Authorization.GetNormalClaims(c)
	}

	// Try to read the contents, accept JSON or form data
	if c.Body() == nil {
		slot, err := strconv.ParseUint(c.FormValue("save_slot", "1"), 10, 8)
		if err != nil {
			return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
		}
		save.Slot = uint8(slot)
		save.Data = c.FormValue("save_data", "")
		save.UGI = c.FormValue("ugi", "")
	}

	usersave := &types.UserGameSave{
		UserID:          claims.ULID,
		SaveSlot:        save.Slot,
		SaveData:        save.Data,
		DeveloperGameID: save.UGI,
	}

	// Check if the UGI exists
	var count int64
	if err := a.Database.
		Model(&types.DeveloperGame{}).
		Where("id = ?", save.UGI).
		Count(&count).
		Error; err != nil {

		return APIResult(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	if count == 0 {
		return APIResult(c, fiber.StatusInternalServerError, "Invalid UGI!", nil)
	}

	// Create or update save
	result := a.Database.
		Model(&types.UserGameSave{}).
		Where("user_id = ?", claims.ULID).
		Where("save_slot = ?", save.Slot).
		Where("developer_game_id = ?", save.UGI).
		Save(&usersave)

	if result.Error != nil {
		return APIResult(c, fiber.StatusInternalServerError, result.Error.Error(), nil)
	}
	if result.RowsAffected > 0 {
		return APIResult(c, fiber.StatusOK, "OK", nil)
	}
	if result.RowsAffected > 1 {
		return APIResult(c, fiber.StatusTeapot, "Database anomaly detected; contact an administrator!", nil)
	}

	return APIResult(c, fiber.StatusOK, "OK", nil)
}

func (a *APIv0) Load(c *fiber.Ctx) error {

	// Attempt to get claims based on token or cookie
	var claims *authorization_structs.Claims
	var load LoadArgs
	if c.Body() != nil {
		if err := json.Unmarshal(c.Body(), &load); err != nil {
			return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
		}
	}

	// Require authorization
	if load.Token != "" {
		log.Println("Using legacy token")
		if !a.ParentServer.Authorization.ValidFromToken(c, load.Token) {
			return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
		}
		claims = a.ParentServer.Authorization.GetClaimsFromToken(c, load.Token)
	} else {
		log.Println("Using cookie-based token")
		if !a.ParentServer.Authorization.ValidFromNormal(c) {
			return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
		}
		claims = a.ParentServer.Authorization.GetNormalClaims(c)
	}

	// Try to read the contents, accept JSON or form data
	slot, err := strconv.ParseUint(c.FormValue("save_slot", "1"), 10, 8)
	if err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	load.Slot = uint8(slot)
	load.UGI = c.FormValue("ugi", "")

	// Get save
	var usersave types.UserGameSave
	result := a.Database.
		Model(&types.UserGameSave{}).
		Where("user_id = ?", claims.ULID).
		Where("save_slot = ?", load.Slot).
		Where("developer_game_id = ?", load.UGI).
		First(&usersave)

	switch result.Error {
	case nil:
		return APIResult(c, fiber.StatusOK, "OK", usersave.SaveData)

	case gorm.ErrRecordNotFound:
		return APIResult(c, fiber.StatusInternalServerError, "Slot not found.", nil)

	default:
		return APIResult(c, fiber.StatusInternalServerError, result.Error.Error(), nil)
	}
}
