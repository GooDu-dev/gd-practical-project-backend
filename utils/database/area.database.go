package database

import (
	"gorm.io/gorm"
)

type AreaModel struct {
	gorm.Model
	Latitude   float32 `gorm:"index:uniqueIndex; type:float(8);"`
	Longitude  float32 `gorm:"index:uniqueIndex; type:float(8);"`
	Width      uint16  `gorm:"index:uniqueIndex;"`
	Height     uint16  `gorm:"index:uniqueIndex;"`
	FloorID    uint
	Floor      FloorModel `gorm:"foreignKey:FloorID"`
	BuildingID uint
	Building   BuildingModel `gorm:"foreignKey:BuildingID"`
	AreaTypeID uint
	AreaType   AreaTypeModel `gorm:"foreignKey:AreaTypeID"`
	Name       string        `gorm:"index:uniqueIndex; type:varchar(24)"`
	X          int           `gorm:"index:uniqueIndex"`
	Y          int           `gorm:"index:uniqueIndex"`
	InAreaID   *int
	AreaFK     *AreaModel `gorm:"foreignKey:InAreaID"`
}

type AreaTypeModel struct {
	gorm.Model
	Name string `gorm:"index:uniqueIndex; type:varchar(24);"`
}

func (AreaModel) TableName() string {
	return "tb_area"
}

func (AreaTypeModel) TableName() string {
	return "tb_area_type"
}
