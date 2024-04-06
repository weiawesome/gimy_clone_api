package service

import (
	"api_upload_worker/repository"
	"os"
)

func downloadResource(r repository.Repository, bucket string, filePath string, fileName string) error {
	err := os.MkdirAll(filePath+"/resource", os.ModePerm)
	if err != nil {
		return err
	}
	return r.FRead(bucket, fileName, filePath+"/"+fileName)
}
