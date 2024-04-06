package repository

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func (m *minIORepository) Create(data multipart.File, bucket string, file string, contentType string) error {
	_, err := m.client.PutObject(context.Background(), bucket, file, data, -1, minio.PutObjectOptions{ContentType: contentType})
	return err
}
func (m *minIORepository) Read(bucket string, file string) (*minio.Object, error) {
	return m.client.GetObject(context.Background(), bucket, file, minio.GetObjectOptions{})
}
func (m *minIORepository) Update(data multipart.File, bucket string, file string, contentType string) error {
	var err error

	_, err = m.client.StatObject(context.Background(), bucket, file, minio.StatObjectOptions{})
	if err != nil {
		return err
	}

	_, err = m.client.PutObject(context.Background(), bucket, file, data, -1, minio.PutObjectOptions{ContentType: contentType})
	return err
}
func (m *minIORepository) Delete(bucket string, file string) error {
	return m.client.RemoveObject(context.Background(), bucket, file, minio.RemoveObjectOptions{})
}
