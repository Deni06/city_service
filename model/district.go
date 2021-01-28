package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"time"
)

func (t District) TableName() string {
	return constant.DISTRICT_TABLE_NAME
}

type District struct {
	DistrictId uint `gorm:"primary_key;AUTO_INCREMENT"`
	DistrictName string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CityId int32 `gorm:"type:int;not null"`
}

func(District) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&District{},).Error
	if err != nil {
		return err
	}

	return nil
}
