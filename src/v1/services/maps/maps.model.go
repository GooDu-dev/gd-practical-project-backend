package maps

import (
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/database"
	customError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
)

type MapModel struct {
}

func (m *MapModel) InitModel() *MapModel {
	return &MapModel{}
}

func (m *MapModel) GetAreasZone(building_id int, floor_id int, area_type_id int) (*[]Area, error) {
	response := []Area{}

	result := database.DB.
		Table("tb_area").
		Select(
			"tb_area.id as area_id",
			"tb_area.name as name",
			"tb_area.latitude as latitude",
			"tb_area.longitude as longitude",
			"tb_area.width as width",
			"tb_area.height as height",
			"tb_area_type.name as area_type",
			"tb_area.x as x",
			"tb_area.y as y",
			"tb_area.in_area_id as in_area_id",
		).
		Joins("INNER JOIN tb_area_type ON tb_area.area_type_id = tb_area_type.id").
		Where("tb_area.building_id = ? AND tb_area.floor_id = ? AND tb_area.area_type_id = ?", building_id, floor_id, area_type_id).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		err := customError.InternalServerError
		return nil, err
	}

	return &response, nil
}

func (m *MapModel) GetUnwalkableZone(building_id int, floor_id int) (*[]Area, error) {

	var response []Area

	result := database.DB.Table("tb_area").
		Select(
			"tb_area.id as area_id",
			"tb_area.name as name",
			"tb_area.latitude as latitude",
			"tb_area.longitude as longitude",
			"tb_area.width as width",
			"tb_area.height as height",
			"tb_area_type.name as area_type",
			"tb_area.x as x",
			"tb_area.y as y",
		).
		Joins("INNER JOIN tb_area_type ON tb_area.area_type_id = tb_area_type.id").
		Where("tb_area.building_id = ? AND tb_area.floor_id = ? AND tb_area.area_type_id != 1 AND tb_area.area_type_id != 7", building_id, floor_id).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil
}

func (m *MapModel) GetFloorData(floor int) (*FloorData, error) {
	response := FloorData{}

	result := database.DB.Table("tb_floor").
		Select(
			"tb_floor.id as floor_id",
			"tb_floor.name as floor_name",
			"tb_floor.x as size_x",
			"tb_floor.y as size_y",
			"tb_floor.building_id as building_id",
		).
		Where("tb_floor.id = ?", floor).
		Limit(1).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		err := customError.InternalServerError
		return nil, err
	}

	return &response, nil
}

func (m *MapModel) GetRoomEntrance() (*[]Area, error) {

	var response []Area

	result := database.DB.Table("tb_area").
		Select(
			"tb_area.id as area_id",
			"tb_area.name as name",
			"tb_area.latitude as latitude",
			"tb_area.longitude as longitude",
			"tb_area.width as width",
			"tb_area.height as height",
			"tb_area_type.name as area_type",
			"tb_area.x as x",
			"tb_area.y as y",
			"tb_area.in_area_id as in_area_id",
		).
		Joins("INNER JOIN tb_area_type ON tb_area_type.id = tb_area.area_type_id").
		Where("tb_area_type.id = 7").
		Limit(10).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil
}

func (m *MapModel) GetFloorBound(floor_id int) (*GetFloorBoundResponse, error) {

	query := struct {
		ID         int    `gorm:"column:id"`
		Name       string `gorm:"column:name"`
		BuildingID int    `gorm:"column:building_id"`
		X          int    `gorm:"column:x"`
		Y          int    `gorm:"column:y"`
		TLLat      string `gorm:"column:tl_lat"`
		TLLng      string `gorm:"column:tl_lng"`
		TRLat      string `gorm:"column:tr_lat"`
		TRLng      string `gorm:"column:tr_lng"`
		BLLat      string `gorm:"column:bl_lat"`
		BLLng      string `gorm:"column:bl_lng"`
		BRLat      string `gorm:"column:br_lat"`
		BRLng      string `gorm:"column:br_lng"`
	}{}

	result := database.DB.Table("tb_floor").
		Select(
			"tb_floor.id as id",
			"tb_floor.name as name",
			"tb_floor.x as x",
			"tb_floor.y as y",
			"tb_floor.building_id as building_id",
			"tb_floor.tl_lat as tl_lat",
			"tb_floor.tl_lng as tl_lng",
			"tb_floor.tr_lat as tr_lat",
			"tb_floor.tr_lng as tr_lng",
			"tb_floor.bl_lat as bl_lat",
			"tb_floor.bl_lng as bl_lng",
			"tb_floor.br_lat as br_lat",
			"tb_floor.br_lng as br_lng",
		).
		Where("tb_floor.id = ?", floor_id).
		Find(&query)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	response := GetFloorBoundResponse{
		ID:         query.ID,
		Name:       query.Name,
		BuildingID: query.BuildingID,
		X:          query.X,
		Y:          query.Y,
		GISData: []GISData{
			{
				Lat: query.TLLat,
				Lng: query.TLLng,
			},
			{
				Lat: query.TRLat,
				Lng: query.TRLng,
			},
			{
				Lat: query.BLLat,
				Lng: query.BLLng,
			},
			{
				Lat: query.BRLat,
				Lng: query.BRLng,
			},
		},
	}

	return &response, nil
}
