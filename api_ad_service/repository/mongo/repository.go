package mongo

import (
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/repository/model"
	"api_ad_service/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	GetAdvertisement(context.Context, pb.AdType) (model.Advertisement, error)
	SaveAdvertisement(context.Context, pb.AdType, time.Time, string, string) error
}
type repository struct {
	client *mongo.Client
}

func NewRepository() Repository {
	return &repository{client: utils.GetClient()}
}
