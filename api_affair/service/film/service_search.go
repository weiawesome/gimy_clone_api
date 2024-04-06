package film

import (
	"api_affair/api/response/film"
	pb "api_affair/proto/film_service"
	"api_affair/utils"
	"context"
)

func (s *filmService) GetSearchCelebrityFilms(celebrity string, offset int32, limit int32) (film.FilmsSearch, error) {
	var result film.FilmsSearch

	client := pb.NewFilmClient(s.c)
	request := pb.FilmSearchRequest{SearchType: pb.SearchType_CELEBRITY, Content: celebrity, Offset: offset, Limit: limit}
	films, err := client.GetSearchFilms(context.Background(), &request)
	if err != nil {
		return result, err
	}

	cdnAddress := utils.EnvCDNAddress()

	result.SearchFilms = make([]film.SearchFilm, len(films.Results))
	for i, searchResult := range films.Results {
		result.SearchFilms[i] = film.SearchFilm{
			Id:           searchResult.Id,
			Title:        searchResult.Title,
			State:        searchResult.State,
			Resource:     cdnAddress + searchResult.Resource,
			Category:     searchResult.Category.String(),
			Actors:       searchResult.Actors,
			Directors:    searchResult.Director,
			Location:     searchResult.Location.String(),
			Language:     searchResult.Language,
			ReleaseYear:  searchResult.ReleaseYear,
			UpdateTime:   searchResult.UpdateTime.AsTime(),
			Introduction: searchResult.Introduction,
		}
		if result.SearchFilms[i].Actors == nil {
			result.SearchFilms[i].Actors = []string{}
		}
		if result.SearchFilms[i].Directors == nil {
			result.SearchFilms[i].Directors = []string{}
		}
	}
	return result, nil
}

func (s *filmService) GetSearchContentFilms(content string, offset int32, limit int32) (film.FilmsSearch, error) {
	var result film.FilmsSearch

	client := pb.NewFilmClient(s.c)
	request := pb.FilmSearchRequest{SearchType: pb.SearchType_TITLE, Content: content, Offset: offset, Limit: limit}
	films, err := client.GetSearchFilms(context.Background(), &request)
	if err != nil {
		return result, err
	}

	cdnAddress := utils.EnvCDNAddress()

	result.SearchFilms = make([]film.SearchFilm, len(films.Results))
	for i, searchResult := range films.Results {
		result.SearchFilms[i] = film.SearchFilm{
			Id:           searchResult.Id,
			State:        searchResult.State,
			Title:        searchResult.Title,
			Resource:     cdnAddress + searchResult.Resource,
			Category:     searchResult.Category.String(),
			Actors:       searchResult.Actors,
			Directors:    searchResult.Director,
			Location:     searchResult.Location.String(),
			Language:     searchResult.Language,
			ReleaseYear:  searchResult.ReleaseYear,
			UpdateTime:   searchResult.UpdateTime.AsTime(),
			Introduction: searchResult.Introduction,
		}
		if result.SearchFilms[i].Actors == nil {
			result.SearchFilms[i].Actors = []string{}
		}
		if result.SearchFilms[i].Directors == nil {
			result.SearchFilms[i].Directors = []string{}
		}
	}
	return result, nil
}
