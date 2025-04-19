package database

import (
	"gorm.io/gorm"

	"github.com/cloudlink-omega/storage/pkg/types"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) RunMigrations() error {
	if err := d.DB.AutoMigrate(
		&types.Developer{},
		&types.DeveloperGame{},
		&types.DeveloperMember{},
		&types.DeveloperArt{},
		&types.UserGameSave{},
		&types.GameFeature{},
		&types.GameArt{},
	); err != nil {
		return err
	}

	// Seed the database with the test Developer and test Game IDs
	if err := d.DB.Model(&types.Developer{}).FirstOrCreate(&types.Developer{
		Name: "Test Developer",
		ID:   "01HNPHQM5SPAG43J68R3NRX4M6",
	}).Error; err != nil {
		return err
	}

	if err := d.DB.Model(&types.DeveloperGame{}).FirstOrCreate(&types.DeveloperGame{
		Name:        "Test Game",
		ID:          "01HNPHRWS0N0AYMM5K4HN31V4W",
		DeveloperID: "01HNPHQM5SPAG43J68R3NRX4M6",
	}).Error; err != nil {
		return err
	}

	return nil
}
