package database

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

func (d *Database) GetGame(id string) *types.DeveloperGame {
	var game *types.DeveloperGame
	if d.DB.Preload("Developer").Find(&game, "id = ?", id).Error == gorm.ErrRecordNotFound {
		return nil
	}
	return game
}
