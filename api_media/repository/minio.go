package repository

import (
	"api_media/utils"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type MinIORepository interface {
	Create(multipart.File, string, string, string) error
	Read(string, string) (*minio.Object, error)
	Update(multipart.File, string, string, string) error
	Delete(string, string) error
}

type minIORepository struct {
	client *minio.Client
}

func NewMinIORepository() MinIORepository {
	return &minIORepository{client: utils.GetMinIOClient()}
}
