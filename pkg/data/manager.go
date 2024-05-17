package data

import (
	"context"
	"database/sql"
	"log"

	"github.com/cloudlink-omega/backend/pkg/structs"
)

type Manager struct {
	Ctx                  context.Context
	ServerNickname       string
	EnableEmail          bool
	MailConfig           structs.MailConfig
	DB                   *sql.DB
	AuthlessMode         bool
	UseInMemoryClientMgr bool
	AuthlessUserMap      map[string]string // ULID session token -> username. Used for authless mode.
}

func New(serverNickname string, sqlDriver string, sqlUrl string, authlessMode bool, useInMemoryClientMgr bool, enableEmail bool, emailPort int, emailServer string, emailUsername string, emailPassword string) *Manager {

	// Create background context
	ctx := context.Background()

	if authlessMode {
		log.Println("[Data Manager] Bypassing DB connection due to authless mode.")

		// Return manager
		return &Manager{
			ServerNickname:       serverNickname,
			DB:                   nil,
			Ctx:                  ctx,
			AuthlessMode:         true,
			UseInMemoryClientMgr: useInMemoryClientMgr,
			AuthlessUserMap:      make(map[string]string),
			EnableEmail:          enableEmail,
			MailConfig: structs.MailConfig{
				Port:     emailPort,
				Server:   emailServer,
				Username: emailUsername,
				Password: emailPassword,
			},
		}
	}

	log.Println("[Data Manager] Connecting to a SQL DB, please wait...")

	// Open database
	db, err := sql.Open(sqlDriver, sqlUrl)
	if err != nil {
		log.Fatal("[Data Manager] Failed to connect to a SQL DB: ", err)
	}

	// Verify connection to database
	err = db.Ping()
	if err != nil {
		log.Fatal("[Data Manager] Failed to connect to a SQL DB: ", err)
	}

	log.Println("[Data Manager] Successfully connected to a SQL DB.")

	// Return manager
	return &Manager{
		ServerNickname:       serverNickname,
		DB:                   db,
		Ctx:                  ctx,
		AuthlessMode:         false,
		UseInMemoryClientMgr: useInMemoryClientMgr,
		AuthlessUserMap:      nil, // Not used in authless mode
		EnableEmail:          enableEmail,
		MailConfig: structs.MailConfig{
			Port:     emailPort,
			Server:   emailServer,
			Username: emailUsername,
			Password: emailPassword,
		},
	}
}
