package v1

import (
	"fmt"
	"strings"

	"github.com/cloudlink-omega/accounts/pkg/constants"
	"github.com/cloudlink-omega/accounts/pkg/email"
	account_structs "github.com/cloudlink-omega/accounts/pkg/structs"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/oklog/ulid/v2"
)

type RegisterDevArgs struct {
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Members     []string `json:"members" form:"members"`
}

type RegisterGameArgs struct {
	DeveloperID string   `json:"developerid" form:"developerid"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Features    []string `json:"features" form:"features"`
}

/*
 * Registers a new developer account.
 * POST /api/v1/developer/register
 *
 * Body Types:
 * JSON, Multipart Form
 *
 * Content:
 *
 *	{
 *		"name": "My Developer Studio Name",
 *		"description": "This is my new developer studio!",
 *		"members": ["user ID"...]
 *	}
 *
 * Response: application/json
 * 200 OK
 *	{
 *		"result": "OK",
 *		"data": "(Developer ID)"
 *	}
 * or
 * 200 OK
 * {
 *		"result": "OK; warnings: [...]",
 *		"data": "(Developer ID)"
 *	}
 */
func (a *APIv1) RegisterDeveloper(c *fiber.Ctx) error {
	if !a.ParentServer.Authorization.ValidFromNormal(c) {
		return APIResult(c, fiber.StatusUnauthorized, "Not logged in!", nil)
	}

	claims := a.ParentServer.Authorization.GetNormalClaims(c)

	// Check if a server admin exists
	if !a.Database.AdminExists() {
		return APIResult(c, fiber.StatusTeapot, "Can't create developer; no server admin to notify!", nil)
	}

	// Read form
	var args RegisterDevArgs
	if err := c.BodyParser(&args); err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// Check if developer exists
	if a.Database.CheckForDeveloperNameConflict(args.Name) {
		return APIResult(c, fiber.StatusBadRequest, "Developer already exists!", nil)
	}

	// Set owner
	owner, err := a.ParentServer.Authorization.DB.GetUser(claims.ULID)
	if err != nil {
		return APIResult(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	owner_member := &types.DeveloperMember{UserID: owner.ID}
	owner_member.State.ManySet(
		constants.DEVMEMBER_IS_OWNER,
		constants.DEVMEMBER_IS_ACTIVE,
	)

	// Prepare members
	var warnings []string
	var members []*types.DeveloperMember
	members = append(members, owner_member)
	for _, member := range args.Members {
		user, err := a.ParentServer.Authorization.DB.GetUser(member)
		if err != nil {
			return APIResult(c, fiber.StatusInternalServerError, err.Error(), nil)
		}
		if user == nil {
			warnings = append(warnings, fmt.Sprintf("%s not found", member))
			continue
		}
		member := &types.DeveloperMember{UserID: user.ID}
		member.State.Set(constants.DEVMEMBER_IS_ACTIVE)
		members = append(members, member)
	}

	// Create the developer account
	developer_id := ulid.Make().String()
	developer := &types.Developer{
		ID:               developer_id,
		Name:             args.Name,
		Description:      args.Description,
		DeveloperMembers: members,
	}
	developer.State.Set(constants.DEVELOPER_IS_ACTIVE)
	a.Database.DB.Create(&developer)

	// Notify server admin
	server_admin := a.Database.GetAdmin()
	email.SendPlainEmail((*account_structs.MailConfig)(a.EmailConfig), &account_structs.EmailArgs{
		Subject:  "New developer registration",
		To:       server_admin.Email,
		Nickname: a.EmailConfig.Username,
	}, fmt.Sprintf(`Hello %s,
	
	A new developer account has been registered and is awaiting approval.

	Owner: %s (ID: %s, Email: %s) 
	ID: %s
	Name: %s
	Description: %s
	
	Regards,
	- %s`, server_admin.Username, owner.Username, owner.ID, owner.Email, developer_id, developer.Name, developer.Description, a.ParentServer.ServerName))

	if len(warnings) > 0 {
		res := "OK; Warnings: " + strings.Join(warnings, ", ")
		return APIResult(c, fiber.StatusOK, res, developer_id)
	}
	return APIResult(c, fiber.StatusOK, "OK", developer_id)
}

/*
 * Registers a new game.
 * POST /api/v1/developer/newgame
 *
 * Body Types:
 * JSON, Multipart Form
 *
 * Content:
 *
 *	{
 *		"developerid": "(Developer ID)",
 *		"name": "My Game Name",
 *		"description": "This is my new game!",
 *		"features": ["feature 1", "feature 2", ...]
 *	}
 *
 * Response: application/json
 * 200 OK
 *	{
 *		"result": "OK",
 *		"data": "(Game ID)"
 *	}
 * or
 * 200 OK
 * {
 *		"result": "OK; warnings: [...]",
 *		"data": "(Game ID)"
 *	}
 */
func (a *APIv1) RegisterGame(c *fiber.Ctx) error {
	if !a.ParentServer.Authorization.ValidFromNormal(c) {
		return APIResult(c, fiber.StatusUnauthorized, "Not logged in!", nil)
	}

	claims := a.ParentServer.Authorization.GetNormalClaims(c)

	// Check if a server admin exists
	if !a.Database.AdminExists() {
		return APIResult(c, fiber.StatusTeapot, "Can't create game; no server admin to notify!", nil)
	}

	// Read form
	var args RegisterGameArgs
	if err := c.BodyParser(&args); err != nil {
		return APIResult(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// Check if developer exists
	if !a.Database.DeveloperExists(args.DeveloperID) {
		return APIResult(c, fiber.StatusBadRequest, "Developer not found!", nil)
	}

	// Get developer
	developer := a.Database.GetDeveloper(args.DeveloperID)
	if developer == nil {
		return APIResult(c, fiber.StatusBadRequest, "Developer not found!", nil)
	}

	// Check if user is developer owner
	var found bool
	log.Debugf("Developer members: %s", developer.DeveloperMembers)
	for _, dev := range developer.DeveloperMembers {
		if dev.UserID == claims.ULID && dev.State.Read(constants.DEVMEMBER_IS_OWNER) {
			found = true
			break
		}
	}
	if !found {
		return APIResult(c, fiber.StatusUnauthorized, "Unauthorized.", nil)
	}

	// Prepare features
	var warnings []string
	var features []*types.FeatureTag
	for _, feature := range args.Features {
		// Verify the feature flag exists
		if a.Database.DB.Where("ID = ?", feature).First(&types.FeatureTag{}).RowsAffected == 0 {
			warnings = append(warnings, fmt.Sprintf("Feature flag %s not found", feature))
			continue
		}
		features = append(features, &types.FeatureTag{ID: feature})
	}

	game_id := ulid.Make().String()
	game := &types.DeveloperGame{
		ID:          game_id,
		DeveloperID: args.DeveloperID,
		Name:        args.Name,
		Description: args.Description,
		Features:    features,
	}

	// Create the game
	a.Database.DB.Create(&game)

	// Notify server admin
	server_admin := a.Database.GetAdmin()
	email.SendPlainEmail((*account_structs.MailConfig)(a.EmailConfig), &account_structs.EmailArgs{
		Subject:  "New developer game",
		To:       server_admin.Email,
		Nickname: a.EmailConfig.Username,
	}, fmt.Sprintf(`Hello %s,
	
	A developer has created a new game. It will not be visible until it is approved.

	Submitted by: %s (ID: %s, Email: %s) 
	Developer: %s (%s)
	ID: %s
	Name: %s
	Description: %s
	
	Regards,
	- %s`, server_admin.Username, claims.Username, claims.ULID, claims.Email, developer.Name, developer.ID, game_id, args.Name, args.Description, a.ParentServer.ServerName))

	if len(warnings) > 0 {
		res := "OK; Warnings: " + strings.Join(warnings, ", ")
		return APIResult(c, fiber.StatusOK, res, game_id)
	}
	return APIResult(c, fiber.StatusOK, "OK", game_id)
}
