package main

import (
	"api_film_service/api"
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/elasticsearch"
	"api_film_service/repository/mongodb"
	"api_film_service/service"
	"api_film_service/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := utils.InitMongoDb(); err != nil {
		log.Panic("error to connect mongo db")
	}
	defer func() {
		err := utils.CloseMongoDbClient()
		if err != nil {
			return
		}
	}()
	if err := utils.InitElasticsearch(); err != nil {
		log.Panic("error to connect elasticsearch")
	}

	mongodbRepository := mongodb.NewRepository()
	elasticRepository := elasticsearch.NewRepository()
	service.InitService(mongodbRepository, elasticRepository)

	lis, err := net.Listen(utils.EnvServerProtocol(), ":"+utils.EnvServerPort())
	if err != nil {
		log.Panic("failed to listen: " + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterFilmServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		log.Panic("failed to serve: " + err.Error())
	}
}
