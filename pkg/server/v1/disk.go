package v1

import (
	"github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SaveArgs struct {
	Slot uint8  `json:"save_slot" form:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	Data string `json:"save_data" form:"save_data" validate:"required,max=10000" label:"save_data"`
	UGI  string `json:"ugi" form:"ugi" validate:"ulid" label:"ugi"`
}

type LoadArgs struct {
	Slot uint8  `json:"save_slot" form:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	UGI  string `json:"ugi" form:"ugi" validate:"ulid" label:"ugi"`
}

func (a *APIv1) Save(c *fiber.Ctx) error {
	var args SaveArgs
	if err := c.BodyParser(&args); err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if args.Slot < 1 || args.Slot > 10 {
		return APIResult(c, fiber.StatusBadRequest, "Invalid save slot (must be a number between 1-10).", nil)
	}

	// Attempt to get session
	var claims *structs.Claims
	if a.ParentServer.Authorization.ValidFromNormal(c) {
		claims = a.ParentServer.Authorization.GetNormalClaims(c)
	} else {
		return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
	}

	// Get user from database
	user, _ := a.ParentServer.Accounts.DB.GetUser(claims.ULID)
	if user == nil {
		return APIResult(c, fiber.StatusInternalServerError, "Failed to get user for encrypting save.", nil)
	}

	// Encrypt save data
	encrypted, err := a.ParentServer.Accounts.DB.Encrypt(user, args.Data)
	if err != nil {
		return APIResult(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	// Check if developer game exists
	if a.Database.DB.Model(&types.DeveloperGame{}).First(&types.DeveloperGame{}, "id = ?", args.UGI).Error == gorm.ErrRecordNotFound {
		return APIResult(c, fiber.StatusBadRequest, "Invalid Game ID (UGI not found).", nil)
	}

	// Save
	var count int64
	a.Database.DB.Model(&types.UserGameSave{}).First(&types.UserGameSave{}, "user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI).Count(&count)
	if count > 0 {
		a.Database.DB.Model(&types.UserGameSave{}).Where("user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI).Update("save_data", encrypted)
	} else {
		a.Database.DB.Create(&types.UserGameSave{UserID: claims.ULID, SaveSlot: args.Slot, DeveloperGameID: args.UGI, SaveData: encrypted})
	}

	return APIResult(c, fiber.StatusOK, "OK", nil)
}

func (a *APIv1) Load(c *fiber.Ctx) error {

	var args LoadArgs
	if err := c.BodyParser(&args); err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if args.Slot < 1 || args.Slot > 10 {
		return APIResult(c, fiber.StatusBadRequest, "Invalid save slot (must be a number between 1-10).", nil)
	}

	// Attempt to get session
	var claims *structs.Claims
	if a.ParentServer.Authorization.ValidFromNormal(c) {
		claims = a.ParentServer.Authorization.GetNormalClaims(c)
	} else {
		return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
	}

	// Load
	var slot types.UserGameSave
	result := a.Database.DB.Model(&types.UserGameSave{}).First(&slot, "user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			return APIResult(c, fiber.StatusNotFound, "Save slot not found.", nil)
		default:
			return APIResult(c, fiber.StatusInternalServerError, result.Error.Error(), nil)
		}
	}

	// Get user from database
	user, _ := a.ParentServer.Accounts.DB.GetUser(claims.ULID)
	if user == nil {
		return APIResult(c, fiber.StatusInternalServerError, "Failed to get user for decrypting save.", nil)
	}

	// Decrypt save data
	decrypted, err := a.ParentServer.Accounts.DB.Decrypt(user, slot.SaveData)
	if err != nil {
		return APIResult(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return APIResult(c, fiber.StatusOK, "OK", decrypted)
}
