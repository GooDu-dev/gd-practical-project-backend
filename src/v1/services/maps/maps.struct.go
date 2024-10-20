package maps

type GetMapRequest struct {
	BuildingID int `json:"building_id"`
	Floor      int `json:"floor"`
}

type GetMapResponse struct {
	ID         int    `json:"floor_id"`
	Name       string `json:"floor_name"`
	Rooms      []Area `json:"rooms"`
	DangerZone []Area `json:"danger_zone"`
	SizeX      int    `json:"size_x"`
	SizeY      int    `json:"size_y"`
}

type FloorData struct {
	FloorID    int    `json:"floor_id"`
	FloorName  string `json:"floor_name"`
	SizeX      int    `json:"size_x"`
	SizeY      int    `json:"size_y"`
	BuildingID int    `json:"building_id"`
}

type Area struct {
	AreaID    int    `json:"area_id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	AreaType  string `json:"area_type"`
}
