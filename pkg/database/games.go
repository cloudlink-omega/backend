package database

import (
	"fmt"
	"math"

	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

func (d *Database) GameExists(id string) bool {
	res := d.DB.Where("id = ?", id).First(&types.DeveloperGame{})

	return res.RowsAffected > 0 && res.Error == nil
}

func (d *Database) GetGame(id string) (game *types.DeveloperGame) {
	if game, ok := d.Cache.Get("game", id); ok {
		return game.(*types.DeveloperGame)
	}

	res := d.DB.Preload("Developer").Preload("Features").Where("id = ?", id).First(&game)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(res.Error)
		}
	}

	d.Cache.Set("game", game, id)
	return game
}

type cached_all struct {
	Games    []*types.DeveloperGame
	Total    int64
	MaxPages int64
}

func (d *Database) GetAllGames(page int, limit int) (games []*types.DeveloperGame, total int64, max_pages int64) {
	cache_key := fmt.Sprintf("page_%d:limit_%d", page, limit)

	if games, ok := d.Cache.Get("allgames", cache_key); ok {
		return games.(cached_all).Games, games.(cached_all).Total, games.(cached_all).MaxPages
	}

	d.DB.Preload("Developer").Preload("Features").Preload("Thumbnail").Limit(limit).Offset(page * limit).Where("state > 0").Order("created_at DESC").Find(&games)
	d.DB.Find(&types.DeveloperGame{}).Count(&total)
	max_pages = int64(math.Ceil(float64(total) / float64(limit)))

	d.Cache.Set("allgames", cached_all{games, total, max_pages}, cache_key)
	return games, total, max_pages
}
