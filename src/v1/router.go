package v1

import (
	"net/http"
	"time"

	validator "github.com/GooDu-dev/gd-practical-project-backend/src/v1/middlewares"
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/services/maps"
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/services/search"
	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/services/test"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Path        string
	Validation  gin.HandlerFunc
	Endpoint    gin.HandlerFunc
}

type Router struct {
	testService   []route
	searchService []route
	mapService    []route
}

func (r Router) InitRouter() http.Handler {
	testEndPoint := test.NewEndPoint()
	r.testService = []route{
		{
			Name:        "[GET] : Ping",
			Description: "If works, 'Pong' will returns",
			Method:      http.MethodGet,
			Path:        "/test/ping",
			Validation:  validator.NoValidation,
			Endpoint:    testEndPoint.GetTestHealth,
		},
	}

	searchEndPoint := search.NewEndPoint()
	r.searchService = []route{
		{
			Name:        "[GET] : Get all rooms in the building",
			Description: "Get all room in the building, required data is building name and room name, floor is not requirement",
			Method:      http.MethodGet,
			Path:        "/search",
			Validation:  validator.NoValidation,
			Endpoint:    searchEndPoint.GetRoomDetails,
		},
		{
			Name:        "[GET] : Get room name and floor list from building",
			Description: "Get only room name and floor from building id or name, use for search engine",
			Method:      http.MethodGet,
			Path:        "/search/roomlist",
			Validation:  validator.NoValidation,
			Endpoint:    searchEndPoint.GetRoomNameSearchList,
		},
		{
			Name:        "[GET] : Get top N of building list",
			Description: "Get N amount of building, use this for building name list in search bar",
			Method:      http.MethodGet,
			Path:        "/search/buildinglist",
			Validation:  validator.NoValidation,
			Endpoint:    searchEndPoint.GetBuildingSearchList,
		},
		{
			Name:        "[GET] : Get building's floor",
			Description: "Get building's floor, no duplicated",
			Method:      http.MethodGet,
			Path:        "/search/floorlist",
			Validation:  validator.NoValidation,
			Endpoint:    searchEndPoint.GetFloorSearchList,
		},
	}
	mapEndpoint := maps.NewMapEndpoint()
	r.mapService = []route{
		{
			Name:        "[GET] : Get all area in user's position",
			Description: "Get by building_id and floor",
			Method:      http.MethodGet,
			Path:        "/maps",
			Validation:  validator.NoValidation,
			Endpoint:    mapEndpoint.GetMap,
		},
		{
			Name:        "[GET] Get floor boundaries",
			Description: "Get lat and lng of 4 cornors in specified floor",
			Method:      http.MethodGet,
			Path:        "/maps/bounds/gis",
			Validation:  validator.NoValidation,
			Endpoint:    mapEndpoint.GetFloorBound,
		},
	}

	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://gd-practical-project-frontend-alpha.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{utils.CONTENT_TYPE, utils.CONTENT_CODE, utils.ACCESS_CONTROL, utils.SOURCE_CONTROL},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	mainRoute := ro.Group(utils.PATH)
	for _, e := range r.testService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	for _, e := range r.searchService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	for _, e := range r.mapService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	return ro
}
