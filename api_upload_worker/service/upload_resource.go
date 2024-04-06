package service

import (
	"api_upload_worker/repository"
	"os"
	"path/filepath"
)

func uploadResource(r repository.Repository, bucket string, filePath string, fileName string) error {
	return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() == fileName {
			return err
		}
		return r.FCreate(bucket, path, info.Name())
	})
}
