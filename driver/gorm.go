package driver

import (
	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
)

var dbInstanceGorm = gorm.DB{}

func InitGorm() (*gorm.DB, error){
	db, err := newDBGorm()
	if err != nil {
		return nil, err
	}
	return db,nil
}

func SetParamGorm(input Parameter){
	parameter = input
}

func newDBGorm() (*gorm.DB, error) {
	if parameter.Dialect == "" {
		parameter.Dialect = constant.POSTGRESQL_DIALECT
	}
	connectionHandler, err := GetConnectionHandler(parameter.Dialect)
	if err != nil {
		return nil, err
	}
	db, err := connectionHandler.newDBGorm()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBInstanceGorm()(gorm.DB,error){
	if dbInstanceGorm .DB == nil{
		return dbInstanceGorm,util.UnhandledError{ErrorMessage:constant.UNHANDLED_ERROR}
	}else{
		return dbInstanceGorm, nil
	}
}