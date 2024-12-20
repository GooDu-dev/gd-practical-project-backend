package search

import (
	"net/http"
	"strconv"

	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	customError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
	"github.com/gin-gonic/gin"
)

type SearchEndPoint struct {
	Services *SearchService
}

func NewEndPoint() *SearchEndPoint {
	var service SearchService
	return &SearchEndPoint{
		Services: service.InitService(),
	}
}

func (e *SearchEndPoint) GetRoomDetails(c *gin.Context) {

	params := map[string]string{
		"room":      c.Query("room"),
		"restroom":  c.Query("restroom"),
		"booth":     c.Query("booth"),
		"connector": c.Query("connector"),
	}

	var selectedParam, selectedValue string

	for key, value := range params {
		if value != "" {
			if selectedParam != "" {
				log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), params)
				status, res := customError.InvalidRequestError.ErrorResponse()
				c.JSON(status, res)
				return
			}
			selectedParam = key
			selectedValue = value
		}
	}

	if common.IsDefaultValueOrNil(selectedParam) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]string{"params": selectedParam})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	id, err := strconv.Atoi(selectedValue)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}
	if common.IsDefaultValueOrNil(id) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"id": id})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	request := AreaDetailsRequest{
		AreaID: id,
	}

	var response *Building

	switch selectedParam {
	case "room":
		response, err = e.Services.GetRoomDetails(request)
	case "booth":
		response, err = e.Services.GetBoothDetails(request)
	case "restroom":
		response, err = e.Services.GetRestroomDetails(request)
	case "connector":
		response, err = e.Services.GetConnectorDetails(request)
	}

	if err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), res)
		c.JSON(status, res)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)

}

func (e *SearchEndPoint) GetRoomNameSearchList(c *gin.Context) {

	building := c.Query("building")
	floor := c.Query("floor")

	building_id, err := strconv.Atoi(building)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	if common.IsDefaultValueOrNil(building_id) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"building_id": building_id})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	var floor_id *int
	if _floor, err := strconv.Atoi(floor); err != nil {
		log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), err)
	} else if _floor == 0 {
		floor_id = nil
	} else {
		floor_id = &_floor
	}

	request := BuildingSearchRequest{
		ID:    building_id,
		Floor: floor_id,
	}

	var response *[]RoomNameSearchList
	if response, err = e.Services.GetRoomSearchListFromBuilding(request); err != nil {
		status, response := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), response)
		c.JSON(status, response)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)

}

func (e *SearchEndPoint) GetBuildingSearchList(c *gin.Context) {

	_count := c.Query("count")

	count, err := strconv.Atoi(_count)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	if common.IsDefaultValueOrNil(count) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"count": count})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	request := BuildingSearchListRequest{
		Count: count,
	}

	var response *[]BuildingSearchList
	if response, err = e.Services.GetBuildingSearchList(request); err != nil {
		status, response := customError.MappingRequestBodyError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, response)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)
}

func (e *SearchEndPoint) GetFloorSearchList(c *gin.Context) {

	building := c.Query("building")

	building_id, err := strconv.Atoi(building)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	if common.IsDefaultValueOrNil(building_id) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"building_id": building_id})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	request := FloorSearchListRequest{
		BuildingID: building_id,
	}

	var response *[]FloorSearchList
	if response, err = e.Services.GetFloorSearchList(request); err != nil {
		status, response := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, response)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)
}
