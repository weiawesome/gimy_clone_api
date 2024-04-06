package main

import (
	"api_upload_worker/api/consumer"
	"api_upload_worker/utils"
	"log"
)

func main() {
	if err := utils.InitMinIODB(); err != nil {
		log.Panic("error to init minio " + err.Error())
		return
	}

	if err := utils.InitAdServiceConnection(); err != nil {
		log.Panic("error to init ad service " + err.Error())
		return
	}
	defer func() {
		err := utils.CloseAdServiceConnection()
		if err != nil {
			return
		}
	}()
	if err := utils.InitFilmServiceConnection(); err != nil {
		log.Panic("error to init film service " + err.Error())
		return
	}
	defer func() {
		err := utils.CloseFilmServiceConnection()
		if err != nil {
			return
		}
	}()

	if err := consumer.InitAdConsumers(); err != nil {
		log.Panic("error to init ad consumers " + err.Error())
		return
	}
	if err := consumer.InitFilmConsumers(); err != nil {
		log.Panic("error to init film consumers " + err.Error())
		return
	}

	select {}
}
