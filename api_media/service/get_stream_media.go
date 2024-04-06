package service

import (
	"github.com/minio/minio-go/v7"
)

func (s mediaService) GetStreamMedia(bucket string, file string) (*minio.Object, error) {
	return s.m.Read(bucket, file)
}
