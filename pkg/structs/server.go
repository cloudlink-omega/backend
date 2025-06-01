package structs

import (
	"github.com/cloudlink-omega/accounts"
	"github.com/cloudlink-omega/accounts/pkg/authorization"
	"github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/backend/pkg/database"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

type Server struct {
	ServerName    string
	ServerURL     string
	DB            *database.Database
	App           *fiber.App
	Authorization *authorization.Auth
	Accounts      *accounts.Accounts
	Cache         *types.DBCache
	Policy        *bluemonday.Policy
	HostedPath    string
	MailConfig    *structs.MailConfig
}
