package database

import (
	"gorm.io/gorm"
)

type FloorModel struct {
	gorm.Model
	X          int           `gorm:""`
	Y          int           `gorm:""`
	Name       string        `gorm:""`
	BuildingID uint          `gorm:""`
	BuildingFK BuildingModel `gorm:"foreignKey: BuildingID"`
}

func (FloorModel) TableName() string {
	return "tb_floor"
}
