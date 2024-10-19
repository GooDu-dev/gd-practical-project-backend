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

	rooms, err := s.model.GetAreasZone(request.BuildingID, request.Floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	dangerZone, err := s.model.GetDangerZone(request.BuildingID, request.Floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	floor_data, err := s.model.GetFloorData(request.Floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	response = &GetMapResponse{
		ID:         floor_data.FloorID,
		Rooms:      *rooms,
		DangerZone: *dangerZone,
		SizeX:      floor_data.SizeX,
		SizeY:      floor_data.SizeY,
	}

	return response, nil

}
