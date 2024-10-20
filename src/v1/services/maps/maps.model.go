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

func (m *MapModel) GetAreasZone(building_id int, floor int) (*[]Area, error) {
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
		).
		Joins("INNER JOIN tb_area_type ON tb_area.area_type_id = tb_area_type.id").
		Where("tb_area.building_id = ? AND tb_area.floor_id = ? AND tb_area.area_type_id = 1", building_id, floor).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		err := customError.InternalServerError
		return nil, err
	}

	return &response, nil
}

func (m *MapModel) GetDangerZone(building_id int, floor int) (*[]Area, error) {
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
		).
		Joins("INNER JOIN tb_area_type ON tb_area.area_type_id = tb_area_type.id").
		Where("tb_area.building_id = ? AND tb_area.floor_id = ? AND tb_area.area_type_id = 2", building_id, floor).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		err := customError.InternalServerError
		return nil, err
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
