package service

import (
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/repository/model"
	"api_ad_service/repository/mongo"
	"api_ad_service/repository/redis"
	"api_ad_service/utils"
	"context"
	"log"
	"time"
)

type AdvertisementService interface {
	Get(context.Context, *pb.GetAdvertisementRequest) (model.Advertisement, error)
	Save(context.Context, *pb.SaveAdvertisementRequest) error
}
type advertisementService struct {
	redisRepository redis.Repository
	repository      mongo.Repository
}

var service advertisementService

func InitService(r mongo.Repository, rr redis.Repository) {
	service = advertisementService{repository: r, redisRepository: rr}
}
func GetService() AdvertisementService {
	return service
}
func (s advertisementService) Get(ctx context.Context, request *pb.GetAdvertisementRequest) (model.Advertisement, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := request.Type.String()

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		response, err := s.redisRepository.LoadCache(key)
		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			return response, nil
		} else {
			result, err := s.repository.GetAdvertisement(ctx, request.Type)
			if err != nil {
				return result, err
			}
			go func() {
				err := s.redisRepository.SaveCache(key, result)
				if err != nil {
					log.Println("error to save cache " + err.Error())
				}
			}()

			return result, nil
		}
	})

	select {
	case <-ctx.Done():
		return model.Advertisement{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(model.Advertisement), ret.Err
	}

}
func (s advertisementService) Save(ctx context.Context, request *pb.SaveAdvertisementRequest) error {
	return s.repository.SaveAdvertisement(ctx, request.Type, request.ExpireTime.AsTime(), request.Bucket, request.File)
}
