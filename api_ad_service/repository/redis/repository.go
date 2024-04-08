package redis

import (
	"api_ad_service/repository/model"
	"api_ad_service/utils"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SaveCache(key string, value model.Advertisement) error
	LoadCache(key string) (model.Advertisement, error)
}

type repository struct {
	client *redis.Client
}

func NewRepository() Repository {
	return &repository{client: utils.GetRedisClient()}
}
