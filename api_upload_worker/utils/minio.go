package utils

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func InitMinIODB() error {
	var err error

	endpoint := EnvMinIOAddress()
	accessKeyID := EnvMinIOAccessKeyID()
	secretAccessKey := EnvMinIOAccessKeySecret()
	token := EnvMinIOAccessKeyToken()
	useSSL := GetMinIOUseSSL()

	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, token),
		Secure: useSSL,
	})

	return err
}

func GetMinIOClient() *minio.Client {
	return client
}
