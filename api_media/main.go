package main

import (
	"api_media/api/route"
	"api_media/repository"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := utils.InitMinIODB(); err != nil {
		log.Panic(err)
	}
	if err := utils.InitConnection(); err != nil {
		log.Panic(err)
	}
	defer func() {
		err := utils.CloseConnection()
		if err != nil {
			return
		}
	}()

	r := repository.NewMinIORepository()
	c := utils.GetConnection()

	gin.SetMode(gin.ReleaseMode)
	router := route.InitRoutes(&r, c)

	err := router.Run()
	if err != nil {
		log.Panic("error to start service " + err.Error())
	}
}
