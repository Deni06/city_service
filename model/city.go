package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"time"
)

func (t City) TableName() string {
	return constant.CITY_TABLE_NAME
}

type City struct {
	CityId uint `gorm:"primary_key;AUTO_INCREMENT"`
	CityName string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(200);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProjectId int32 `gorm:"type:int;"`
	ProvinceId int32 `gorm:"type:int;not null"`
}

func(City) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&City{},).Error
	if err != nil {
		return err
	}

	return nil
}
