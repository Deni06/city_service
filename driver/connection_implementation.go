package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
)

type PostgreeImplementation struct {

}

type MysqlImplementation struct {

}

type SqliteImplementation struct {

}

type SqlServerImplementation struct {

}

func (p PostgreeImplementation)newDBGorm() (*gorm.DB, error){
	psqlInfo := ""
	fmt.Println("masuk postgree")
	if parameter.UseCli {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			parameter.Host, parameter.HostParam, parameter.User, parameter.Password, parameter.DbName, parameter.SslMode)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			constant.HOST_POSTGRESQL, constant.PORT_POSTGRESQL, constant.POSTGRESQL_USERNAME, constant.POSTGRESQL_PASSWORD,
			constant.POSTGRESQL_DBNAME, constant.SSL_MODE,
		)
	}

	db, err := gorm.Open(constant.POSTGRESQL_DIALECT, psqlInfo)
	if err!=nil{
		return nil, err
	}
	dbInstanceGorm = *db
	return db, nil
}

func (m MysqlImplementation)newDBGorm() (*gorm.DB, error){
	mysqlInfo := ""
	fmt.Println("masuk mysql")
	if parameter.UseCli {
		mysqlInfo = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True", parameter.User,parameter.Password,parameter.Host,parameter.HostParam,parameter.DbName)
	} else {
		mysqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			constant.HOST_POSTGRESQL, constant.PORT_POSTGRESQL, constant.POSTGRESQL_USERNAME, constant.POSTGRESQL_PASSWORD,
			constant.POSTGRESQL_DBNAME, constant.SSL_MODE,
		)
	}

	db, err := gorm.Open(constant.MYSQL_DIALECT, mysqlInfo)
	if err!=nil{
		return nil, err
	}
	dbInstanceGorm = *db
	return db, nil
}

func (s SqliteImplementation)newDBGorm() (*gorm.DB, error){
	sqliteInfo := ""

	if parameter.UseCli {
		sqliteInfo = fmt.Sprintf("%v", parameter.DbPath)
	} else {
		sqliteInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			constant.HOST_POSTGRESQL, constant.PORT_POSTGRESQL, constant.POSTGRESQL_USERNAME, constant.POSTGRESQL_PASSWORD,
			constant.POSTGRESQL_DBNAME, constant.SSL_MODE,
		)
	}

	db, err := gorm.Open(constant.SQLITE_DIALECT, sqliteInfo)
	if err!=nil{
		return nil, err
	}
	dbInstanceGorm = *db
	return db, nil
}

func (ss SqlServerImplementation)newDBGorm() (*gorm.DB, error){
	fmt.Println("masuk sql server")
	sqlserverInfo := ""
	if parameter.UseCli {
		sqlserverInfo = fmt.Sprintf("sqlserver://%v:%v@%v%v?database=%v", parameter.User,parameter.Password,parameter.Host,parameter.HostParam,parameter.DbName)
	} else {
		sqlserverInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			constant.HOST_POSTGRESQL, constant.PORT_POSTGRESQL, constant.POSTGRESQL_USERNAME, constant.POSTGRESQL_PASSWORD,
			constant.POSTGRESQL_DBNAME, constant.SSL_MODE,
		)
	}

	db, err := gorm.Open(constant.SQL_SERVER_DIALECT, sqlserverInfo)
	if err!=nil{
		return nil, err
	}
	dbInstanceGorm = *db
	return db, nil
}