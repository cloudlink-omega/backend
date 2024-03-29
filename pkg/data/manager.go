package data

import (
	"context"
	"database/sql"
	"log"
)

type Manager struct {
	Ctx                  context.Context
	DB                   *sql.DB
	AuthlessMode         bool
	UseInMemoryClientMgr bool
	AuthlessUserMap      map[string]string // ULID session token -> username. Used for authless mode.
}

func New(sqlDriver string, sqlUrl string, authlessMode bool, useInMemoryClientMgr bool) *Manager {
	// Create background context
	ctx := context.Background()

	if authlessMode {
		log.Println("[Data Manager] Bypassing DB connection due to authless mode.")

		// Return manager
		return &Manager{
			DB:                   nil,
			Ctx:                  ctx,
			AuthlessMode:         true,
			UseInMemoryClientMgr: useInMemoryClientMgr,
			AuthlessUserMap:      make(map[string]string),
		}
	}

	log.Println("[Data Manager] Connecting to DB, please wait...")

	// Open database
	db, err := sql.Open(sqlDriver, sqlUrl)
	if err != nil {
		log.Fatal("[Data Manager] Failed to connect to DB: ", err)
	}

	// Verify connection to database
	err = db.Ping()
	if err != nil {
		log.Fatal("[Data Manager] Failed to connect to DB: ", err)
	}

	log.Println("[Data Manager] Successfully connected to DB.")

	// Return manager
	return &Manager{
		DB:                   db,
		Ctx:                  ctx,
		AuthlessMode:         false,
		UseInMemoryClientMgr: useInMemoryClientMgr,
		AuthlessUserMap:      nil, // Not used in authless mode
	}
}
