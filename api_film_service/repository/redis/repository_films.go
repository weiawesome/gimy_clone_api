package redis

import (
	"api_film_service/repository/model"
	"context"
	"encoding/json"
	"math/rand"
	"time"
)

func (r *repository) SaveFilmsCache(key string, value []model.Film) error {
	ctx := context.Background()

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	minDuration := GetMinCacheMinute()
	maxDuration := GetMaxCacheMinute()

	randomDuration := rand.Intn(maxDuration-minDuration+1) + minDuration
	expireDuration := time.Minute * time.Duration(randomDuration)

	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}

func (r *repository) LoadFilmsCache(key string) ([]model.Film, error) {
	var response []model.Film

	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(result), &response)
	return response, err
}
