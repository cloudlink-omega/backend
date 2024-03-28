package data

import (
	"context"
	"database/sql"
	"log"
)

type Manager struct {
	Ctx context.Context
	DB  *sql.DB
}

func New(sqlDriver string, sqlUrl string) *Manager {
	// Create background context
	ctx := context.Background()

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
		DB:  db,
		Ctx: ctx,
	}
}
