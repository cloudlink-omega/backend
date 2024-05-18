package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	constants "github.com/cloudlink-omega/backend/pkg/constants"
	dm "github.com/cloudlink-omega/backend/pkg/data"
	errors "github.com/cloudlink-omega/backend/pkg/errors"
	structs "github.com/cloudlink-omega/backend/pkg/structs"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func VerifyAdminSession(validate *validator.Validate, dm *dm.Manager, w http.ResponseWriter, r *http.Request) (bool, *structs.Client) {

	// Load request body as JSON into token struct
	var s structs.AdminToken
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return false, nil
	}

	// Validate save struct
	if handleValidationError(w, validate.Struct(s)) {
		return false, nil
	}

	// Find and read user account given session token
	var session *structs.Client
	session, err := dm.VerifySessionToken(s.Token)

	// Handle errors
	if err != nil {
		switch err {
		case errors.ErrSessionNotFound:
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return false, nil
	}

	// Check if session is expired
	if session.Expiry <= time.Now().Unix() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Session token has expired."))
		return false, nil
	}

	// Check if user is admin
	if !session.UserState.Read(constants.USER_IS_ADMIN) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("You don't have permission to perform this action."))
		return false, nil
	}

	return true, session
}

func AdminRouter(r chi.Router) {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	// Register custom label function for validator
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	r.Post("/resend_hello", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		var admin *structs.Client
		var ok bool
		if ok, admin = VerifyAdminSession(validate, dm, w, r); !ok {
			return
		}

		var verifLink string
		var err error
		if verifLink, err = dm.GenerateMagicLink(admin.ULID, constants.LINKMODE_EMAIL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := dm.SendHTMLEmail(&structs.EmailArgs{
			Subject:  "Hello, " + admin.Username + "!",
			To:       admin.Email,
			Template: "hello",
		}, &structs.TemplateData{
			Name:             admin.Username,
			VerificationLink: fmt.Sprintf("%s/api/v0/verify?token=%s", dm.PublicHostname, verifLink),
			UnsubscribeLink:  fmt.Sprintf("%s/api/v0/unsubscribe?token=%s", dm.PublicHostname, verifLink),
		}); err != nil {
			log.Printf("[Admin] Error sending email: %s", err)
		} else {
			// Update user's account state
			admin.UserState.Set(constants.USER_IS_EMAIL_REGISTERED)
			if err := dm.UpdateUserState(uint(admin.UserState), admin.ULID); err != nil {
				log.Printf("[Admin] Error updating user state: %s", err)
			}
		}

		w.Write([]byte("OK"))
	})

	r.Post("/test_email", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		var admin *structs.Client
		var ok bool
		if ok, admin = VerifyAdminSession(validate, dm, w, r); !ok {
			return
		}

		if err := dm.SendPlainEmail(&structs.EmailArgs{
			Subject: "Admin plaintext email test",
			To:      admin.Email,
		}, fmt.Sprintf("Hello %s. This is a test.", admin.Username)); err != nil {
			log.Printf("[Admin] Error sending email: %s", err)
		}

		w.Write([]byte("OK"))
	})

	r.Post("/all_new_users", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		if ok, _ := VerifyAdminSession(validate, dm, w, r); !ok {
			return
		}

		var users []*structs.BasicUserQuery
		users, err := dm.GetAllNewUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Dump the users slice as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	r.Post("/test_html_email", func(w http.ResponseWriter, r *http.Request) {
		dm := r.Context().Value(constants.DataMgrCtx).(*dm.Manager)

		var admin *structs.Client
		var ok bool
		if ok, admin = VerifyAdminSession(validate, dm, w, r); !ok {
			return
		}

		if err := dm.SendHTMLEmail(&structs.EmailArgs{
			Subject:  "Admin HTML email test",
			To:       admin.Email,
			Template: "admin_test",
		}, &structs.TemplateData{
			Name: admin.Username,
		}); err != nil {
			log.Printf("[Admin] Error sending email: %s", err)
		}

		w.Write([]byte("OK"))
	})
}
