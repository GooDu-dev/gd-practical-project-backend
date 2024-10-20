package search

import (
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
)

type SearchService struct {
	Model *SearchModel
}

func (s *SearchService) InitService() *SearchService {
	var model SearchModel
	return &SearchService{
		Model: model.InitModel(),
	}
}

func (s *SearchService) GetRoomDetails(request SearchRequest) (_ *Building, err error) {
	var response *Building

	if response, err = s.Model.GetRoomDetails(request.RoomID); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}

func (s *SearchService) GetRoomSearchListFromBuilding(request BuildingSearchRequest) (_ *[]RoomNameSearchList, err error) {
	var response *[]RoomNameSearchList

	if response, err = s.Model.GetRoomSearchListFromBuilding(request.ID); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}

func (s *SearchService) GetBuildingSearchList(request BuildingSearchListRequest) (_ *[]BuildingSearchList, err error) {
	var response *[]BuildingSearchList

	if response, err = s.Model.GetBuildingSearchList(request.Count); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}

func (s *SearchService) GetFloorSearchList(request FloorSearchListRequest) (_ *[]FloorSearchList, err error) {
	var response *[]FloorSearchList

	if response, err = s.Model.GetFloorSearchList(request.BuildingID); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}
