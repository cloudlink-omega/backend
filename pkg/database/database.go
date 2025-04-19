package database

import (
	"gorm.io/gorm"

	"github.com/cloudlink-omega/storage/pkg/types"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) RunMigrations() error {
	return d.DB.AutoMigrate(
		&types.Developer{},
		&types.DeveloperGame{},
		&types.DeveloperMember{},
		&types.UserGameSave{},
	)
}
