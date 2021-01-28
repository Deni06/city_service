package driver

import (
	"github.com/jinzhu/gorm"
)

type GetConnectionGeneric interface {
	newDBGorm() (*gorm.DB, error)
}

type CustomQueryInterface interface {
	GetInsertQueryCity(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB
	GetQueryCity(*string)
}
