package main

import (
	"api_ad_service/api"
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/repository/mongo"
	"api_ad_service/repository/redis"
	"api_ad_service/service"
	"api_ad_service/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	utils.InitSingleFlight()

	err := utils.InitMongoDb()
	if err != nil {
		log.Panic("error to connect mongo db")
	}
	defer func() {
		err := utils.CloseClient()
		if err != nil {
			return
		}
	}()
	if err := utils.InitRedis(); err != nil {
		log.Panic("error to connect redis db")
		return
	}
	defer func() {
		err := utils.CloseRedis()
		if err != nil {
			return
		}
	}()

	repository := mongo.NewRepository()
	redisRepository := redis.NewRepository()
	service.InitService(repository, redisRepository)

	lis, err := net.Listen(utils.EnvServerProtocol(), ":"+utils.EnvServerPort())
	if err != nil {
		log.Panic("failed to listen: " + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterAdvertisementServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		log.Panic("failed to serve: " + err.Error())
	}
}
