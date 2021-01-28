package driver

import (
	"fmt"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
)

func GetConnectionHandler(dialect string) (GetConnectionGeneric, error) {
	switch dialect {
	case constant.POSTGRESQL_DIALECT:
		return new(PostgreeImplementation), nil
	case constant.MYSQL_DIALECT:
		return new(MysqlImplementation), nil
	case constant.SQLITE_DIALECT:
		return new(SqliteImplementation), nil
	case constant.SQL_SERVER_DIALECT:
		return new(SqlServerImplementation), nil
	default:
		break
	}
	return nil, util.UnhandledError{ErrorMessage: fmt.Sprintf("dialect %v doesnt have implementation", dialect)}
}

func GetCustomQuery(dialect string) (CustomQueryInterface, error) {
	switch dialect {
	case constant.POSTGRESQL_DIALECT:
		return new(PostgreeImplementation), nil
	case constant.MYSQL_DIALECT:
		return new(MysqlImplementation), nil
	case constant.SQLITE_DIALECT:
		return new(SqliteImplementation), nil
	case constant.SQL_SERVER_DIALECT:
		return new(SqlServerImplementation), nil
	default:
		break
	}
	return nil, util.UnhandledError{ErrorMessage: fmt.Sprintf("dialect %v doesnt have implementation", dialect)}
}