package minio

import (
	"api_upload/utils"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type Repository interface {
	Create(multipart.File, string, string, int64, string) error
	Read(string, string) (*minio.Object, error)
	Update(multipart.File, string, string, string) error
	Delete(string, string) error
}

type repository struct {
	client *minio.Client
}

func NewRepository() Repository {
	return &repository{client: utils.GetMinIOClient()}
}
