package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
)

type DB struct {
	*sql.DB
}

type Parameter struct {
	UseCli bool
	Host string
	HostParam string
	User string
	Password string
	SslMode string
	DbName string
	Dialect string
	DbPath string
}

var dbInstance = DB{}
var parameter = Parameter{}

func Init() (*DB, error){
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	return db,nil
}

func SetParam(input Parameter){
	parameter = input
}

func newDB() (*DB, error) {
	psqlInfo := ""
	if parameter.UseCli{
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			parameter.Host, parameter.HostParam, parameter.User, parameter.Password, parameter.DbName, parameter.SslMode)
	}else{
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=%s",
			constant.HOST_POSTGRESQL, constant.PORT_POSTGRESQL, constant.POSTGRESQL_USERNAME, constant.POSTGRESQL_PASSWORD, constant.POSTGRESQL_DBNAME, constant.SSL_MODE)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	result := DB{db}
	dbInstance = result
	return &result, nil
}

func GetDBInstance()(DB,error){
	if dbInstance.DB == nil{
		return dbInstance,util.UnhandledError{ErrorMessage:constant.UNHANDLED_ERROR}
	}else{
		return dbInstance, nil
	}
}