package main

import (
	"api_affair/api/routes"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := utils.InitFilmServiceConnection(); err != nil {
		log.Println(err.Error())
	}
	defer func() {
		err := utils.CloseFilmServiceConnection()
		if err != nil {
			return
		}
	}()
	if err := utils.InitAdServiceConnection(); err != nil {
		log.Panic("error to init ad service " + err.Error())
	}
	defer func() {
		err := utils.CloseAdServiceConnection()
		if err != nil {
			return
		}
	}()

	filmConn := utils.GetFilmServiceConnection()
	adConn := utils.GetAdServiceConnection()

	gin.SetMode(gin.ReleaseMode)
	r := routes.InitRoutes(filmConn, adConn)

	err := r.Run()
	if err != nil {
		log.Panic("error to start service " + err.Error())
	}

}
