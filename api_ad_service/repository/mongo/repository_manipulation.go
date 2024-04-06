package mongo

import (
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/repository/model"
	"api_ad_service/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (r *repository) GetAdvertisement(ctx context.Context, adType pb.AdType) (model.Advertisement, error) {

	var result model.Advertisement

	collection := r.client.Database(utils.EnvMongoDb()).Collection(adType.String())
	now := time.Now().UTC()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"expire_at", bson.D{{"$gt", now}}}}}},
		{{"$sample", bson.D{{"size", 1}}}},
	}

	var results []model.Advertisement
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return result, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return result, err
	}

	if len(results) == 0 {
		return result, errors.New("no advertisement for type " + adType.String())
	}

	return results[0], nil

}

func (r *repository) SaveAdvertisement(ctx context.Context, adType pb.AdType, expireAt time.Time, bucket string, file string) error {
	content := model.Advertisement{ExpireAt: expireAt, Bucket: bucket, File: file}
	collection := r.client.Database(utils.EnvMongoDb()).Collection(adType.String())
	_, err := collection.InsertOne(ctx, content)
	return err
}
