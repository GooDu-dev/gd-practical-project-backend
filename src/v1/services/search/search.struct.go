package search

type SearchRequest struct {
	BuildingID int  `json:"building"`
	Floor      *int `json:"floor"`
	RoomID     int  `json:"room"`
}

type SearchResponse struct {
	BuildingID        int    `json:"building_id"`
	BuildingName      string `json:"building_name"`
	BuildingLatitude  string `json:"building_latitude"`
	BuildingLongitude string `json:"building_longitude"`
	AreaID            int    `json:"area_id"`
	AreaName          string `json:"area_name"`
	AreaLatitude      string `json:"area_latitude"`
	AreaLongitude     string `json:"area_longitude"`
	AreaWidth         int    `json:"area_width"`
	AreaHeight        int    `json:"area_height"`
	AreaFloorID       int    `json:"area_floor_id"`
	AreaFloor         string `json:"area_floor"`
	AreaTypeID        int    `json:"area_type_id"`
	AreaType          string `json:"area_type"`
}

type Building struct {
	ID         int    `json:"building_id"`
	Name       string `json:"name"`
	Latitude   string `json:"latitude"`
	Longtitude string `json:"longtitude"`
	Area       []Area `json:"area"`
}

type Area struct {
	ID        int             `json:"area_id"`
	Name      string          `json:"name"`
	Latittude string          `json:"latitude"`
	Longitude string          `json:"longitude"`
	Width     int             `json:"width"`
	Height    int             `json:"height"`
	Floor     FloorSearchList `json:"floor"`
	AreaType  string          `json:"area_type"`
}

type BuildingSearchRequest struct {
	ID    int  `json:"building_id"`
	Floor *int `json:"building_floor"`
}

type RoomNameSearchList struct {
	RoomID    int    `json:"room_id"`
	RoomName  string `json:"room_name"`
	RoomFloor int    `json:"room_floor"`
}

type BuildingSearchListRequest struct {
	Count int `json:"count"`
}

type BuildingSearchList struct {
	BuildingID   int    `json:"building_id"`
	BuildingName string `json:"building_name"`
}

type FloorSearchListRequest struct {
	BuildingID int `json:"building_id"`
}

type FloorSearchList struct {
	ID    int    `json:"floor_id"`
	Floor string `json:"floor"`
}
