package main

import (
	"api_upload/api/routes"
	"api_upload/repository/minio"
	"api_upload/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := utils.InitMinIODB(); err != nil {
		log.Panic("error to init minio db " + err.Error())
	}
	if err := utils.InitFilmServiceConnection(); err != nil {
		log.Panic("error to init film service " + err.Error())
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

	if err := utils.InitPublisher(); err != nil {
		log.Panic("error to init publisher " + err.Error())
	}

	repository := minio.NewRepository()
	filmServiceConnection := utils.GetFilmServiceConnection()
	adServiceConnection := utils.GetAdServiceConnection()

	publisher := utils.GetPublisher()

	gin.SetMode(gin.ReleaseMode)

	r := routes.InitRoutes(repository, filmServiceConnection, adServiceConnection, publisher)

	err := r.Run()
	if err != nil {
		log.Panic("error to start service " + err.Error())
	}
}
