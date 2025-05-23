package database

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

type Database struct {
	DB    *gorm.DB
	Cache *types.DBCache
}
