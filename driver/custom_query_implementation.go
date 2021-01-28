package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

func (p PostgreeImplementation)GetInsertQueryCity(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	query = query+fmt.Sprintf(" Returning %v ",primaryKey)
	newQuery := db.Raw(query)
	return newQuery
}

func (m MysqlImplementation)GetInsertQueryCity(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	newQuery := db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Exec(query)
	}, func(db *gorm.DB) *gorm.DB {
		return db.Select("LAST_INSERT_ID()").Table(tableName)
	})
	return newQuery
}

func (s SqliteImplementation)GetInsertQueryCity(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	newQuery := db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Exec(query)
	}, func(db *gorm.DB) *gorm.DB {
		return db.Select("last_insert_rowid()").Table(tableName)
	})
	return newQuery
}

func (ss SqlServerImplementation)GetInsertQueryCity(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	query = query+" select ID = convert(bigint, SCOPE_IDENTITY()) "
	newQuery := db.Raw(query)
	return newQuery
}

func (p PostgreeImplementation)GetQueryCity(in *string){

}

func (m MysqlImplementation)GetQueryCity(in *string){
}

func (s SqliteImplementation)GetQueryCity(in *string){
}

func (ss SqlServerImplementation)GetQueryCity(in *string){
	val := strings.Replace(*in,"false", "0",-1)
	*in = val
}