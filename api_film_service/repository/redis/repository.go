package redis

import (
	"api_film_service/repository/model"
	"api_film_service/utils"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SaveFilmsCache(key string, value []model.Film) error
	SaveFilmCache(key string, value model.Film) error
	SaveFilmRoutesCache(key string, value []model.FilmRoute) error
	LoadFilmsCache(key string) ([]model.Film, error)
	LoadFilmCache(key string) (model.Film, error)
	LoadFilmRoutesCache(key string) ([]model.FilmRoute, error)
}

type repository struct {
	client *redis.Client
}

func NewRepository() Repository {
	return &repository{client: utils.GetRedisClient()}
}
