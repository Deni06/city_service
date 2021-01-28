package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/driver"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/model"
	"gopkg.in/gormigrate.v1"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/mssql"

)

var city = model.City{}
var district = model.District{}
var province = model.Province{}

func main()  {
	dbUsage := flag.Bool("db_usage", false, "Config for flag usage")
	dbHost := flag.String("db_host", "", "Config for db host")
	dbUsername := flag.String("db_username", "", "Config for username")
	dbPassword := flag.String("db_password", "", "Config for db password")
	dbName := flag.String("db_name", "", "Config for db name")
	dbSslMode := flag.String("db_ssl_mode", "", "Config for db ssl mode")
	dbDialect := flag.String("db_dialect", "", "Config for db_dialect")
	dbPath := flag.String("db_path", "", "Config for db_path")
	dbHostParam := flag.String("db_host_param", "", "Config db port and db instance, db port ex: :8080 or db instance ex: /sql2014 for sql server" +
		" and db port ex : 8080 for other than sql server")

	flag.Parse()
	driver.SetParam(driver.Parameter{
		UseCli:*dbUsage,
		SslMode:*dbSslMode,
		DbName:*dbName,
		Password:*dbPassword,
		User:*dbUsername,
		HostParam:*dbHostParam,
		Host:*dbHost,
		Dialect:*dbDialect,
		DbPath:*dbPath,
	})
	db,errInitDB := driver.InitGorm()
	if errInitDB!=nil{
		fmt.Print(errInitDB.Error())
	}else{
		m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
			{
				ID: time.Now().UTC().String(),
				Migrate: func(tx *gorm.DB) error {

					err := city.Migrate(tx)
					if err != nil {
						return err
					}
					err = district.Migrate(tx)
					if err != nil {
						return err
					}
					err = province.Migrate(tx)
					if err != nil {
						return err
					}
					return nil
				},
			},

		})

		db.LogMode(true)
		err := m.Migrate()
		if err == nil {
			fmt.Println("Migration did run successfully")
		} else {
			fmt.Printf("Could not migrate: %v", err)
		}
		defer db.Close()
	}
}