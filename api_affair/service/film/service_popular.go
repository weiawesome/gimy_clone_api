package film

import (
	"api_affair/api/response/film"
	pb "api_affair/proto/film_service"
	"api_affair/utils"
	"context"
)

func (s *filmService) GetPopularTypeFilms(filmType string, offset int32, limit int32) (film.Films, error) {
	var result film.Films

	client := pb.NewFilmClient(s.c)
	request := pb.FilmPopularTypeRequest{Type: pb.MediaType(pb.MediaType_value[filmType]), Offset: offset, Limit: limit}
	films, err := client.GetPopularTypeFilms(context.Background(), &request)
	if err != nil {
		return result, err
	}

	cdnAddress := utils.EnvCDNAddress()

	result.Films = make([]film.Film, len(films.FilmInformation))
	for i, information := range films.FilmInformation {
		result.Films[i] = film.Film{
			Id:       information.Id,
			Title:    information.Title,
			Resource: cdnAddress + information.Resource,
			State:    information.State,
			Actors:   information.Actors,
		}
		if result.Films[i].Actors == nil {
			result.Films[i].Actors = []string{}
		}
	}
	return result, nil
}

func (s *filmService) GetPopularCategoryFilms(filmCategory string, offset int32, limit int32) (film.Films, error) {
	var result film.Films

	client := pb.NewFilmClient(s.c)
	request := pb.FilmPopularCategoryRequest{Category: pb.MediaCategory(pb.MediaCategory_value[filmCategory]), Offset: offset, Limit: limit}
	films, err := client.GetPopularCategoryFilms(context.Background(), &request)
	if err != nil {
		return result, err
	}

	cdnAddress := utils.EnvCDNAddress()

	result.Films = make([]film.Film, len(films.FilmInformation))
	for i, information := range films.FilmInformation {
		result.Films[i] = film.Film{
			Id:       information.Id,
			Title:    information.Title,
			Resource: cdnAddress + information.Resource,
			State:    information.State,
			Actors:   information.Actors,
		}
		if result.Films[i].Actors == nil {
			result.Films[i].Actors = []string{}
		}
	}
	return result, nil
}
