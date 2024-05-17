package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	accounts "github.com/cloudlink-omega/backend/pkg/accounts"
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager) // TODO: implement some sort of status check on the root endpoint (uptime, OS, load, memory use, etc.)
		w.Write([]byte("Hello, World!"))
	})

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
		res, err := dm.RegisterUser(&u)

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

		dm.SendHTMLEmail(&structs.EmailArgs{
			Subject:  "Welcome to CloudLink Omega!",
			To:       u.Email,
			Template: "hello",
		}, &structs.TemplateData{
			Name: u.Username,
		})

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
