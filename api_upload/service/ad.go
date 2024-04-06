package service

import (
	"api_upload/proto/ad_service"
	"api_upload/repository/minio"
	"api_upload/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mime/multipart"
	"time"
)

type AdService interface {
	CreateAd(string, multipart.File, string, string, int64, time.Time) error
}

type adService struct {
	r minio.Repository
	c *grpc.ClientConn
	p message.Publisher
}

func NewAdService(r minio.Repository, c *grpc.ClientConn, p message.Publisher) AdService {
	return &adService{r: r, c: c, p: p}
}

func (s *adService) CreateAd(adType string, file multipart.File, fileExtension string, contentType string, size int64, expireTime time.Time) error {
	if ad_service.AdType(ad_service.AdType_value[adType]) == ad_service.AdType_UNKNOWN {
		return errors.New("error ad type")
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	bucket := utils.GetAdBucket()
	description := utils.GetOriginalFileDescription()

	err = s.r.Create(file, bucket, id.String()+description+fileExtension, size, contentType)
	if err != nil {
		return err
	}
	if ad_service.AdType(ad_service.AdType_value[adType]) == ad_service.AdType_FILM {
		film := utils.AdInformation{Bucket: bucket, Id: id.String(), FileExtension: fileExtension, ExpireTime: expireTime}
		bytes, err := json.Marshal(film)
		if err != nil {
			return err
		}
		msg := message.NewMessage(watermill.NewUUID(), bytes)
		return s.p.Publish(utils.EnvKafkaAdTopic(), msg)
	}
	client := ad_service.NewAdvertisementClient(s.c)
	ad := ad_service.SaveAdvertisementRequest{Type: ad_service.AdType(ad_service.AdType_value[adType]), Bucket: bucket, File: id.String() + description + fileExtension, ExpireTime: timestamppb.New(expireTime.UTC())}
	_, err = client.SaveAd(context.Background(), &ad)
	return err
}
