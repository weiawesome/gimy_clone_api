package film

import (
	"api_affair/api/response/film"
	"google.golang.org/grpc"
)

type Service interface {
	GetBasicFilms(int32, int32) (film.Films, error)
	GetFilmInformation(string) (film.FilmInformation, error)
	GetFilmRouteInformation(string) (film.FilmRouteInformation, error)
	GetFilterFilms(string, string, string, uint32, string, int32, int32) (film.Films, error)
	GetPopularTypeFilms(string, int32, int32) (film.Films, error)
	GetPopularCategoryFilms(string, int32, int32) (film.Films, error)
	GetRankedFilms(string, int32, int32) (film.FilmsRanked, error)
	GetSearchCelebrityFilms(string, int32, int32) (film.FilmsSearch, error)
	GetSearchContentFilms(string, int32, int32) (film.FilmsSearch, error)
}

type filmService struct {
	c *grpc.ClientConn
}

func NewFilmService(c *grpc.ClientConn) Service {
	return &filmService{c: c}
}
