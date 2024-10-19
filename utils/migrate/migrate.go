package main

import (
	"github.com/GooDu-dev/gd-practical-project-backend/utils/database"
	customLog "github.com/GooDu-dev/gd-practical-project-backend/utils/log"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
	customLog.InitLogger()
	database.ConDB()
}

func main() {
	database.DB.AutoMigrate(&database.BuildingModel{})
	database.DB.AutoMigrate(&database.AreaModel{})
	database.DB.AutoMigrate(&database.AreaTypeModel{})
}
