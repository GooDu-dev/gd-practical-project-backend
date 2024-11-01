package maps

import (
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
)

type MapService struct {
	model *MapModel
}

func (s *MapService) InitService() *MapService {
	var model MapModel
	return &MapService{
		model: model.InitModel(),
	}
}

func (s *MapService) GetMap(request GetMapRequest) (response *GetMapResponse, err error) {

	var area_type_room int = 1
	rooms, err := s.model.GetAreasZone(request.BuildingID, request.Floor, area_type_room)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	unwalkableZone, err := s.model.GetUnwalkableZone(request.BuildingID, request.Floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	floor_data, err := s.model.GetFloorData(request.Floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	entrances, err := s.model.GetRoomEntrance()
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	areaRooms := []Room{}
	unwalkableRooms := []Room{}

	for _, room := range *rooms {
		r := Room{
			Area: room,
		}
		e := []Area{}
		for _, ent := range *entrances {
			if *(ent.InAreaID) == room.AreaID {
				e = append(e, ent)
			}
		}
		r.Entrance = e
		areaRooms = append(areaRooms, r)
	}

	for _, zone := range *unwalkableZone {
		log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), zone)
		r := Room{
			Area: zone,
		}
		e := []Area{}
		for _, ent := range *entrances {
			if zone.AreaType == "danger_zone" {
				break
			}
			if *(ent.InAreaID) == zone.AreaID {
				e = append(e, ent)
			}
		}
		r.Entrance = e
		unwalkableRooms = append(unwalkableRooms, r)
	}

	response = &GetMapResponse{
		ID:             floor_data.FloorID,
		Name:           floor_data.FloorName,
		Rooms:          areaRooms,
		UnwalkableZone: unwalkableRooms,
		SizeX:          floor_data.SizeX,
		SizeY:          floor_data.SizeY,
	}

	return response, nil

}

func (s *MapService) GetFloorBound(floor_id int) (response *GetFloorBoundResponse, err error) {

	response, err = s.model.GetFloorBound(floor_id)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}
