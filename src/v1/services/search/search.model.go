package search

import (
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/database"
	customError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
)

type SearchModel struct{}

func (m *SearchModel) InitModel() *SearchModel {
	return &SearchModel{}
}

func (m *SearchModel) GetAreaDetails(area_id int, area_type_id int) (*Building, error) {

	response := []SearchResponse{}

	result := database.DB.Table("tb_building").
		Select("tb_building.id as building_id",
			"tb_building.name as building_name",
			"tb_building.latitude as building_latitude",
			"tb_building.longitude as building_longitude",
			"tb_area.id as area_id",
			"tb_area.name as area_name",
			"tb_area.latitude as area_latitude",
			"tb_area.longitude as area_longitude",
			"tb_area.width as area_width",
			"tb_area.height as area_height",
			"tb_area.in_area_id as in_area_id",
			"tb_area.x as area_x",
			"tb_area.y as area_y",
			"tb_floor.id as area_floor_id",
			"tb_floor.name as area_floor",
			"tb_area.area_type_id as area_type_id",
			"tb_area_type.name as area_type").
		Joins("INNER JOIN tb_area ON tb_building.id = tb_area.building_id").
		Joins("INNER JOIN tb_area_type ON tb_area.area_type_id = tb_area_type.id").
		Joins("INNER JOIN tb_floor ON tb_area.floor_id = tb_floor.id").
		Where("tb_area.id = ? AND tb_area_type.id = ?", area_id, area_type_id).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	output := Building{
		ID:         response[0].BuildingID,
		Name:       response[0].BuildingName,
		Latitude:   response[0].BuildingLatitude,
		Longtitude: response[0].BuildingLongitude,
		Area:       []Area{},
	}

	for _, res := range response {
		output.Area = append(output.Area, Area{
			ID:        res.AreaID,
			Name:      res.AreaName,
			Latittude: res.AreaLatitude,
			Longitude: res.AreaLongitude,
			Width:     res.AreaWidth,
			Height:    res.AreaHeight,
			Floor: FloorSearchList{
				ID:    res.AreaFloorID,
				Floor: res.AreaFloor,
			},
			AreaType: res.AreaType,
			InAreaID: res.InAreaID,
			AreaX:    res.AreaX,
			AreaY:    res.AreaY,
		})
	}

	return &output, nil
}

func (m *SearchModel) GetRoomSearchListFromBuilding(id int) (_ *[]RoomNameSearchList, err error) {
	response := []RoomNameSearchList{}

	result := database.DB.Table("tb_area").
		Select(
			"tb_area.id as room_id",
			"tb_area.name as room_name",
			"tb_floor.name as room_floor",
		).
		Joins("INNER JOIN tb_building ON tb_area.building_id = tb_building.id").
		Joins("INNER JOIN tb_floor ON tb_area.floor_id = tb_floor.id").
		Where("tb_building.id = ? AND tb_area.tb_area_id = 3", id).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil

}

func (m *SearchModel) GetBuildingSearchList(count int) (_ *[]BuildingSearchList, err error) {
	var response []BuildingSearchList

	result := database.DB.Table("tb_building").
		Select(
			"tb_building.id as building_id",
			"tb_building.name as building_name",
		).
		Limit(count).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil
}

func (m *SearchModel) GetFloorSearchList(building_id int) (_ *[]FloorSearchList, err error) {
	var response []FloorSearchList

	result := database.DB.Table("tb_floor").
		Select(
			"tb_floor.id as id",
			"tb_floor.name as floor",
		).
		Where("tb_floor.building_id = ?", building_id).
		Limit(10).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, result.Error
	}

	return &response, nil
}
