package maps

import (
	"net/http"
	"strconv"

	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"
	customError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
	"github.com/GooDu-dev/gd-practical-project-backend/utils/log"
	"github.com/gin-gonic/gin"
)

type MapEndpoint struct {
	Service *MapService
}

func NewMapEndpoint() *MapEndpoint {
	var service MapService
	return &MapEndpoint{
		Service: service.InitService(),
	}
}

func (e *MapEndpoint) GetMap(c *gin.Context) {

	building := c.Query("building")
	floor := c.Query("floor")

	building_id, err := strconv.Atoi(building)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}
	if common.IsDefaultValueOrNil(building_id) || common.IsDefaultValueOrNil(floor) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]interface{}{"building": building_id, "floor": floor})
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	floor_id, err := strconv.Atoi(floor)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.InvalidRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}
	if common.IsDefaultValueOrNil(floor_id) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"floor_id": floor_id})
		status, res := customError.MissingRequestError.ErrorResponse()
		c.JSON(status, res)
		return
	}

	request := GetMapRequest{
		BuildingID: building_id,
		Floor:      floor_id,
	}

	if err := common.DeepIsDefaultValueOrNil(request); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.GetErrorResponse(err)
		c.JSON(status, res)
		return
	}

	response, err := e.Service.GetMap(request)

	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		status, res := customError.GetErrorResponse(err)
		c.JSON(status, res)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)
}
