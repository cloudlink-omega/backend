package database

import (
	"github.com/cloudlink-omega/accounts/pkg/constants"
	"github.com/cloudlink-omega/storage/pkg/bitfield"
	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

func (d *Database) AdminExists() bool {
	var bitfield bitfield.Bitfield8
	bitfield.Set(constants.USER_IS_ADMIN)
	var rows int64
	res := d.DB.Model(&types.User{}).Count(&rows).Where("state >= ?", bitfield)
	return res.Error != gorm.ErrRecordNotFound || (rows > 0 && res.Error == nil)
}

func (d *Database) GetAdmin() *types.User {
	var bitfield bitfield.Bitfield8
	bitfield.Set(constants.USER_IS_ADMIN)
	var res *types.User
	d.DB.Where("state >= ?", bitfield).First(&res)
	return res
}
