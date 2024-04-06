package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDb() error {
	var err error
	url := "mongodb://" + EnvMongoDbUser() + ":" + EnvMongoDbPassword() + "@" + EnvMongoDbAddress()
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	database := EnvMongoDb()

	for _, typeName := range GetCollectionList() {
		if err := mongoClient.Database(database).CreateCollection(context.Background(), typeName); err != nil {
			if mongo.IsDuplicateKeyError(err) {
				continue
			} else {
				return err
			}
		}
	}

	return err
}
func GetMongoDbClient() *mongo.Client {
	return mongoClient
}
func CloseMongoDbClient() error {
	return mongoClient.Disconnect(context.Background())
}
