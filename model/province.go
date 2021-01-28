package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"time"
)

func (t Province) TableName() string {
	return constant.PROVINCE_TABLE_NAME
}

type Province struct {
	ProvinceId uint `gorm:"primary_key;AUTO_INCREMENT"`
	ProvinceName string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func(Province) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Province{},).Error
	if err != nil {
		return err
	}

	return nil
}
