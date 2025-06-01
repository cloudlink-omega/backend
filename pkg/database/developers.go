package database

import "github.com/cloudlink-omega/storage/pkg/types"

func (d *Database) CheckForDeveloperNameConflict(name string) bool {
	res := d.DB.Where("name LIKE ?", name).First(&types.Developer{})

	return res.RowsAffected > 0 && res.Error == nil
}

func (d *Database) DeveloperExists(id string) bool {
	res := d.DB.Where("id = ?", id).First(&types.Developer{})

	return res.RowsAffected > 0 && res.Error == nil
}

func (d *Database) GetDeveloper(id string) *types.Developer {
	var res *types.Developer
	d.DB.Preload("DeveloperMembers").Where("id = ?", id).First(&res)
	return res
}
