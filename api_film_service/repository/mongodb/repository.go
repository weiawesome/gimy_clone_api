package mongodb

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"api_film_service/utils"
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryMongo interface {
	GetBasicFilms(context.Context, *pb.FilmBasicRequest) ([]model.Film, error)
	GetPopularFilmsByType(context.Context, *pb.FilmPopularTypeRequest) ([]model.Film, error)
	GetPopularFilmsByCategory(context.Context, *pb.FilmPopularCategoryRequest) ([]model.Film, error)
	GetSpecificFilm(context.Context, *pb.FilmSpecificRequest) (model.Film, error)
	GetSpecificFilmRoutes(context.Context, *pb.FilmSpecificRequest) ([]model.FilmRoute, error)
	GetRankedFilms(context.Context, *pb.FilmRankedRequest) ([]model.Film, error)
	GetFilterFilms(context.Context, *pb.FilmFilterRequest) ([]model.Film, error)
	GetSearchFilms(context.Context, *pb.FilmSearchRequest) ([]model.Film, error)
	SaveFilm(context.Context, *pb.FilmSaveRequest) error
	SaveFilmEpisode(context.Context, *pb.FilmSaveEpisodeRequest) error
	DeleteFilmEpisode(context.Context, *pb.FilmSaveEpisodeRequest) error
	DeleteFilm(context.Context, *pb.FilmSpecificRequest) error
	AddPopularity(context.Context, *pb.FilmSpecificRequest) error
}
type repository struct {
	client              *mongo.Client
	elasticsearchClient *elasticsearch.Client
}

func NewRepository() RepositoryMongo {
	return &repository{client: utils.GetMongoDbClient()}
}
