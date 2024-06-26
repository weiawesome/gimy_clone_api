package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/elasticsearch"
	"api_film_service/repository/mongodb"
	"api_film_service/repository/redis"
	"context"
)

type FilmService interface {
	AddPopularity(context.Context, *pb.FilmSpecificRequest) error
	GetBasicFilms(context.Context, *pb.FilmBasicRequest) (*pb.FilmInformationListReply, error)
	GetFilterFilms(context.Context, *pb.FilmFilterRequest) (*pb.FilmInformationListReply, error)
	GetRankedFilms(context.Context, *pb.FilmRankedRequest) (*pb.FilmRankedReply, error)
	GetTypePopularityFilms(context.Context, *pb.FilmPopularTypeRequest) (*pb.FilmInformationListReply, error)
	GetCategoryPopularityFilms(context.Context, *pb.FilmPopularCategoryRequest) (*pb.FilmInformationListReply, error)
	GetSearchFilms(context.Context, *pb.FilmSearchRequest) (*pb.FilmSearchReply, error)
	GetSpecificFilm(context.Context, *pb.FilmSpecificRequest) (*pb.FilmSpecificReply, error)
	GetSpecificFilmRoutes(context.Context, *pb.FilmSpecificRequest) (*pb.FilmSpecificRoutesReply, error)
	SaveFilm(context.Context, *pb.FilmSaveRequest) error
	SaveFilmEpisode(context.Context, *pb.FilmSaveEpisodeRequest) error
	DeleteFilmEpisode(context.Context, *pb.FilmSaveEpisodeRequest) error
	DeleteFilm(context.Context, *pb.FilmSpecificRequest) error
	DeleteFilmToSearchEngine(*pb.FilmSpecificRequest) error
	AddFilmToSearchEngine(context.Context, *pb.FilmSpecificRequest) error
}
type filmService struct {
	redisRepository         redis.Repository
	mongodbRepository       mongodb.RepositoryMongo
	elasticsearchRepository elasticsearch.RepositoryElasticsearch
}

var service filmService

func InitService(mr mongodb.RepositoryMongo, er elasticsearch.RepositoryElasticsearch, rr redis.Repository) {
	service = filmService{mongodbRepository: mr, elasticsearchRepository: er, redisRepository: rr}
}
func GetService() FilmService {
	return service
}
