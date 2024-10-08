package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	accounts "github.com/cloudlink-omega/backend/pkg/accounts"
	"github.com/cloudlink-omega/backend/pkg/bitfield"
	constants "github.com/cloudlink-omega/backend/pkg/constants"
	dm "github.com/cloudlink-omega/backend/pkg/data"
	errors "github.com/cloudlink-omega/backend/pkg/errors"
	structs "github.com/cloudlink-omega/backend/pkg/structs"
	utils "github.com/cloudlink-omega/backend/pkg/utils"
)

func RootRouter(r chi.Router) {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	// Register custom label function for validator
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	// Display server nickname and version
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)
		w.Write([]byte(fmt.Sprintf("%s | v%s\n", dm.ServerNickname, constants.Version)))
	})

	// Verify magic link
	r.Get("/verify", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// Read query parameters from URL
		queryParams := r.URL.Query()
		var token = queryParams.Get("token")

		var user *structs.Client
		var mode uint8
		var err error
		if user, mode, err = dm.VerifyMagicToken(token); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Verify mode
		if mode != constants.LINKMODE_EMAIL {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("This token is not a valid email verification or unsubscribe token."))
			return
		}

		// Update user's account state
		user.UserState.Set(constants.USER_IS_ACTIVE)
		if err := dm.UpdateUserState(uint(user.UserState), user.ULID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Delete magic link
		if err := dm.DestroyMagicLink(token); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Write response to client
		w.Write([]byte(fmt.Sprintf("Hello %s, your email has been verified successfully.", user.Username)))
	})

	// Unsubscribe magic link
	r.Get("/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// Read query parameters from URL
		queryParams := r.URL.Query()
		var token = queryParams.Get("token")

		var user *structs.Client
		var mode uint8
		var err error
		if user, mode, err = dm.VerifyMagicToken(token); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Verify mode
		if mode != constants.LINKMODE_EMAIL {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("This token is not a valid email verification or unsubscribe token."))
			return
		}

		// Update user's account state
		user.UserState.Set(constants.USER_IS_ACTIVE)
		user.UserState.Set(constants.USER_IS_EMAIL_DISABLED)
		if err := dm.UpdateUserState(uint(user.UserState), user.ULID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Delete magic link
		if err := dm.DestroyMagicLink(token); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Write response to client
		w.Write([]byte(fmt.Sprintf("The email address %s for user %s has been unsubscribed successfully. If this was in error, contact support.", user.Email, user.Username)))
	})

	// Login
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// If authless mode is enabled, return a randomly generated auth token
		if dm.AuthlessMode {

			// Load request body as JSON into login struct
			var u structs.Login
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			// Generate random session token
			var usertoken string
			if res, err := dm.GenerateSessionToken(u.Email, r.URL.Hostname()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				usertoken = res
			}

			// Write response to client with the session token
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(usertoken))
			return
		}

		// Load request body as JSON into login struct
		var u structs.Login
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate login struct
		if handleValidationError(w, validate.Struct(u)) {
			return
		}

		// Define vars
		var hash string
		var userid string
		var usertoken string

		// Grab user's password hash
		if res, err := dm.GetUserPasswordHash(u.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		} else {
			hash = res
		}

		// Verify the hash matches the provided password
		if err := accounts.VerifyPassword(u.Password, hash); err != nil {
			// Incorrect password
			if strings.Contains(err.Error(), "does not match") {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Incorrect password"))
				return
			}
			// Something else went wrong
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong while verifying your login credentials. Please try again."))
			return
		}

		// Grab user ID from email
		if res, err := dm.GetUserID(u.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		} else {
			userid = res
		}

		// Generate session token
		if res, err := dm.GenerateSessionToken(userid, r.URL.Hostname()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		} else {
			usertoken = res
		}

		// Write response to client with the session token
		w.Write([]byte(usertoken))
	})

	// Save to slot
	r.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// If authless mode is enabled, disable this endpoint
		if dm.AuthlessMode {
			w.WriteHeader(http.StatusGone)
			w.Write([]byte("Authless mode is enabled on this server. Save slots are not available."))
			return
		}

		// Load request body as JSON into save struct
		var s structs.Save
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate save struct
		if handleValidationError(w, validate.Struct(s)) {
			return
		}

		// Validate session token & UGI format
		if errmsg := utils.VariableContainsValidationError("token", validate.Var(s.Token, "ulid")); errmsg != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed session token."))
			return
		}
		if errmsg := utils.VariableContainsValidationError("ugi", validate.Var(s.UGI, "ulid")); errmsg != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed UGI."))
			return
		}

		// Validate UGI exists
		if _, _, err := dm.VerifyUGI(s.UGI); err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		// Find and read user account given session token
		var session *structs.Session
		session, err := dm.GetSessionInfoFromToken(s.Token)

		// Handle errors
		if err != nil {
			switch err {
			case errors.ErrSessionNotFound:
				w.WriteHeader(http.StatusUnauthorized)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}

		// Check if session is expired
		if session.Expiry <= time.Now().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Session token has expired."))
			return
		}

		// Write save slot
		if err = dm.WriteSaveSlot(s.SaveSlot, fmt.Sprint(s.SaveData), session.UserID, s.UGI); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("OK"))
	})

	// Load from slot
	r.Post("/load", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// If authless mode is enabled, disable this endpoint
		if dm.AuthlessMode {
			w.WriteHeader(http.StatusGone)
			w.Write([]byte("Authless mode is enabled on this server. Save slots are not available."))
			return
		}

		// Load request body as JSON into load struct
		var s structs.Load
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate save struct
		if handleValidationError(w, validate.Struct(s)) {
			return
		}

		// Validate session token & UGI format
		if errmsg := utils.VariableContainsValidationError("token", validate.Var(s.Token, "ulid")); errmsg != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed session token."))
			return
		}
		if errmsg := utils.VariableContainsValidationError("ugi", validate.Var(s.UGI, "ulid")); errmsg != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed UGI."))
			return
		}

		// Validate UGI exists
		if _, _, err := dm.VerifyUGI(s.UGI); err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		// Find and read user account given session token
		var session *structs.Session
		session, err := dm.GetSessionInfoFromToken(s.Token)

		// Handle errors
		if err != nil {
			switch err {
			case errors.ErrSessionNotFound:
				w.WriteHeader(http.StatusUnauthorized)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}

		// Check if session is expired
		if session.Expiry <= time.Now().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Session token has expired."))
			return
		}

		// Read save slot
		data, err := dm.ReadSaveSlot(s.SaveSlot, session.UserID, s.UGI)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(data))
	})

	// Register an account
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		// If authless mode is enabled, disable this endpoint
		if dm.AuthlessMode {
			w.WriteHeader(http.StatusGone)
			w.Write([]byte("Authless mode is enabled on this server. User registration is not available."))
			return
		}

		// Load request body as JSON into register struct
		var u structs.Register
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate register struct
		if handleValidationError(w, validate.Struct(u)) {
			return
		}

		// Hash password
		u.Password = accounts.HashPassword(u.Password)

		// Register user
		res, id, err := dm.RegisterUser(&u)

		// Handle errors
		if err != nil {
			switch err {
			case errors.ErrUsernameInUse:
				w.WriteHeader(http.StatusConflict)
			case errors.ErrEmailInUse:
				w.WriteHeader(http.StatusConflict)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}

		if !res {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong while registering your account. Please try again later."))
			return
		}

		fmt.Printf("Registered user %s\n", u.Username)

		var verifLink string
		if verifLink, err = dm.GenerateMagicLink(id, constants.LINKMODE_EMAIL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Send welcome email - TODO: Somehow make this a template.
		if err := dm.SendPlainEmail(&structs.EmailArgs{
			Subject: "Please verify your email address",
			To:      u.Email,
		}, fmt.Sprintf(
`Hello %s. You are receiving this email because you created an account with CloudLink Omega.
		
You can verify your email address by going to this link: %s
		
If you are not the intended recipient (or you've changed your mind), please visit this link: %s
		
If you have questions, comments, or concerns, please reach out to MikeDEV via Discord: https://discord.gg/BZ7TWeMF75
		
Please support the project by marking this email as "not spam". Beware phishing: We will never ask for your login credentials over email.`,
			u.Username,
			fmt.Sprintf("%s/api/v0/verify?token=%s", dm.PublicHostname, verifLink),
			fmt.Sprintf("%s/api/v0/unsubscribe?token=%s", dm.PublicHostname, verifLink),
		)); err != nil {
			log.Printf("Error sending welcome email: %s", err)
		} else {

			// Update user's account state
			var state bitfield.Bitfield8

			state.Set(constants.USER_IS_EMAIL_REGISTERED)
			if err := dm.UpdateUserState(uint(state), id); err != nil {
				log.Printf("Failed to update user state: %s", err)
			}
		}

		// Scan output
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	})
}

func handleValidationError(w http.ResponseWriter, err error) bool {
	if err != nil && len(err.(validator.ValidationErrors)) > 0 {
		// Create error message
		msg := "Validation failed:\n"
		for _, err := range err.(validator.ValidationErrors) {
			msg += fmt.Sprintf("%s: %s\n", err.Field(), err.Tag())
		}

		// Write error message to response
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return true
	}
	return false
}
