package data

import (
	"database/sql"
	"log"
	"strings"

	"github.com/cloudlink-omega/backend/pkg/constants"
	errors "github.com/cloudlink-omega/backend/pkg/errors"
	structs "github.com/cloudlink-omega/backend/pkg/structs"
	"github.com/huandu/go-sqlbuilder"
	"github.com/oklog/ulid/v2"
)

func (mgr *Manager) RunSelectQuery(sb *sqlbuilder.SelectBuilder) (*sql.Rows, error) {
	query, args := sb.Build()
	if res, err := mgr.DB.Query(query, args...); err != nil {
		log.Printf("[DB] Failed to execute select request:\n\tquery: %s\n\targs: %v\n\tmessage: %s", query, args, err)
		return nil, err
	} else {
		return res, nil
	}
}

func (mgr *Manager) RunUpdateQuery(sb *sqlbuilder.UpdateBuilder) (sql.Result, error) {
	query, args := sb.Build()
	if res, err := mgr.DB.Exec(query, args...); err != nil {
		log.Printf("[DB] Failed to execute update request:\n\tquery: %s\n\targs: %v\n\tmessage: %s", query, args, err)
		return nil, err
	} else {
		return res, nil
	}
}

func (mgr *Manager) RunInsertQuery(sb *sqlbuilder.InsertBuilder) (sql.Result, error) {
	query, args := sb.Build()
	if res, err := mgr.DB.Exec(query, args...); err != nil {
		log.Printf("[DB] Failed to execute insert request:\n\tquery: %s\n\targs: %v\n\tmessage: %s", query, args, err)
		return nil, err
	} else {
		return res, nil
	}
}

func (mgr *Manager) RunDeleteQuery(sb *sqlbuilder.DeleteBuilder) (sql.Result, error) {
	query, args := sb.Build()
	if res, err := mgr.DB.Exec(query, args...); err != nil {
		log.Printf("[DB] Failed to execute delete request:\n\tquery: %s\n\targs: %v\n\tmessage: %s", query, args, err)
		return nil, err
	} else {
		return res, nil
	}
}

func (mgr *Manager) FindAllUsers() map[string]*structs.UserQuery {
	qy := sqlbuilder.NewSelectBuilder().
		Select("id", "username", "email", "created").
		From("users")

	if res, err := mgr.RunSelectQuery(qy); err != nil {
		res.Close()
		return nil
	} else {
		// Scan all rows, using ID as the key and User as the value
		rows := make(map[string]*structs.UserQuery)

		for res.Next() {
			var u structs.UserQuery
			err := res.Scan(&u.ID, &u.Username, &u.Email, &u.Created)
			if err != nil {
				log.Printf(`[DB] Failed to find all users: %s`, err)
				return nil
			}
			rows[u.ID] = &u
		}

		res.Close()
		return rows
	}
}

// RegisterUser registers a user in the Manager.
//
// u *structs.User - the user to be registered
// (sql.Result, error) - the result of the registration and any error encountered
func (mgr *Manager) RegisterUser(u *structs.Register) (bool, string, error) {

	var id = ulid.Make().String()

	qy := sqlbuilder.NewInsertBuilder().
		InsertInto("users").
		Cols("id", "username", "password", "email").
		Values(id, u.Username, u.Password, u.Email)

	res, err := mgr.RunInsertQuery(qy)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				return false, "", errors.ErrUsernameInUse
			} else if strings.Contains(err.Error(), "email") {
				return false, "", errors.ErrEmailInUse
			}
		}
		return false, "", err
	}
	rows, _ := res.RowsAffected()
	return rows == 1, id, nil
}

// GetUserPasswordHash retrieves the password hash for the given email.
//
// email string
// string, error
func (mgr *Manager) GetUserPasswordHash(email string) (string, error) {
	qy := sqlbuilder.NewSelectBuilder()
	qy.Distinct().Select("password")
	qy.From("users")
	qy.Where(
		qy.E("email", email),
	)
	var hash string
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		return "", err
	}
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&hash); err != nil {
			return "", err
		}
	} else {
		return "", errors.ErrUserNotFound
	}
	return hash, nil
}

// GetUserID retrieves the user ID for the given email.
//
// email string - the email of the user
// string, error - the user ID and any error encountered
func (mgr *Manager) GetUserID(email string) (string, error) {
	qy := sqlbuilder.NewSelectBuilder()
	qy.Distinct().Select("id")
	qy.From("users")
	qy.Where(
		qy.E("email", email),
	)
	var userid string
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		return "", err
	}
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&userid); err != nil {
			return "", err
		}
	} else {
		return "", errors.ErrUserNotFound
	}
	return userid, nil
}

// GenerateSessionToken generates a session token for the given user ID and origin.
//
// userid: string representing the user ID
// origin: string representing the origin of the session
// string: the generated session token
// error: an error, if any
func (mgr *Manager) GenerateSessionToken(userid string, origin string) (string, error) {
	usertoken := ulid.Make().String()

	// Bypass insert if in authless mode
	if mgr.AuthlessMode {
		mgr.AuthlessUserMap[usertoken] = userid
		return usertoken, nil
	}

	qy := sqlbuilder.NewInsertBuilder().
		InsertInto("sessions").
		Cols("id", "userid", "origin").
		Values(usertoken, userid, origin)
	res, err := mgr.RunInsertQuery(qy)
	if err != nil {
		return "", err
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", errors.ErrDatabaseError
	}
	return usertoken, nil
}

// GenerateMagicLink generates a magic link token for the given user ID and mode.
//
// userid: uint64 representing the user ID
// mode: uint8 representing the mode of the magic link
// state: uint8 representing the bitfield state of the magic link
// string: the generated magic link token
// error: an error, if any
func (mgr *Manager) GenerateMagicLink(userid string, mode uint8) (string, error) {
	token := ulid.Make().String()

	// Cannot work in authless mode
	if mgr.AuthlessMode {
		return "", errors.ErrAuthlessMode
	}

	qy := sqlbuilder.NewInsertBuilder().
		InsertInto("magic_links").
		Cols("id", "mode", "userid").
		Values(token, mode, userid)
	res, err := mgr.RunInsertQuery(qy)
	if err != nil {
		return "", err
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", errors.ErrDatabaseError
	}
	return token, nil
}

// DestroyMagicLink removes a magic link token from the database.
func (mgr *Manager) DestroyMagicLink(token string) error {
	// Cannot work in authless mode
	if mgr.AuthlessMode {
		return errors.ErrAuthlessMode
	}

	qy := sqlbuilder.NewDeleteBuilder()
	qy.DeleteFrom("magic_links").Where(qy.Equal("id", token))
	_, err := mgr.RunDeleteQuery(qy)
	if err != nil {
		return err
	}
	return nil
}

// newSaveSlotEntry saves a new slot entry to the database.
//
// Parameters:
//
//	slotnumber int - the slot number
//	slotdata string - the data to be saved in the slot
//	userid string - the user ID
//	ugi string - the game ID
//
// Return type:
//
//	error
func (mgr *Manager) newSaveSlotEntry(slotnumber uint8, slotdata string, userid string, ugi string) error {
	qy := sqlbuilder.NewInsertBuilder().
		InsertInto("saves").
		Cols("userid", "gameid", "slotid", "contents").
		Values(userid, ugi, slotnumber, slotdata)

	// Run the query
	res, err := mgr.RunInsertQuery(qy)
	if err != nil {
		return err
	}

	// Check if any rows were created
	if _, err := res.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (mgr *Manager) UpdateUserState(newstate uint, userid string) error {
	qy := sqlbuilder.NewUpdateBuilder()
	qy.Update("users").
		Set(
			qy.Assign("state", newstate),
		).
		Where(
			qy.E("id", userid),
		).
		Limit(1)

	// Run the query
	res, err := mgr.RunUpdateQuery(qy)
	if err != nil {
		return err
	}

	// Check if any errors occurred getting the number of rows affected
	if _, err := res.RowsAffected(); err != nil {
		return err
	}

	return nil
}

// updateSaveSlotEntry updates the save slot entry in the database.
//
// Parameters:
//
//	slotnumber int - the slot number to update.
//	slotdata string - the data to update in the slot.
//	userid string - the user ID associated with the save slot.
//	ugi string - the game ID associated with the save slot.
//
// Return:
//
//	error - returns an error if any operation fails.
func (mgr *Manager) updateSaveSlotEntry(slotnumber uint8, slotdata string, userid string, ugi string) error {
	qy := sqlbuilder.NewUpdateBuilder()
	qy.Update("saves").
		Set(
			qy.Assign("contents", slotdata),
		).
		Where(
			qy.E("userid", userid),
			qy.E("gameid", ugi),
			qy.E("slotid", slotnumber),
		).
		Limit(1)

	// Run the query
	res, err := mgr.RunUpdateQuery(qy)
	if err != nil {
		return err
	}

	// Check if any errors occurred getting the number of rows affected
	if _, err := res.RowsAffected(); err != nil {
		return err
	}

	return nil
}

// doesSaveSlotEntryExist checks if the slot entry exists in the Manager.
//
// slotnumber: the slot number to check
// userid: the user ID to check against
// ugi: the game ID to check against
// int: the existence status of the slot entry, error if any
func (mgr *Manager) doesSaveSlotEntryExist(slotnumber uint8, userid string, ugi string) (int, error) {

	// Check if the slot exists.
	var exists int
	qy := sqlbuilder.NewSelectBuilder()
	qy.Select("COUNT(*) > 0").
		From("saves").
		Where(
			qy.E("userid", userid),
			qy.E("gameid", ugi),
			qy.E("slotid", slotnumber),
		)

	// Run the query
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		return 0, err
	}

	// Get the result
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&exists); err != nil {
			return 0, err
		}
	}

	return exists, nil
}

func (mgr *Manager) loadSaveSlotEntry(slotnumber uint8, userid string, ugi string) (string, error) {

	var contents string
	qy := sqlbuilder.NewSelectBuilder()
	qy.Select("contents").
		From("saves").
		Where(
			qy.E("userid", userid),
			qy.E("gameid", ugi),
			qy.E("slotid", slotnumber),
		).
		Limit(1)

	// Run the query
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		return "", err
	}

	// Get the result
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&contents); err != nil {
			return "", err
		}
	}

	return contents, nil
}

// WriteSaveSlot writes or updates a save slot for a given user.
//
// Parameters:
//   - slotnumber: the slot number to write or update
//   - slotdata: the data to be saved in the slot
//   - userid: the user ID associated with the save slot
//   - ugi: the user group identifier
//
// Return type: error
func (mgr *Manager) WriteSaveSlot(slotnumber uint8, slotdata string, userid string, ugi string) error {

	// This function is not possible in authless mode
	if mgr.AuthlessMode {
		return errors.ErrAuthlessMode
	}

	// Check if the slot already exists
	exists, err := mgr.doesSaveSlotEntryExist(slotnumber, userid, ugi)
	if err != nil {
		return err
	}

	// If the slot doesn't exist, create it. Otherwise, update it.
	if exists == 0 {
		return mgr.newSaveSlotEntry(slotnumber, slotdata, userid, ugi)
	} else {
		return mgr.updateSaveSlotEntry(slotnumber, slotdata, userid, ugi)
	}
}

func (mgr *Manager) ReadSaveSlot(slotnumber uint8, userid string, ugi string) (string, error) {

	// This function is not possible in authless mode
	if mgr.AuthlessMode {
		return "", errors.ErrAuthlessMode
	}

	// Check if the slot exists
	exists, err := mgr.doesSaveSlotEntryExist(slotnumber, userid, ugi)
	if err != nil {
		return "", err
	}

	// If the slot doesn't exist, return an empty string
	if exists == 0 {
		return "", nil
	} else {
		// Retrieve slot data
		return mgr.loadSaveSlotEntry(slotnumber, userid, ugi)
	}
}

// VerifyUGI is a function that verifies the given UGI (Unique Game Identifier).
//
// It takes a parameter ugi string and returns an error.
func (mgr *Manager) VerifyUGI(ugi string) (string, string, error) {

	// Bypass UGI check if in authless mode
	if mgr.AuthlessMode {
		// TODO: somehow allow sysadmin to customize these parameters in authless mode
		return "", "", nil
	}

	qy := sqlbuilder.NewSelectBuilder()
	qy.Select(
		qy.As("g.name", "gameName"),
		qy.As("d.name", "developerName"),
	).
		From("games g", "developers d").
		Where(
			qy.E("g.id", ugi),
			qy.And("g.developerid = d.id"),
		)

	if res, err := mgr.RunSelectQuery(qy); err != nil {
		res.Close()
		return "", "", err
	} else {
		var gameName string
		var developerName string

		// Check if there's any output from the query (there should be 1 row if the game exists)
		if !res.Next() {
			return "", "", errors.ErrGameNotFound
		}
		// Scan the output into the variables
		if err := res.Scan(&gameName, &developerName); err != nil {
			return "", "", err
		}

		res.Close()
		return gameName, developerName, nil
	}
}

// VerifyMagicToken verifies a magic link token.
//
// It takes a linktoken string as a parameter and returns a client struct or an error.
func (mgr *Manager) VerifyMagicToken(linktoken string) (*structs.Client, uint8, error) {

	// This does not work in authless mode
	if mgr.AuthlessMode {
		return nil, constants.LINKMODE_UNDEFINED, errors.ErrAuthlessMode
	}

	qy := sqlbuilder.NewSelectBuilder()
	qy.Select(
		qy.As("u.username", "username"),
		qy.As("u.email", "email"),
		qy.As("u.id", "userid"),
		qy.As("u.state", "state"),
		qy.As("m.mode", "mode"),
	).
		From("magic_links m", "users u").
		Where(
			qy.E("m.id", linktoken),
			qy.And("u.id = m.userid"),
		)
	var linkmode uint8
	client := &structs.Client{}
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		res.Close()
		return nil, constants.LINKMODE_UNDEFINED, err
	}
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&client.Username, &client.Email, &client.ULID, &client.UserState, &linkmode); err != nil {
			return nil, constants.LINKMODE_UNDEFINED, err
		}
	} else {
		return nil, constants.LINKMODE_UNDEFINED, errors.ErrLinkNotFound
	}
	return client, linkmode, nil
}

// VerifySessionToken verifies the session token.
//
// It takes a usertoken string as a parameter and returns a client struct or an error.
func (mgr *Manager) VerifySessionToken(usertoken string) (*structs.Client, error) {

	// Bypass token check if in authless mode
	if mgr.AuthlessMode {
		client := &structs.Client{}

		// Check if the token is in the authless user map
		if _, ok := mgr.AuthlessUserMap[usertoken]; !ok {
			return nil, errors.ErrSessionNotFound
		}

		client.Username = mgr.AuthlessUserMap[usertoken]

		// Generate a random User ID
		client.ULID = ulid.Make().String()

		// Free memory by deleting the token from the map
		delete(mgr.AuthlessUserMap, usertoken)

		return client, nil
	}

	qy := sqlbuilder.NewSelectBuilder()
	qy.Select(
		qy.As("u.username", "username"),
		qy.As("u.email", "email"),
		qy.As("s.userid", "userid"),
		qy.As("s.origin", "origin"),
		qy.As("u.state", "userstate"),
		qy.As("s.state", "sessionstate"),
		qy.As("s.expires", "expires"),
	).
		From("sessions s", "users u").
		Where(
			qy.E("s.id", usertoken),
			qy.And("u.id = s.userid"),
		)
	client := &structs.Client{}
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		res.Close()
		return nil, err
	}
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&client.Username, &client.Email, &client.ULID, &client.Origin, &client.UserState, &client.SessionState, &client.Expiry); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.ErrSessionNotFound
	}
	return client, nil
}

// GetSessionInfoFromToken retrieves the session info from the session token.
//
// It takes a usertoken string as a parameter and returns a client struct or an error.
func (mgr *Manager) GetSessionInfoFromToken(usertoken string) (*structs.Session, error) {

	// Cannot work in authless mode
	if mgr.AuthlessMode {
		return nil, errors.ErrAuthlessMode
	}

	qy := sqlbuilder.NewSelectBuilder()
	qy.Select(
		qy.As("s.userid", "userid"),
		qy.As("s.state", "state"),
		qy.As("s.origin", "origin"),
		qy.As("s.created", "created"),
		qy.As("s.expires", "expires"),
	).
		From("sessions s").
		Where(
			qy.E("s.id", usertoken),
		)
	session := &structs.Session{}
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		res.Close()
		return nil, err
	}
	defer res.Close()
	if res.Next() {
		if err := res.Scan(&session.UserID, &session.State, &session.Origin, &session.Created, &session.Expiry); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.ErrSessionNotFound
	}
	return session, nil
}

// GetAllNewUsers retrieves a list of all new users in the database (for migration purposes).
//
// Returns a slice containing username and email pairs of all new users (users that haven't been given a welcome email for verification), or an error.
func (mgr *Manager) GetAllNewUsers() ([]*structs.BasicUserQuery, error) {

	// Cannot work in authless mode
	if mgr.AuthlessMode {
		return nil, errors.ErrAuthlessMode
	}

	qy := sqlbuilder.NewSelectBuilder()
	qy.Select(
		qy.As("u.username", "username"),
		qy.As("u.email", "email"),
	).
		From("users u").
		Where(
			qy.E("u.state", 0), // 0 = new
		)

	var users []*structs.BasicUserQuery
	res, err := mgr.RunSelectQuery(qy)
	if err != nil {
		res.Close()
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var u structs.BasicUserQuery
		if err := res.Scan(&u.Username, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
