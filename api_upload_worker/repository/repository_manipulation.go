package repository

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func (m *repository) Create(data multipart.File, bucket string, file string, fileSize int64, contentType string) error {
	exists, err := m.client.BucketExists(context.Background(), bucket)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	_, err = m.client.PutObject(context.Background(), bucket, file, data, fileSize, minio.PutObjectOptions{ContentType: contentType})
	return err
}
func (m *repository) FCreate(bucket string, filePath string, file string) error {
	exists, err := m.client.BucketExists(context.Background(), bucket)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	_, err = m.client.FPutObject(context.Background(), bucket, file, filePath, minio.PutObjectOptions{})
	return err
}
func (m *repository) Read(bucket string, file string) (*minio.Object, error) {
	return m.client.GetObject(context.Background(), bucket, file, minio.GetObjectOptions{})
}
func (m *repository) FRead(bucket string, file string, filePath string) error {
	return m.client.FGetObject(context.Background(), bucket, file, filePath, minio.GetObjectOptions{})
}
func (m *repository) Update(data multipart.File, bucket string, file string, contentType string) error {
	var err error

	_, err = m.client.StatObject(context.Background(), bucket, file, minio.StatObjectOptions{})
	if err != nil {
		return err
	}

	_, err = m.client.PutObject(context.Background(), bucket, file, data, -1, minio.PutObjectOptions{ContentType: contentType})
	return err
}
func (m *repository) Delete(bucket string, file string) error {
	return m.client.RemoveObject(context.Background(), bucket, file, minio.RemoveObjectOptions{})
}
