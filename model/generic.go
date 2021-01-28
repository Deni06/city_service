package model

import "github.com/jinzhu/gorm"

type MigrateInterface interface {
	Migrate(db *gorm.DB)error
}
