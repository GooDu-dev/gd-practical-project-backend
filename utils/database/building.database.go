package database

import (
	"gorm.io/gorm"
)

type BuildingModel struct {
	gorm.Model
	Name      string  `gorm:"size:255; index:uniqueIndex;"`
	Latitude  float32 `gorm:"type:float(8); index:uniqueIndex;"`
	Longitude float32 `gorm:"type:float(8); index:uniqueIndex;"`
}

func (BuildingModel) TableName() string {
	return "tb_building"
}
