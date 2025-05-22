package database

import "github.com/cloudlink-omega/storage/pkg/types"

func (d *Database) CommentExists(id string) bool {
	res := d.DB.Where("id = ?", id).First(&types.GameComment{})

	return res.RowsAffected > 0 && res.Error == nil
}
