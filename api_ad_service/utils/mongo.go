package utils

import (
	pb "api_ad_service/proto/ad_service"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDb() error {
	var err error
	url := "mongodb://" + EnvMongoDbUser() + ":" + EnvMongoDbPassword() + "@" + EnvMongoDbAddress()
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	database := EnvMongoDb()
	for _, typeName := range pb.AdType_name {
		if err := client.Database(database).CreateCollection(context.Background(), typeName); err != nil {
			if mongo.IsDuplicateKeyError(err) {
				continue
			} else {
				return err
			}
		}
	}

	return err
}
func GetClient() *mongo.Client {
	return client
}
func CloseClient() error {
	return client.Disconnect(context.Background())
}
