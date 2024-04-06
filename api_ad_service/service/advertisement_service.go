package service

import (
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/repository/model"
	"api_ad_service/repository/mongo"
	"context"
)

type AdvertisementService interface {
	Get(context.Context, *pb.GetAdvertisementRequest) (model.Advertisement, error)
	Save(context.Context, *pb.SaveAdvertisementRequest) error
}
type advertisementService struct {
	repository mongo.Repository
}

var service advertisementService

func InitService(r mongo.Repository) {
	service = advertisementService{repository: r}
}
func GetService() AdvertisementService {
	return service
}
func (s advertisementService) Get(ctx context.Context, request *pb.GetAdvertisementRequest) (model.Advertisement, error) {
	return s.repository.GetAdvertisement(ctx, request.Type)

}
func (s advertisementService) Save(ctx context.Context, request *pb.SaveAdvertisementRequest) error {
	return s.repository.SaveAdvertisement(ctx, request.Type, request.ExpireTime.AsTime(), request.Bucket, request.File)
}
