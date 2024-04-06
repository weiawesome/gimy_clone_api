package elasticsearch

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"api_film_service/utils"
	"context"
	"github.com/elastic/go-elasticsearch/v8"
)

type RepositoryElasticsearch interface {
	FilmSave(context.Context, model.Film) error
	SearchFilmsGet(context.Context, *pb.FilmSearchRequest) ([]model.Film, error)
}
type repository struct {
	client *elasticsearch.Client
}

func NewRepository() RepositoryElasticsearch {
	return &repository{client: utils.GetElasticsearchClient()}
}
