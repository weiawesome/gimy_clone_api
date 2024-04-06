package service

import (
	"log"
	"os"
)

func deleteResource(filePath string) {
	err := os.RemoveAll(filePath)
	if err != nil {
		log.Println("error to delete file " + err.Error())
	}
}
