package database

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

func (d *Database) GetGame(id string) (game *types.DeveloperGame) {
	res := d.DB.Preload("Developer").Where("id = ?", id).First(&game)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(res.Error)
		}
	}
	return game
}
