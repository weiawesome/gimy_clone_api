package service

import (
	"api_media/repository"
	"github.com/grafov/m3u8"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc"
)

type MediaService interface {
	GetStreamMediaList(string, string) (*m3u8.MediaPlaylist, error)
	GetStreamMedia(string, string) (*minio.Object, error)
	InsertAd(*m3u8.MediaPlaylist, *m3u8.MediaPlaylist)
	RefreshBucketUrl(*m3u8.MediaPlaylist, string)
	RefreshMediaUrl(*m3u8.MediaPlaylist)
	GetAd() (string, string, error)
}

type mediaService struct {
	m repository.MinIORepository
	c *grpc.ClientConn
}

func NewMediaService(m repository.MinIORepository, c *grpc.ClientConn) MediaService {
	return mediaService{m: m, c: c}
}
