package repository

import (
	"api_upload_worker/utils"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type Repository interface {
	Create(multipart.File, string, string, int64, string) error
	FCreate(string, string, string) error
	Read(string, string) (*minio.Object, error)
	FRead(string, string, string) error
	Update(multipart.File, string, string, string) error
	Delete(string, string) error
}

type repository struct {
	client *minio.Client
}

func NewRepository() Repository {
	return &repository{client: utils.GetMinIOClient()}
}
