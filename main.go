package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	router "github.com/GooDu-dev/gd-practical-project-backend/src/v1"
	"github.com/GooDu-dev/gd-practical-project-backend/utils"

	"github.com/GooDu-dev/gd-practical-project-backend/utils/database"
	customLog "github.com/GooDu-dev/gd-practical-project-backend/utils/log"
)

func init() {
	utils.LoadEnv()
	customLog.InitLogger()
	database.ConDB()
}

func main() {
	r := router.Router{}
	route_handler := r.InitRouter()
	AppSrv := &http.Server{
		Addr:    utils.SERVER_PORT,
		Handler: route_handler,
	}

	go func() {
		var err error = nil
		err = AppSrv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
}
