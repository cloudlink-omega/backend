package v0

import (
	"github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SaveArgs struct {
	Slot  uint8  `json:"save_slot" form:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	Data  string `json:"save_data" form:"save_data" validate:"required,max=10000" label:"save_data"`
	UGI   string `json:"ugi" form:"ugi" validate:"ulid" label:"ugi"`
	Token string `json:"token" form:"token" validate:"" label:"token"`
}

type LoadArgs struct {
	Slot  uint8  `json:"save_slot" form:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	UGI   string `json:"ugi" form:"ugi" validate:"ulid" label:"ugi"`
	Token string `json:"token" form:"token" validate:"" label:"token"`
}

func (a *APIv0) Save(c *fiber.Ctx) error {
	var args SaveArgs
	if err := c.BodyParser(&args); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if args.Slot < 1 || args.Slot > 10 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid save slot (must be a number between 1-10).")
	}

	// Attempt to get session
	var claims *structs.Claims
	if a.ParentServer.Authorization.ValidFromToken(args.Token) {
		claims = a.ParentServer.Authorization.GetClaimsFromToken(args.Token)
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized.")
	}

	// Get user from database
	user, err := a.ParentServer.Accounts.DB.GetUser(claims.ULID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not find user.")
	}

	// Encrypt save data
	encrypted, err := a.ParentServer.Accounts.DB.Encrypt(user, args.Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Check if developer game exists
	if a.Database.DB.Model(&types.DeveloperGame{}).First(&types.DeveloperGame{}, "id = ?", args.UGI).Error == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Game ID (UGI not found).")
	}

	// Save
	var count int64
	a.Database.DB.Model(&types.UserGameSave{}).First(&types.UserGameSave{}, "user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI).Count(&count)
	if count > 0 {
		a.Database.DB.Model(&types.UserGameSave{}).Where("user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI).Update("save_data", encrypted)
	} else {
		a.Database.DB.Create(&types.UserGameSave{UserID: claims.ULID, SaveSlot: args.Slot, DeveloperGameID: args.UGI, SaveData: encrypted})
	}

	return c.Status(fiber.StatusOK).SendString("OK")
}

func (a *APIv0) Load(c *fiber.Ctx) error {

	var args LoadArgs
	if err := c.BodyParser(&args); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if args.Slot < 1 || args.Slot > 10 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid save slot (must be a number between 1-10).")
	}

	// Attempt to get session
	var claims *structs.Claims
	if a.ParentServer.Authorization.ValidFromToken(args.Token) {
		claims = a.ParentServer.Authorization.GetClaimsFromToken(args.Token)
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized.")
	}

	// Load
	var slot types.UserGameSave
	result := a.Database.DB.Model(&types.UserGameSave{}).First(&slot, "user_id = ? AND save_slot = ? AND developer_game_id = ?", claims.ULID, args.Slot, args.UGI)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).SendString("Save slot not found.")
		default:
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}
	}

	// Get user from database
	user, err := a.ParentServer.Accounts.DB.GetUser(claims.ULID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not find user.")
	}

	// Decrypt save data
	decrypted, err := a.ParentServer.Accounts.DB.Decrypt(user, slot.SaveData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString(decrypted)
}
